package msnet

import (
	"fmt"
	"log/slog"

	"github.com/zhyonc/msnet/internal/crypt"

	"strings"
	"time"
)

type iPacket struct {
	RawSeq   uint16
	DataLen  int
	RecvBuff []byte
	Length   int
	Offset   int
}

func NewCInPacket(buf []byte) CInPacket {
	p := &iPacket{
		RecvBuff: buf,
		Length:   len(buf),
	}
	return p
}

// CInPacket::AppendBuffer
// DecryptHeader implements CInPacket.
func (p *iPacket) DecryptHeader(cipherType CipherType, pBuff []byte) {
	// Decode packet length
	p.RecvBuff = pBuff
	p.Length = len(pBuff)
	p.Offset = 0
	p.RawSeq = uint16(p.Decode2())
	temp := uint16(p.Decode2())
	if cipherType != XORCipher {
		// XORCipher didn't do this
		temp ^= p.RawSeq
	}
	p.DataLen = int(temp)
}

// DecryptData implements CInPacket.
func (p *iPacket) DecryptData(cipherType CipherType, dwKey []byte) {
	if p.Length <= 0 && p.Length > MaxDataLength {
		slog.Warn("Invalid data length")
		return
	}
	switch cipherType {
	case AESCipher:
		// Decrypt packet data
		gAESCipher.Decrypt(p.RecvBuff, dwKey)
		// IsEncryptedByShanda
		if gSetting.MSRegion > TMS || (gSetting.MSRegion == CMS && gSetting.MSVersion < 86) {
			(*crypt.CIOBufferManipulator).De(nil, p.RecvBuff)
		}
	case XORCipher:
		gXORCipher.Decrypt(p.RecvBuff, dwKey)
	case LinearCipher:
		gLinearCipher.Decrypt(p.RecvBuff, dwKey)
	case NullCipher:
		// Nothing
	default:
		slog.Warn("Unknown cipher type when DecryptData", "cipherType", cipherType)
	}
}

// GetType implements CInPacket.
func (p *iPacket) GetType() uint16 {
	bufLen := len(p.RecvBuff)
	switch {
	case gSetting.IsTypeHeader1Byte && bufLen >= 1:
		return uint16(p.RecvBuff[0])
	case bufLen >= 2:
		return uint16(p.RecvBuff[0]) | uint16(p.RecvBuff[1])<<8
	default:
		return 0
	}
}

// GetRemain implements CInPacket.
func (p *iPacket) GetRemain() int {
	return p.Length - p.Offset
}

// GetOffset implements CInPacket.
func (p *iPacket) GetOffset() int {
	return p.Offset
}

// GetLength implements CInPacket.
func (p *iPacket) GetLength() int {
	return p.Length
}

// DecodeBool implements CInPacket.
func (p *iPacket) DecodeBool() bool {
	return p.Decode1() == 1
}

// Decode1 implements CInPacket.
func (p *iPacket) Decode1() int8 {
	if p.GetRemain() <= 0 {
		return 0
	}
	result := int8(p.RecvBuff[p.Offset])
	p.Offset++
	return result
}

// Decode2 implements CInPacket.
func (p *iPacket) Decode2() int16 {
	if p.GetRemain() < 2 {
		return 0
	}
	var result int16
	for i := range 2 {
		index := p.Offset + i
		result |= int16(p.RecvBuff[index]) << (i * 8)
	}
	p.Offset += 2
	return result
}

// Decode4 implements CInPacket.
func (p *iPacket) Decode4() int32 {
	if p.GetRemain() < 4 {
		return 0
	}
	var result int32
	for i := range 4 {
		index := p.Offset + i
		result |= int32(p.RecvBuff[index]) << (i * 8)
	}
	p.Offset += 4
	return result
}

// Decode8 implements CInPacket.
func (p *iPacket) Decode8() int64 {
	if p.GetRemain() < 8 {
		return 0
	}
	var result int64
	for i := range 8 {
		index := p.Offset + i
		result |= int64(p.RecvBuff[index]) << (i * 8)
	}
	p.Offset += 8
	return result
}

// DecodeBuffer implements CInPacket.
func (p *iPacket) DecodeBuffer(uSize int) []byte {
	remain := p.GetRemain()
	if uSize <= 0 || remain <= 0 {
		return nil
	}
	if remain < uSize {
		uSize = remain
	}
	buf := p.RecvBuff[p.Offset : p.Offset+uSize]
	p.Offset += uSize
	return buf
}

// DecodeStr implements CInPacket.
func (p *iPacket) DecodeStr() string {
	if p.GetRemain() < 2 {
		return ""
	}
	strLen := int(p.Decode2())
	if strLen <= 0 {
		return ""
	}
	if p.GetRemain() < strLen {
		strLen = p.GetRemain()
	}
	rawBuf := p.DecodeBuffer(strLen)
	return GetLocaleStr(rawBuf)
}

// DecodeName implements CInPacket.
func (p *iPacket) DecodeName(uSize ...int) string {
	nameLen := MaxNameLength
	if len(uSize) > 0 {
		nameLen = uSize[0]
	}
	if p.GetRemain() < nameLen {
		nameLen = p.GetRemain()
	}
	rawBuf := p.DecodeBuffer(nameLen)
	return GetLocaleStr(rawBuf)
}

// DecodeTime implements [CInPacket].
func (p *iPacket) DecodeTime() uint32 {
	cTime := uint32(time.Now().UnixMilli())
	isPast := p.DecodeBool()
	offset := uint32(p.Decode4())
	if isPast {
		return cTime - offset
	}
	return cTime + offset
}

// DecodeDateTime implements CInPacket.
func (p *iPacket) DecodeDateTime() time.Time {
	// FileTime is in 100-nanosecond intervals
	ft := p.Decode8()
	if ft < FTEpochDiff {
		return time.Unix(0, 0)
	}
	// Subtract FT_EPOCH_DIFF
	// Multiply by 100 to convert 100ns units -> nanoseconds
	nano := (ft - FTEpochDiff) * 100
	return time.Unix(0, nano)
}

// DumpString implements CInPacket.
func (p *iPacket) DumpString(nSize int) string {
	length := len(p.RecvBuff)
	if nSize <= 0 || nSize > length {
		nSize = length
	}
	var builder strings.Builder
	for i := range nSize {
		v := p.RecvBuff[i]
		fmt.Fprintf(&builder, "%02X", v)
		if i < nSize-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

// Clear implements CInPacket.
func (p *iPacket) Clear() {
	p.Length = 0
	p.Offset = 0
	p.RecvBuff = p.RecvBuff[:0]
}
