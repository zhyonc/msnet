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
	"sync"
	"time"

	"github.com/zhyonc/msnet/internal/crypt"
)

type clientSocket struct {
	id               int32
	delegate         CClientSocketDelegate
	sock             net.Conn
	addr             net.Addr
	recvCipherType   CipherType
	sendCipherType   CipherType
	seqRcv           []byte
	seqSnd           []byte
	recvBuff         []byte
	sendBuff         []byte
	packetRecv       *iPacket
	CPMap            map[uint16]uint16
	stopChan         chan struct{}
	closeOnce        sync.Once
	lastAliveAckTime time.Time
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
		CPMap:      make(map[uint16]uint16),
		stopChan:   make(chan struct{}),
	}
	// Cipher Type
	cs.recvCipherType = gSetting.RecvCipherType
	cs.sendCipherType = gSetting.SendCipherType
	// IV
	if len(rcvIV) == 0 {
		cs.seqRcv = make([]byte, 4)
		crand.Read(cs.seqRcv[:])
	} else {
		cs.seqRcv = rcvIV
	}
	if len(sndIV) == 0 {
		cs.seqSnd = make([]byte, 4)
		crand.Read(cs.seqSnd[:])
	} else {
		cs.seqSnd = sndIV
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

// XORRecvBuff implements [CClientSocket].
func (cs *clientSocket) XORRecvBuff() {
	// The server must use the same XOR key to recover the original packet
	if gSetting.RecvBuffXOR == 0 {
		return
	}
	for i := range cs.recvBuff {
		cs.recvBuff[i] ^= gSetting.RecvBuffXOR
	}
}

// XORSendBuff implements [CClientSocket].
func (cs *clientSocket) XORSendBuff() {
	// The client must use the same XOR key to recover the original packet
	if gSetting.SendBuffXOR == 0 {
		return
	}
	for i := range cs.sendBuff {
		cs.sendBuff[i] ^= gSetting.SendBuffXOR
	}
}

// OnRead implements [CClientSocket].
func (cs *clientSocket) OnRead() {
	readSize := HEADER_LENGTH
	isHeader := true
	defer cs.Close()
	for {
		if cs.sock == nil {
			return
		}
		cs.recvBuff = make([]byte, readSize)
		_, err := cs.sock.Read(cs.recvBuff)
		if err != nil {
			slog.Error("[OnRead]", "err", err)
			return
		}
		cs.XORRecvBuff()
		// CClientSocket::ManipulatePacket
		if isHeader {
			// Decode packet header
			cs.packetRecv.DecryptHeader(cs.recvCipherType, cs.recvBuff)
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
			iPacket.DecryptData(cs.recvCipherType, cs.seqRcv)
			cs.InnoHash(cs.recvCipherType, cs.seqRcv)
			cs.delegate.DebugInPacketLog(cs.id, iPacket)
			cs.delegate.ProcessPacket(cs, iPacket)
			readSize = HEADER_LENGTH
		}
		isHeader = !isHeader
	}
}

// OnConnect implements [CClientSocket].
func (cs *clientSocket) OnConnect() {
	if cs.delegate == nil {
		return
	}
	oPacket := cs.delegate.NewConnectPacket(gSetting.MSRegion, gSetting.MSVersion, gSetting.MSMinorVersion, cs.seqRcv, cs.seqSnd)
	if oPacket == nil {
		oPacket = NewCOutPacket()
		oPacket.EncodeStr(gSetting.MSMinorVersion)
		oPacket.EncodeBuffer(cs.seqRcv)
		oPacket.EncodeBuffer(cs.seqSnd)
		oPacket.Encode1(int8(gSetting.MSRegion))
	}
	cs.sendBuff = oPacket.MakeBufferList(NullCipher, nil)
	cs.XORSendBuff()
	cs.Flush()
}

// InnoHash implements [CClientSocket].
func (cs *clientSocket) InnoHash(cipherType CipherType, dwKey []byte) {
	switch cipherType {
	case XORCipher:
		gXORCipher.Shuffle(dwKey)
	case AESCipher, LinearCipher:
		(*crypt.CIGCipher).InnoHash(nil, dwKey)
	default:
		slog.Warn("Unknown cipher type when Stepping", "cipherType", cipherType)
	}
}

// LoopAliveAck implements [CClientSocket].
func (cs *clientSocket) LoopAliveAck(aliveAckSec int) {
	if aliveAckSec <= 0 {
		return
	}
	aliveTimeout := time.Duration(aliveAckSec) * time.Second
	cs.lastAliveAckTime = time.Now()
	go func() {
		ticker := time.NewTicker(aliveTimeout / 2)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if time.Since(cs.lastAliveAckTime) > aliveTimeout {
					cs.OnError(fmt.Errorf("failed to check client socket %d alive", cs.GetID()))
					return
				}
			case <-cs.stopChan:
				return
			}
		}
	}()
}

// OnAliveAck implements [CClientSocket].
func (cs *clientSocket) OnAliveAck() {
	cs.lastAliveAckTime = time.Now()
}

// LoopAliveReq implements [CClientSocket].
func (cs *clientSocket) LoopAliveReq(aliveReqSec int, LP_AliveReq uint16) {
	if aliveReqSec <= 0 {
		return
	}
	aliveReqTime := time.Duration(aliveReqSec) * time.Second
	go func() {
		ticker := time.NewTicker(aliveReqTime)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				cs.OnAliveReq(LP_AliveReq)
			case <-cs.stopChan:
				return
			}
		}
	}()
}

// OnAliveReq implements [CClientSocket].
func (cs *clientSocket) OnAliveReq(LP_AliveReq uint16) {
	oPacket := NewCOutPacket(LP_AliveReq)
	cs.SendPacket(oPacket)
}

// SetRecvCipherType implements [CClientSocket].
func (cs *clientSocket) SetRecvCipherType(cipherType CipherType) {
	cs.recvCipherType = cipherType
}

// SetSendCipherType implements [CClientSocket].
func (cs *clientSocket) SetSendCipherType(cipherType CipherType) {
	cs.sendCipherType = cipherType
}

// OnOpcodeEncryption implements [CClientSocket].
func (cs *clientSocket) OnOpcodeEncryption(LP_OpcodeEncryption uint16, startOpcode uint16, endOpcode uint16, isSplit bool) {
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
	encryptedBuf, err := gDESCipher.Encrypt(content)
	if err != nil {
		slog.Error("[CClientSocket] Failed to encrypt content using TripleDESCipher", "err", err)
		return
	}
	oPacket := NewCOutPacket(LP_OpcodeEncryption)
	if !isSplit {
		oPacket.Encode4(gDESCipher.GetBlockSize())
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

// SendPacket implements [CClientSocket].
func (cs *clientSocket) SendPacket(oPacket COutPacket) {
	if cs.delegate == nil {
		return
	}
	cs.delegate.DebugOutPacketLog(cs.id, oPacket)
	cs.sendBuff = oPacket.MakeBufferList(cs.sendCipherType, cs.seqSnd)
	cs.InnoHash(cs.sendCipherType, cs.seqSnd)
	cs.XORSendBuff()
	cs.Flush()
}

// Flush implements [CClientSocket].
func (cs *clientSocket) Flush() {
	if cs.sock == nil {
		return
	}
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
	cs.closeOnce.Do(func() {
		if cs.stopChan != nil {
			close(cs.stopChan)
			cs.stopChan = nil
		}
		if cs.sock != nil {
			cs.sock.Close()
			cs.sock = nil
		}
		if cs.delegate != nil {
			cs.delegate.SocketClose(cs.id)
			cs.delegate = nil
		}
	})
}
