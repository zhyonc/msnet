package msnet

import (
	"fmt"
	"log/slog"

	"github.com/zhyonc/msnet/enum"
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
	p := &iPacket{}
	p.RecvBuff = buf
	p.Length = len(buf)
	return p
}

// AppendBuffer implements CInPacket.
func (p *iPacket) AppendBuffer(pBuff []byte, bEnc bool) {
	// Decode packet length
	p.RecvBuff = pBuff
	p.Length = len(pBuff)
	p.Offset = 0
	p.RawSeq = uint16(p.Decode2())
	temp := uint16(p.Decode2())
	if bEnc {
		temp ^= p.RawSeq
	}
	p.DataLen = int(temp)
}

// DecryptData implements CInPacket.
func (p *iPacket) DecryptData(dwKey []byte) {
	if p.Length <= 0 && p.Length > maxDataLength {
		slog.Warn("Invalid data length")
		return
	}
	if gSetting.IsCycleAESKey {
		(*crypt.CAESCipher).Decrypt(nil, crypt.CycleAESKeys[gSetting.MSVersion%20], p.RecvBuff, dwKey)
	} else {
		(*crypt.CAESCipher).Decrypt(nil, gSetting.AESKeyDecrypt, p.RecvBuff, dwKey)
	}
	if gSetting.MSRegion > enum.TMS || (gSetting.MSRegion == enum.CMS && gSetting.MSVersion < 86) {
		(*crypt.CIOBufferManipulator).De(nil, p.RecvBuff)
	}
}

// GetType implements CInPacket.
func (p *iPacket) GetType() int16 {
	if len(p.RecvBuff) >= 2 {
		return int16(p.RecvBuff[0]) | int16(p.RecvBuff[1])<<8
	}
	return 0
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
	p.Offset += 1
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

// DecodeFT implements CInPacket.
func (p *iPacket) DecodeFT() time.Time {
	// FileTime is in 100-nanosecond intervals
	// Convert to nanoseconds by multiplying by 100
	// FileTime epoch is January 1, 1601
	// Unix epoch is January 1, 1970
	// Calculate the difference between the two in nanoseconds
	ft := p.Decode8()
	nano := (ft - fileTimeEpochDiff) * 100
	return time.Unix(0, nano)
}

// DecodeStr implements CInPacket.
func (p *iPacket) DecodeStr() string {
	if p.GetRemain() < 2 {
		return ""
	}
	strLen := p.Decode2()
	if p.GetRemain() < int(strLen) {
		return ""
	}
	start := p.Offset
	end := p.Offset + int(strLen)
	str := string(p.RecvBuff[start:end])
	p.Offset = end
	return str
}

// DecodeLocalStr implements CInPacket.
func (p *iPacket) DecodeLocalStr() string {
	strLen := p.Decode2()
	buf := p.DecodeBuffer(int(strLen))
	return GetLangStr(buf)
}

// DecodeLocalName implements CInPacket.
func (p *iPacket) DecodeLocalName() string {
	buf := p.DecodeBuffer(13)
	return GetLangStr(buf)
}

// DecodeBuffer implements CInPacket.
func (p *iPacket) DecodeBuffer(uSize int) []byte {
	if p.GetRemain() < uSize {
		return nil
	}
	if uSize < 0 {
		uSize = 0
	}
	result := make([]byte, uSize)
	for i := range uSize {
		result[i] = byte(p.Decode1())
	}
	return result
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
		builder.WriteString(fmt.Sprintf("%02X", v))
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
