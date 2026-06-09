package msnet

import (
	"time"

	"github.com/zhyonc/msnet/internal/crypt"
)

var (
	gSetting      *Setting
	gDESCipher    *crypt.TripleDESCipher
	gAESCipher    *crypt.CAESCipher
	gXORCipher    *crypt.XORCipher
	gLinearCipher *crypt.LinearCipher
)

func New(setting *Setting) {
	gSetting = setting
	SetLocale(gSetting.LocaleRegion)
	if gSetting.DESKey != "" {
		gDESCipher = crypt.NewTripleDESCipher(gSetting.DESKey)
	}
	aesKey := setting.AESKey
	if aesKey[0] == 0 {
		switch gSetting.AESKeyType {
		case DefaultKey:
			aesKey = AESKeyDefault
		case CycleKey:
			aesKey = CycleAESKeys[(gSetting.MSVersion)%20]
		case CycleKey13:
			aesKey = CycleAESKeys[(gSetting.MSVersion+13)%20]
		default:
			panic("Unknown aes key type")
		}
	}
	gAESCipher = crypt.NewCAESCipher(aesKey, uint8(gSetting.AESInitType))
	gXORCipher = crypt.NewXORCipher()
	gLinearCipher = crypt.NewLinearCipher()
}

type CClientSocket interface {
	SetID(id int32)
	GetID() int32
	GetAddr() string
	XORRecvBuff()
	XORSendBuff()
	OnRead()
	OnConnect()
	InnoHash(cipherType CipherType, dwKey []byte)
	LoopAliveAck(aliveAckSec int)
	OnAliveAck()
	LoopAliveReq(aliveReqSec int, LP_AliveReq uint16)
	OnAliveReq(LP_AliveReq uint16)
	SetRecvCipherType(cipherType CipherType)
	SetSendCipherType(cipherType CipherType)
	OnOpcodeEncryption(LP_OpcodeEncryption uint16, startOpcode uint16, endOpcode uint16, isSplit bool)
	DecryptOpcode(randNum uint16) uint16
	SendPacket(oPacket COutPacket)
	Flush()
	OnError(err error)
	Close()
}

type CClientSocketDelegate interface {
	DebugInPacketLog(id int32, iPacket CInPacket)
	DebugOutPacketLog(id int32, oPacket COutPacket)
	NewConnectPacket(region Region, version uint16, minorVersion string, seqRcv []byte, seqSnd []byte) COutPacket
	ProcessPacket(cs CClientSocket, iPacket CInPacket)
	SocketClose(id int32)
}

type CInPacket interface {
	DecryptHeader(cipherType CipherType, pBuff []byte)
	DecryptData(cipherType CipherType, dwKey []byte)
	GetType() uint16
	GetRemain() int
	GetOffset() int
	GetLength() int
	DecodeBool() bool
	Decode1() int8
	Decode2() int16
	Decode4() int32
	Decode8() int64
	DecodeBuffer(uSize int) []byte
	DecodeStr() string
	DecodeName(uSize ...int) string
	DecodeTime() uint32
	DecodeDateTime() time.Time
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
	EncodeBuffer(buf []byte)
	EncodeStr(s string)
	EncodeName(s string, uSize ...int)
	EncodeTime(tTime uint32)
	EncodeDateTime(dTime time.Time)
	EncryptHeader(cipherType CipherType, pBuff []byte, dataLen int, dwKey []byte)
	MakeBufferList(cipherType CipherType, dwKey []byte) []byte
	DumpString(nSize int) string
}
