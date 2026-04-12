package msnet

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log/slog"
	"net"

	"github.com/zhyonc/msnet/internal/crypt"
)

type clientSocket struct {
	id         int32
	delegate   CClientSocketDelegate
	sock       net.Conn
	addr       net.Addr
	recvBuff   []byte
	sendBuff   []byte
	packetRecv *iPacket
	seqRcv     [4]byte
	seqSnd     [4]byte
}

func NewCClientSocket(delegate CClientSocketDelegate, conn net.Conn, rcvIV []byte, sndIV []byte) CClientSocket {
	if gSetting == nil {
		panic("Please use msnet.New(setting) to install package")
	}
	c := &clientSocket{
		delegate:   delegate,
		sock:       conn,
		addr:       conn.RemoteAddr(),
		packetRecv: &iPacket{},
	}
	if len(rcvIV) == 0 {
		rand.Read(c.seqRcv[:])
	} else {
		c.seqRcv = [4]byte(rcvIV)
	}
	if len(sndIV) == 0 {
		rand.Read(c.seqSnd[:])
	} else {
		c.seqSnd = [4]byte(sndIV)
	}
	return c
}

// SetID implements [CClientSocket].
func (cs *clientSocket) SetID(id int32) {
	cs.id = id
}

// GetID implements [CClientSocket].
func (cs *clientSocket) GetID() int32 {
	return cs.id
}

// GetAddr implements [CClientSocket].
func (cs *clientSocket) GetAddr() string {
	return cs.addr.String()
}

// XORRecv implements [CClientSocket].
func (cs *clientSocket) XORRecv(buf []byte) {
	// The server must use the same XOR key to recover the original packet
	if gSetting.RecvXOR == 0 {
		return
	}
	for i := range buf {
		buf[i] ^= gSetting.RecvXOR
	}
}

// XORSend implements [CClientSocket].
func (cs *clientSocket) XORSend(buf []byte) {
	// The client must use the same XOR key to recover the original packet
	if gSetting.SendXOR == 0 {
		return
	}
	for i := range buf {
		buf[i] ^= gSetting.SendXOR
	}
}

// OnRead implements [CClientSocket].
func (cs *clientSocket) OnRead() {
	readSize := headerLength
	isHeader := true
	defer cs.Close()
	for {
		cs.recvBuff = make([]byte, readSize)
		_, err := cs.sock.Read(cs.recvBuff)
		if err != nil {
			slog.Error("[OnRead]", "err", err)
			return
		}
		cs.XORRecv(cs.recvBuff)
		// CClientSocket::ManipulatePacket
		if isHeader {
			// Decode packet header
			cs.packetRecv.AppendBuffer(cs.recvBuff, true)
			clientVersion := cs.packetRecv.RawSeq ^ binary.LittleEndian.Uint16(cs.seqRcv[2:4])
			if clientVersion != gSetting.MSVersion {
				if clientVersion == 223 {
					// GMSCW
				} else {
					cs.OnError(fmt.Errorf("failed to decode packet header"))
					return
				}
			}
			readSize = cs.packetRecv.DataLen
		} else {
			// Decode packet data
			iPacket := NewCInPacket(cs.recvBuff)
			iPacket.DecryptData(cs.seqRcv[:]) // Decrypt using AES OFB mode
			// Refresh m_uSeqRcv value
			if gSetting.IsXORCipher {
				(*crypt.XORCipher).Shuffle(nil, cs.seqRcv[:])
			} else {
				(*crypt.CIGCipher).InnoHash(nil, cs.seqRcv[:])
			}
			cs.delegate.DebugInPacketLog(cs.id, iPacket)
			cs.delegate.ProcessPacket(cs, iPacket)
			readSize = headerLength
		}
		isHeader = !isHeader
	}
}

// OnConnect implements [CClientSocket].
func (cs *clientSocket) OnConnect() {
	oPacket := cs.delegate.NewConnectPacket(gSetting.MSRegion, gSetting.MSVersion, gSetting.MSMinorVersion, cs.seqRcv, cs.seqSnd)
	if oPacket == nil {
		oPacket = NewCOutPacket(0)
		oPacket.EncodeStr(gSetting.MSMinorVersion)
		oPacket.EncodeBuffer(cs.seqRcv[:])
		oPacket.EncodeBuffer(cs.seqSnd[:])
		oPacket.Encode1(int8(gSetting.MSRegion))
	}
	cs.sendBuff = oPacket.MakeBufferList(false, nil)
	cs.XORSend(cs.sendBuff)
	cs.Flush()
}

// OnReceiveHotfix implements [CClientSocket].
func (cs *clientSocket) OnReceiveHotfix() {
	oPacket := cs.delegate.NewHotfixPacket()
	if oPacket == nil {
		return
	}
	cs.SendPacket(oPacket)
}

// OnAliveReq implements [CClientSocket].
func (cs *clientSocket) OnAliveReq(LP_AliveReq uint16) {
	var oPacket COutPacket
	if gSetting.IsXORCipher {
		oPacket = NewCOutPacketByte(uint8(LP_AliveReq))
	} else {
		oPacket = NewCOutPacket(LP_AliveReq)
	}
	cs.SendPacket(oPacket)
}

// OnMigrateCommand implements [CClientSocket].
func (cs *clientSocket) OnMigrateCommand(LP_MigrateCommand uint16, ip string, port int16) {
	var oPacket COutPacket
	if gSetting.IsXORCipher {
		oPacket = NewCOutPacketByte(uint8(LP_MigrateCommand))
	} else {
		oPacket = NewCOutPacket(LP_MigrateCommand)
	}
	oPacket.EncodeBool(true)
	ipBytes := net.ParseIP(ip)
	if ipBytes == nil {
		slog.Warn("Invaild ip on migrate command", "ip", ip)
		oPacket.EncodeBuffer([]byte{127, 0, 0, 1})
	} else {
		oPacket.EncodeBuffer(ipBytes.To4())
	}
	oPacket.Encode2(port)
	cs.SendPacket(oPacket)
}

// SendPacket implements [CClientSocket].
func (cs *clientSocket) SendPacket(oPacket COutPacket) {
	cs.delegate.DebugOutPacketLog(cs.id, oPacket)
	cs.sendBuff = oPacket.MakeBufferList(true, cs.seqSnd[:])
	// Refresh SeqSnd value
	if gSetting.IsXORCipher {
		(*crypt.XORCipher).Shuffle(nil, cs.seqSnd[:])
	} else {
		(*crypt.CIGCipher).InnoHash(nil, cs.seqSnd[:])
	}
	cs.XORSend(cs.sendBuff)
	cs.Flush()
}

// Flush implements [CClientSocket].
func (cs *clientSocket) Flush() {
	_, err := cs.sock.Write(cs.sendBuff)
	if err != nil {
		slog.Error("Failed to send packet to client", "err", err)
		return
	}
}

// OnError implements [CClientSocket].
func (cs *clientSocket) OnError(err error) {
	slog.Error("[ClientSocket] OnError", "err", err)
	cs.Close()
}

// Close implements [CClientSocket].
func (cs *clientSocket) Close() {
	if cs.delegate != nil {
		cs.delegate.SocketClose(cs.id)
	}
	cs.sock.Close()
	cs = nil
}
