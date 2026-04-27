package msnet

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log/slog"
	"math"
	mrand "math/rand/v2"
	"net"
	"strings"

	"github.com/zhyonc/msnet/def"
	"github.com/zhyonc/msnet/internal/crypt"
)

type clientSocket struct {
	id             int32
	delegate       CClientSocketDelegate
	sock           net.Conn
	addr           net.Addr
	recvBuff       []byte
	sendBuff       []byte
	packetRecv     *iPacket
	seqRcv         [4]byte
	seqSnd         [4]byte
	desCipher      *crypt.TripleDESCipher
	CPMap          map[uint16]uint16
	isLinearCipher bool
}

func NewCClientSocket(delegate CClientSocketDelegate, conn net.Conn, rcvIV []byte, sndIV []byte) CClientSocket {
	if gSetting == nil {
		panic("[CClientSocket] Please use msnet.New(setting) to install package")
	}
	cs := &clientSocket{
		delegate:   delegate,
		sock:       conn,
		addr:       conn.RemoteAddr(),
		packetRecv: &iPacket{},
	}
	// IV
	if len(rcvIV) == 0 {
		crand.Read(cs.seqRcv[:])
	} else {
		cs.seqRcv = [4]byte(rcvIV)
	}
	if len(sndIV) == 0 {
		crand.Read(cs.seqSnd[:])
	} else {
		cs.seqSnd = [4]byte(sndIV)
	}
	// OpcodeEncryption
	if gSetting.DESKey != "" {
		desCipher, err := crypt.NewTripleDESCipher(gSetting.DESKey)
		if err != nil {
			panic(err)
		}
		cs.desCipher = desCipher
	}
	return cs
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
	readSize := def.HEADER_LENGTH
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
			cs.packetRecv.DecryptHeader(cs.recvBuff)
			clientVersion := cs.packetRecv.RawSeq ^ binary.LittleEndian.Uint16(cs.seqRcv[2:4])
			if clientVersion != gSetting.MSVersion {
				if clientVersion == 223 {
					// GMSCW v1.0
				} else {
					cs.OnError(fmt.Errorf("failed to decode packet header"))
					return
				}
			}
			readSize = cs.packetRecv.DataLen
		} else {
			// Decode packet data
			iPacket := NewCInPacket(cs.recvBuff)
			iPacket.DecryptData(cs.seqRcv[:])
			cs.Stepping(cs.seqRcv[:])
			cs.delegate.DebugInPacketLog(cs.id, iPacket)
			cs.delegate.ProcessPacket(cs, iPacket)
			readSize = def.HEADER_LENGTH
		}
		isHeader = !isHeader
	}
}

// OnConnect implements [CClientSocket].
func (cs *clientSocket) OnConnect() {
	oPacket := cs.delegate.NewConnectPacket(gSetting.MSRegion, gSetting.MSVersion, gSetting.MSMinorVersion, cs.seqRcv, cs.seqSnd)
	if oPacket == nil {
		oPacket = NewCOutPacket()
		oPacket.EncodeStr(gSetting.MSMinorVersion)
		oPacket.EncodeBuffer(cs.seqRcv[:])
		oPacket.EncodeBuffer(cs.seqSnd[:])
		oPacket.Encode1(int8(gSetting.MSRegion))
	}
	cs.sendBuff = oPacket.MakeBufferList(def.NullCipher, nil)
	cs.XORSend(cs.sendBuff)
	cs.Flush()
}

// OnReceiveHotfix implements [CClientSocket].
func (cs *clientSocket) OnReceiveHotfix(LP_ApplyHotfix uint16) {
	oPacket := cs.delegate.NewHotfixPacket()
	if oPacket == nil {
		oPacket = NewCOutPacket(LP_ApplyHotfix)
		oPacket.Encode1(1)
	}
	cs.SendPacket(oPacket)
}

// OnAliveReq implements [CClientSocket].
func (cs *clientSocket) OnAliveReq(LP_AliveReq uint16) {
	var oPacket COutPacket
	if gSetting.IsTypeHeader1Byte {
		oPacket = NewCOutPacket(uint8(LP_AliveReq))
	} else {
		oPacket = NewCOutPacket(LP_AliveReq)
	}
	cs.SendPacket(oPacket)
}

// OnMigrateCommand implements [CClientSocket].
func (cs *clientSocket) OnMigrateCommand(LP_MigrateCommand uint16, ip string, port int16) {
	var oPacket COutPacket
	if gSetting.IsTypeHeader1Byte {
		oPacket = NewCOutPacket(uint8(LP_MigrateCommand))
	} else {
		oPacket = NewCOutPacket(LP_MigrateCommand)
	}
	oPacket.EncodeBool(true)
	ipBytes := net.ParseIP(ip)
	if ipBytes == nil {
		slog.Warn("[CClientSocket] Invaild ip on migrate command", "ip", ip)
		oPacket.EncodeBuffer([]byte{127, 0, 0, 1})
	} else {
		oPacket.EncodeBuffer(ipBytes.To4())
	}
	oPacket.Encode2(port)
	cs.SendPacket(oPacket)
}

// OnOpcodeEncryption implements [CClientSocket].
func (cs *clientSocket) OnOpcodeEncryption(LP_OpcodeEncryption uint16, startOpcode uint16, endOpcode uint16, isSplit bool) {
	cs.CPMap = make(map[uint16]uint16)
	var builder strings.Builder
	if isSplit {
		// Write opcodes range
		fmt.Fprintf(&builder, "%04d", endOpcode-startOpcode+1)
		builder.WriteRune('|')
	}
	// Convert opcode to rand num
	max := int32(math.MaxUint16)
	for op := startOpcode; op <= endOpcode; op++ {
		// Bind rand num with opcode
		min := int32(op)
		var randNum uint16
		for {
			randNum = uint16(mrand.Int32N(max-min+1) + min)
			if _, ok := cs.CPMap[randNum]; !ok {
				break
			}
		}
		cs.CPMap[randNum] = op
		// Write rand num to replace opcode
		fmt.Fprintf(&builder, "%04d", randNum)
		if isSplit && op < endOpcode {
			// Add separator symbol
			builder.WriteRune('|')
		}
	}
	// Using TripleDESCipher encrypt content
	content := builder.String()
	encryptedBuf, err := cs.desCipher.Encrypt(content)
	if err != nil {
		slog.Error("[CClientSocket] Failed to encrypt content using TripleDESCipher", "err", err)
		return
	}
	oPacket := NewCOutPacket(LP_OpcodeEncryption)
	if !isSplit {
		oPacket.Encode4(cs.desCipher.GetBlockSize())
	}
	oPacket.Encode4(int32(len(encryptedBuf)))
	oPacket.EncodeBuffer(encryptedBuf)
	cs.SendPacket(oPacket)
}

// DecryptOpcode implements [CClientSocket].
func (cs *clientSocket) DecryptOpcode(randNum uint16) uint16 {
	op, ok := cs.CPMap[randNum]
	if !ok {
		return 0
	}
	return op
}

// SetLinearCipher implements [CClientSocket].
func (cs *clientSocket) SetLinearCipher(toggle bool) {
	cs.isLinearCipher = toggle
}

// SendPacket implements [CClientSocket].
func (cs *clientSocket) SendPacket(oPacket COutPacket) {
	cs.delegate.DebugOutPacketLog(cs.id, oPacket)
	if cs.isLinearCipher {
		cs.sendBuff = oPacket.MakeBufferList(def.LinearCipher, cs.seqSnd[:])
	} else {
		cs.sendBuff = oPacket.MakeBufferList(gSetting.CipherType, cs.seqSnd[:])
	}
	cs.Stepping(cs.seqSnd[:])
	cs.XORSend(cs.sendBuff)
	cs.Flush()
}

// Stepping implements [CClientSocket].
func (cs *clientSocket) Stepping(iv []byte) {
	if iv == nil {
		return
	}
	// Refresh SeqSnd value
	switch gSetting.CipherType {
	case def.XORCipher:
		(*crypt.XORCipher).Shuffle(nil, iv)
	case def.AESCipher, def.LinearCipher:
		(*crypt.CIGCipher).InnoHash(nil, iv)
	}
}

// Flush implements [CClientSocket].
func (cs *clientSocket) Flush() {
	_, err := cs.sock.Write(cs.sendBuff)
	if err != nil {
		slog.Error("[CClientSocket] Failed to send packet to client", "err", err)
		return
	}
}

// OnError implements [CClientSocket].
func (cs *clientSocket) OnError(err error) {
	slog.Error("[CClientSocket] OnError", "err", err)
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
