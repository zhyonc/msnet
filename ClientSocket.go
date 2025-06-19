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
	impl       CClientSocketImpl
	sock       net.Conn
	addr       net.Addr
	recvBuff   []byte
	sendBuff   []byte
	packetRecv *iPacket
	seqRcv     [4]byte
	seqSnd     [4]byte
}

func NewCClientSocket(conn net.Conn, rcvIV []byte, sndIV []byte, impl CClientSocketImpl) CClientSocket {
	if gSetting == nil {
		panic("Please use msnet.New(setting) to install package")
	}
	c := &clientSocket{
		impl:       impl,
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

// OnMigrateCommand implements CClientSocket.
func (c *clientSocket) OnMigrateCommand(LP_MigrateCommand int16, ip string, port int16) {
	oPacket := NewCOutPacket(LP_MigrateCommand)
	oPacket.EncodeBool(true)
	ipBytes := net.ParseIP(ip)
	if ipBytes == nil {
		slog.Warn("Invaild ip on migrate command", "ip", ip)
		oPacket.EncodeBuffer([]byte{127, 0, 0, 1})
	} else {
		oPacket.EncodeBuffer(ipBytes.To4())
	}
	oPacket.Encode2(port)
	c.SendPacket(oPacket)
}

// OnConnect implements CClientSocket.
func (c *clientSocket) OnConnect() {
	oPacket := NewCOutPacket(0)
	oPacket.EncodeStr(gSetting.MSMinorVersion)
	oPacket.EncodeBuffer(c.seqRcv[:])
	oPacket.EncodeBuffer(c.seqSnd[:])
	oPacket.Encode1(int8(gSetting.MSRegion))
	c.sendBuff = oPacket.MakeBufferList(gSetting.MSVersion, false, nil)
	c.XORSend(c.sendBuff)
	c.Flush()
}

// Flush implements CClientSocket.
func (c *clientSocket) Flush() {
	_, err := c.sock.Write(c.sendBuff)
	if err != nil {
		slog.Error("Failed to send packet to client", "err", err)
		return
	}
}

// OnAliveReq implements CClientSocket.
func (c *clientSocket) OnAliveReq(LP_AliveReq int16) {
	oPacket := NewCOutPacket(LP_AliveReq)
	c.SendPacket(oPacket)
}

// XORRecv implements CClientSocket.
func (c *clientSocket) XORRecv(buf []byte) {
	// The server must use the same XOR key to recover the original packet
	if gSetting.RecvXOR == 0 {
		return
	}
	for i := range buf {
		buf[i] ^= gSetting.RecvXOR
	}
}

// XORSend implements CClientSocket.
func (c *clientSocket) XORSend(buf []byte) {
	// The client must use the same XOR key to recover the original packet
	if gSetting.SendXOR == 0 {
		return
	}
	for i := range buf {
		buf[i] ^= gSetting.SendXOR
	}
}

// OnRead implements CClientSocket.
func (c *clientSocket) OnRead() {
	readSize := headerLength
	isHeader := true
	defer c.Close()
	for {
		c.recvBuff = make([]byte, readSize)
		_, err := c.sock.Read(c.recvBuff)
		if err != nil {
			slog.Error("[OnRead]", "err", err)
			return
		}
		c.XORRecv(c.recvBuff)
		// CClientSocket::ManipulatePacket
		if isHeader {
			// Decode packet header
			c.packetRecv.AppendBuffer(c.recvBuff, true)
			HIWORD := binary.LittleEndian.Uint16(c.seqRcv[2:4])
			if c.packetRecv.RawSeq^HIWORD != gSetting.MSVersion {
				c.OnError(fmt.Errorf("failed to decode packet header"))
				return
			}
			readSize = c.packetRecv.DataLen
		} else {
			// Decode packet data
			iPacket := NewCInPacket(c.recvBuff)
			iPacket.DecryptData(c.seqRcv[:])              // Decrypt using AES OFB mode
			(*crypt.CIGCipher).InnoHash(nil, c.seqRcv[:]) // Refresh m_uSeqRcv value
			c.impl.DebugInPacketLog(iPacket)
			c.impl.ProcessPacket(c, iPacket)
			readSize = headerLength
		}
		isHeader = !isHeader
	}
}

// SendPacket implements CClientSocket.
func (c *clientSocket) SendPacket(oPacket COutPacket) {
	c.impl.DebugOutPacketLog(oPacket)
	c.sendBuff = oPacket.MakeBufferList(gSetting.MSVersion, true, c.seqSnd[:])
	(*crypt.CIGCipher).InnoHash(nil, c.seqSnd[:]) // Refresh SeqSnd value
	c.XORSend(c.sendBuff)
	c.Flush()
}

// OnError implements CClientSocket.
func (c *clientSocket) OnError(err error) {
	slog.Error("[ClientSocket] OnError", "err", err)
	c.Close()
}

// Close implements CClientSocket.
func (c *clientSocket) Close() {
	if c.impl != nil {
		c.impl.SocketClose()
	}
	c.sock.Close()
	c = nil
}

// GetAddr implements CClientSocket.
func (c *clientSocket) GetAddr() string {
	return c.addr.String()
}
