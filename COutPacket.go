package msnet

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/zhyonc/msnet/internal/crypt"
)

type oPacket struct {
	SendBuff []byte
	Offset   int
}

func NewCOutPacket(nType ...any) COutPacket {
	p := &oPacket{
		SendBuff: make([]byte, 0),
	}
	if len(nType) > 0 {
		switch v := nType[0].(type) {
		case uint8:
			p.Encode1(int8(v))
		case int8:
			p.Encode1(v)
		case uint16:
			p.Encode2(int16(v))
		case int16:
			p.Encode2(v)
		case int:
			if gSetting.IsTypeHeader1Byte {
				p.Encode1(int8(v))
			} else {
				p.Encode2(int16(v))
			}
		case []byte:
			p.EncodeBuffer(v)
		default:
			slog.Warn(fmt.Sprintf("unsupported opcode type: %T", v))
		}
	}
	return p
}

// GetType implements COutPacket
func (p *oPacket) GetType() uint16 {
	bufLen := len(p.SendBuff)
	switch {
	case gSetting.IsTypeHeader1Byte && bufLen >= 1:
		return uint16(p.SendBuff[0])
	case bufLen >= 2:
		return uint16(p.SendBuff[0]) | uint16(p.SendBuff[1])<<8
	default:
		return 0
	}
}

// GetOffset implements COutPacket.
func (p *oPacket) GetOffset() int {
	return p.Offset
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
	p.Offset++
}

// Encode1 implements COutPacket
func (p *oPacket) Encode1(n int8) {
	p.SendBuff = append(p.SendBuff, byte(n))
	p.Offset++
}

// Encode2 implements COutPacket
func (p *oPacket) Encode2(n int16) {
	p.SendBuff = append(p.SendBuff, byte(n), byte(n>>8))
	p.Offset += 2
}

// Encode4 implements COutPacket
func (p *oPacket) Encode4(n int32) {
	buf := make([]byte, 4)
	for i := range 4 {
		buf[i] = byte(n >> (i * 8))
	}
	p.SendBuff = append(p.SendBuff, buf...)
	p.Offset += 4
}

// Encode8 implements COutPacket
func (p *oPacket) Encode8(n int64) {
	buf := make([]byte, 8)
	for i := range 8 {
		buf[i] = byte(n >> (i * 8))
	}
	p.SendBuff = append(p.SendBuff, buf...)
	p.Offset += 8
}

// EncodeBuffer implements COutPacket
func (p *oPacket) EncodeBuffer(buf []byte) {
	p.SendBuff = append(p.SendBuff, buf...)
	p.Offset += len(buf)
}

// EncodeStr implements COutPacket
func (p *oPacket) EncodeStr(s string) {
	rawBuf := GetLocaleBuf(s)
	rawBufLen := len(rawBuf)
	p.Encode2(int16(rawBufLen))
	if rawBufLen > 0 {
		p.EncodeBuffer(rawBuf)
	}
}

// EncodeName implements COutPacket
func (p *oPacket) EncodeName(s string, uSize ...int) {
	nameLen := MAX_NAME_LENGTH
	if len(uSize) > 0 {
		nameLen = uSize[0]
	}
	buf := make([]byte, nameLen)
	rawBuf := GetLocaleBuf(s)
	copy(buf, rawBuf)
	p.EncodeBuffer(buf)
}

// EncodeTime implements [COutPacket].
func (p *oPacket) EncodeTime(tTime uint32) {
	cTime := uint32(time.Now().UnixMilli())
	if cTime >= tTime {
		p.EncodeBool(true)
		p.Encode4(int32(cTime - tTime))
	} else {
		p.EncodeBool(false)
		p.Encode4(int32(tTime - cTime))
	}
}

// EncodeDateTime implements COutPacket
func (p *oPacket) EncodeDateTime(dTime time.Time) {
	// FileTime is in 100-nanosecond intervals
	nano := dTime.UnixNano()
	// Divide by 100 to convert nanoseconds -> 100ns units
	// Add FT_EPOCH_DIFF
	ft := nano/100 + FT_EPOCH_DIFF
	p.Encode8(ft)
}

// EncryptHeader implements [COutPacket].
func (p *oPacket) EncryptHeader(cipherType CipherType, pBuff []byte, dataLen int, dwKey []byte) {
	uSeqBaseN := ^gSetting.MSVersion
	HIWORD := binary.LittleEndian.Uint16(dwKey[2:4])
	uRawSeq := HIWORD ^ uSeqBaseN
	temp := uint16(dataLen)
	if cipherType != XORCipher {
		// XORCipher didn't do this
		temp ^= uRawSeq
	}
	binary.LittleEndian.PutUint16(pBuff, uRawSeq)
	binary.LittleEndian.PutUint16(pBuff[2:4], temp)
}

// MakeBufferList implements COutPacket
func (p *oPacket) MakeBufferList(cipherType CipherType, dwKey []byte) []byte {
	dataLen := len(p.SendBuff)
	bufferList := make([]byte, HEADER_LENGTH+dataLen)
	copy(bufferList[HEADER_LENGTH:], p.SendBuff)
	switch cipherType {
	case AESCipher:
		// Encrypt packet header
		p.EncryptHeader(cipherType, bufferList, dataLen, dwKey)
		// IsEncryptedByShanda
		if gSetting.MSRegion > TMS || (gSetting.MSRegion == CMS && gSetting.MSVersion < 86) {
			(*crypt.CIOBufferManipulator).En(nil, bufferList[HEADER_LENGTH:])
		}
		// Encrypt packet data
		bufferListLen := len(bufferList)
		blockSize := HEADER_LENGTH + MAX_DATA_LENGTH
		// Encrypt First Block
		firstEnd := min(bufferListLen, blockSize)
		gAESCipher.Encrypt(bufferList[4:firstEnd], dwKey)
		// Encrypt Remain Block
		for i := firstEnd; i < bufferListLen; i += blockSize {
			remainEnd := min(i+blockSize, bufferListLen)
			gAESCipher.Encrypt(bufferList[i:remainEnd], dwKey)
		}
	case XORCipher:
		// Encrypt packet header
		p.EncryptHeader(cipherType, bufferList, dataLen, dwKey)
		// Encrypt packet data
		gXORCipher.Encrypt(bufferList[HEADER_LENGTH:], dwKey)
	case LinearCipher:
		// Encrypt packet header
		p.EncryptHeader(cipherType, bufferList, dataLen, dwKey)
		// Encrypt packet data
		gLinearCipher.Encrypt(bufferList[HEADER_LENGTH:], dwKey)
	case NullCipher:
		// Encode packet header for CClientSocket::OnConnect
		binary.LittleEndian.PutUint16(bufferList, uint16(dataLen+2)) // +2 for MSVersion
		binary.LittleEndian.PutUint16(bufferList[2:4], gSetting.MSVersion)
	default:
		slog.Warn("Unknown cipher type when MakeBufferList", "cipherType", cipherType)
		return nil
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
		fmt.Fprintf(&builder, "%02X", v)
		if i < nSize-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}
