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
	SetID(id int32)
	GetID() int32
	GetAddr() string
	XORRecv(buf []byte)
	XORSend(buf []byte)
	OnRead()
	OnConnect()
	OnAliveReq(LP_AliveReq uint16)
	OnMigrateCommand(LP_MigrateCommand uint16, ip string, port int16)
	SendPacket(oPacket COutPacket)
	Flush()
	OnError(err error)
	Close()
}

type CClientSocketDelegate interface {
	DebugInPacketLog(id int32, iPacket CInPacket)
	DebugOutPacketLog(id int32, oPacket COutPacket)
	ProcessPacket(cs CClientSocket, iPacket CInPacket)
	SocketClose(id int32)
}

type CInPacket interface {
	AppendBuffer(pBuff []byte, bEnc bool)
	DecryptData(dwKey []byte)
	GetType() uint16
	GetTypeByte() uint8
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
	GetTypeByte() uint8
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
