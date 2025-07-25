package msnet

import (
	"time"

	"github.com/zhyonc/msnet/enum"
	"github.com/zhyonc/msnet/internal/crypt"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

const (
	headerLength      int   = 4
	maxDataLength     int   = 1456
	fileTimeEpochDiff int64 = 116444736000000000 // FileTime epoch is January 1, 1601
)

var (
	gSetting *Setting
)

func New(setting *Setting) {
	// Setting
	gSetting = setting
	// Language coder
	switch gSetting.MSRegion {
	case enum.CMS:
		langEncoder = simplifiedchinese.GBK.NewEncoder()
		langDecoder = simplifiedchinese.GBK.NewDecoder()
	case enum.TMS:
		langEncoder = traditionalchinese.Big5.NewEncoder()
		langDecoder = traditionalchinese.Big5.NewDecoder()
	default:
		langEncoder = encoding.Nop.NewEncoder()
		langDecoder = encoding.Nop.NewDecoder()
	}
	// AESInitType
	crypt.AESInitType = gSetting.AESInitType
}

type CClientSocket interface {
	OnMigrateCommand(LP_MigrateCommand int16, ip string, port int16)
	OnConnect()
	Flush()
	OnAliveReq(LP_AliveReq int16)
	XORRecv(buf []byte)
	XORSend(buf []byte)
	OnRead()
	SendPacket(oPacket COutPacket)
	OnError(err error)
	Close()
	GetAddr() string
}

type CClientSocketDelegate interface {
	DebugInPacketLog(iPacket CInPacket)
	DebugOutPacketLog(oPacket COutPacket)
	ProcessPacket(cs CClientSocket, iPacket CInPacket)
	SocketClose()
}

type CInPacket interface {
	AppendBuffer(pBuff []byte, bEnc bool)
	DecryptData(dwKey []byte)
	GetType() int16
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
	GetType() int16
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
	MakeBufferList(uSeqBase uint16, bEnc bool, dwKey []byte) []byte
	DumpString(nSize int) string
}
