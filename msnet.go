package msnet

import (
	"time"

	"github.com/zhyonc/msnet/internal/crypt"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

var (
	gSetting *Setting
)

func New(setting *Setting) {
	// Setting
	gSetting = setting
	// Language coder
	switch gSetting.MSRegion {
	case CMS:
		langEncoder = simplifiedchinese.GBK.NewEncoder()
		langDecoder = simplifiedchinese.GBK.NewDecoder()
	case TMS:
		langEncoder = traditionalchinese.Big5.NewEncoder()
		langDecoder = traditionalchinese.Big5.NewDecoder()
	default:
		langEncoder = encoding.Nop.NewEncoder()
		langDecoder = encoding.Nop.NewDecoder()
	}
	// AESInitType
	crypt.AESInitType = uint8(gSetting.AESInitType)
}

type CClientSocket interface {
	SetID(id int32)
	GetID() int32
	GetAddr() string
	XORRecv(buf []byte)
	XORSend(buf []byte)
	OnRead()
	OnConnect()
	LoopAliveAck(aliveAckSec int)
	OnAliveAck()
	LoopAliveReq(aliveReqSec int, LP_AliveReq uint16)
	OnAliveReq(LP_AliveReq uint16)
	OnOpcodeEncryption(LP_OpcodeEncryption uint16, startOpcode uint16, endOpcode uint16, isSplit bool)
	DecryptOpcode(randNum uint16) uint16
	SetLinearCipher(toggle bool)
	SendPacket(oPacket COutPacket)
	Stepping(iv []byte)
	Flush()
	OnError(err error)
	Close()
}

type CClientSocketDelegate interface {
	DebugInPacketLog(id int32, iPacket CInPacket)
	DebugOutPacketLog(id int32, oPacket COutPacket)
	NewConnectPacket(region Region, version uint16, minorVersion string, seqRcv [4]byte, seqSnd [4]byte) COutPacket
	ProcessPacket(cs CClientSocket, iPacket CInPacket)
	SocketClose(id int32)
}

type CInPacket interface {
	DecryptHeader(pBuff []byte)
	DecryptData(dwKey []byte)
	GetType() uint16
	GetRemain() int
	GetOffset() int
	GetLength() int
	DecodeBool() bool
	Decode1() int8
	Decode2() int16
	Decode4() int32
	Decode8() int64
	DecodeFT() time.Time
	DecodeStr() string
	DecodeLocalStr() string
	DecodeLocalName() string
	DecodeBuffer(uSize int) []byte
	DumpString(nSize int) string
	Clear()
}

type COutPacket interface {
	GetType() uint16
	GetSendBuffer() []byte
	GetOffset() int
	GetLength() int
	EncodeBool(b bool)
	Encode1(n int8)
	Encode2(n int16)
	Encode4(n int32)
	Encode8(n int64)
	EncodeFT(t time.Time)
	EncodeStr(s string)
	EncodeLocalStr(s string)
	EncodeLocalName(s string)
	EncodeBuffer(buf []byte)
	EncryptHeader(pBuff []byte, dataLen int, dwKey []byte)
	MakeBufferList(cipherType CipherType, dwKey []byte) []byte
	DumpString(nSize int) string
}
