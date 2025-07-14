package msnet

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"github.com/zhyonc/msnet/internal/crypt"
	"github.com/zhyonc/msnet/internal/enum"
)

type oPacket struct {
	SendBuff            []byte
	IsEncryptedByShanda bool
}

func NewCOutPacket(nType int16) COutPacket {
	p := &oPacket{
		SendBuff:            make([]byte, 0),
		IsEncryptedByShanda: false,
	}
	p.Encode2(nType)
	return p
}

// GetType implements COutPacket
func (p *oPacket) GetType() int16 {
	if len(p.SendBuff) >= 2 {
		return int16(p.SendBuff[0]) | int16(p.SendBuff[1])<<8
	}
	return 0
}

// GetLength implements COutPacket.
func (p *oPacket) GetLength() int {
	return len(p.SendBuff)
}

// GetSendBuffer implements COutPacket
func (p *oPacket) GetSendBuffer() []byte {
	return p.SendBuff
}

// EncodeBool implements COutPacket
func (p *oPacket) EncodeBool(b bool) {
	var n byte
	if b {
		n = 1
	}
	p.SendBuff = append(p.SendBuff, n)
}

// Encode1 implements COutPacket
func (p *oPacket) Encode1(n int8) {
	p.SendBuff = append(p.SendBuff, byte(n))
}

// Encode2 implements COutPacket
func (p *oPacket) Encode2(n int16) {
	p.SendBuff = append(p.SendBuff, byte(n), byte(n>>8))
}

// Encode4 implements COutPacket
func (p *oPacket) Encode4(n int32) {
	buf := make([]byte, 4)
	for i := range 4 {
		buf[i] = byte(n >> (i * 8))
	}
	p.SendBuff = append(p.SendBuff, buf...)
}

// Encode8 implements COutPacket
func (p *oPacket) Encode8(n int64) {
	buf := make([]byte, 8)
	for i := range 8 {
		buf[i] = byte(n >> (i * 8))
	}
	p.SendBuff = append(p.SendBuff, buf...)
}

// EncodeFT implements COutPacket
func (p *oPacket) EncodeFT(t time.Time) {
	// Convert the time.Time value to nanoseconds since the Unix epoch
	nano := t.UnixNano() // nano=currentTime-8hours
	// Add the local time zone offset
	_, offset := t.Zone()
	offsetNano := int64(offset) * int64(time.Second)
	nano += offsetNano
	// Convert from nanoseconds to 100-nanosecond intervals (the unit used by FileTime)
	ft := nano / 100
	// Add the difference between the Unix and FileTime epochs
	ft += fileTimeEpochDiff
	p.Encode8(ft)
}

// EncodeStr implements COutPacket
func (p *oPacket) EncodeStr(s string) {
	buf := []byte(s) // ASCII Code
	p.Encode2(int16(len(buf)))
	p.SendBuff = append(p.SendBuff, buf...)
}

// EncodeLocalStr implements COutPacket
func (p *oPacket) EncodeLocalStr(s string) {
	buf := GetLangBuf(s)
	p.Encode2(int16(len(buf)))
	p.EncodeBuffer(buf)
}

// EncodeLocalName implements COutPacket
func (p *oPacket) EncodeLocalName(s string) {
	localeBuf := make([]byte, 13)
	buf := GetLangBuf(s)
	if len(buf) > 0 {
		copy(localeBuf, buf)
	}
	p.EncodeBuffer(localeBuf)
}

// EncodeBuffer implements COutPacket
func (p *oPacket) EncodeBuffer(buf []byte) {
	p.SendBuff = append(p.SendBuff, buf...)
}

// MakeBufferList implements COutPacket
func (p *oPacket) MakeBufferList(uSeqBase uint16, bEnc bool, dwKey []byte) []byte {
	headerLen := uint16(headerLength)
	dataLen := uint16(len(p.SendBuff))
	var bufferList []byte
	if bEnc {
		bufferList = make([]byte, headerLen+dataLen)
		copy(bufferList[headerLen:], p.SendBuff)
		// Encrypt packet header
		uSeqBaseN := ^uSeqBase
		HIWORD := binary.LittleEndian.Uint16(dwKey[2:4])
		uRawSeq := HIWORD ^ uSeqBaseN
		dataLen ^= uRawSeq
		// Put encrypted header into buffer list
		binary.LittleEndian.PutUint16(bufferList, uRawSeq)
		binary.LittleEndian.PutUint16(bufferList[2:4], dataLen)
		// IsEncryptedByShanda
		if gSetting.MSRegion > enum.TMS || (gSetting.MSRegion == enum.CMS && gSetting.MSVersion < 86) {
			(*crypt.CIOBufferManipulator).En(nil, bufferList[headerLen:])
			p.IsEncryptedByShanda = true
		}
		var aesKey [32]byte
		if gSetting.IsCycleAESKey {
			aesKey = crypt.CycleAESKeys[uSeqBase%20]
		} else {
			aesKey = gSetting.AESKeyEncrypt
		}
		// Encrypt packet data
		for i := 4; i < len(bufferList); i += maxDataLength {
			end := min(i+maxDataLength, len(bufferList))
			(*crypt.CAESCipher).Encrypt(nil, aesKey, bufferList[i:end], dwKey)
		}
	} else {
		// Encode packet header for CClientSocket::OnConnect
		bufferList = make([]byte, headerLen+dataLen-2)
		binary.LittleEndian.PutUint16(bufferList, dataLen)
		binary.LittleEndian.PutUint16(bufferList[2:4], uSeqBase)
		copy(bufferList[headerLen:], p.SendBuff[2:])
	}
	return bufferList
}

// DumpString implements COutPacket
func (p *oPacket) DumpString(nSize int) string {
	length := len(p.SendBuff)
	if nSize <= 0 || nSize > length {
		nSize = length
	}
	var builder strings.Builder
	for i := range nSize {
		v := p.SendBuff[i]
		builder.WriteString(fmt.Sprintf("%02X", v))
		if i < nSize-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}
