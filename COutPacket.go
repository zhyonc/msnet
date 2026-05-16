package msnet

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"github.com/zhyonc/msnet/internal/crypt"
)

type oPacket struct {
	SendBuff            []byte
	Offset              int
	IsEncryptedByShanda bool
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
			panic(fmt.Sprintf("unsupported opcode type: %T", v))
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

// EncodeFT implements COutPacket
func (p *oPacket) EncodeFT(t time.Time) {
	// FileTime is in 100-nanosecond intervals
	nano := t.UnixNano()
	// Divide by 100 to convert nanoseconds -> 100ns units
	// Add FT_EPOCH_DIFF
	ft := nano/100 + FT_EPOCH_DIFF
	p.Encode8(ft)
}

// EncodeStr implements COutPacket
func (p *oPacket) EncodeStr(s string) {
	buf := []byte(s) // ASCII Code
	bufLen := len(buf)
	p.Encode2(int16(bufLen))
	p.SendBuff = append(p.SendBuff, buf...)
	p.Offset += bufLen
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
	bufLen := len(buf)
	if bufLen > 0 {
		copy(localeBuf, buf)
	}
	p.EncodeBuffer(localeBuf)
	p.Offset += bufLen
}

// EncodeBuffer implements COutPacket
func (p *oPacket) EncodeBuffer(buf []byte) {
	p.SendBuff = append(p.SendBuff, buf...)
	p.Offset += len(buf)
}

// EncryptHeader implements [COutPacket].
func (p *oPacket) EncryptHeader(pBuff []byte, dataLen int, dwKey []byte) {
	uSeqBaseN := ^gSetting.MSVersion
	HIWORD := binary.LittleEndian.Uint16(dwKey[2:4])
	uRawSeq := HIWORD ^ uSeqBaseN
	temp := uint16(dataLen)
	if gSetting.CipherType != XORCipher {
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
		p.EncryptHeader(bufferList, dataLen, dwKey)
		// IsEncryptedByShanda
		if gSetting.MSRegion > TMS || (gSetting.MSRegion == CMS && gSetting.MSVersion < 86) {
			(*crypt.CIOBufferManipulator).En(nil, bufferList[HEADER_LENGTH:])
			p.IsEncryptedByShanda = true
		}
		// Switch AESKey
		var aesKey [32]byte
		if gSetting.IsCycleAESKey {
			var version int = int(gSetting.MSVersion)
			if gSetting.MSRegion == KMS || gSetting.MSRegion == KMS && version >= 1112 || gSetting.MSRegion == JMS && version >= 300 {
				version += 13
			}
			aesKey = crypt.CycleAESKeys[version%20]
		} else {
			aesKey = gSetting.AESKeyEncrypt
		}
		// Encrypt packet data
		bufferListLen := len(bufferList)
		blockSize := HEADER_LENGTH + MAX_DATA_LENGTH
		// Encrypt First Block
		firstEnd := min(bufferListLen, blockSize)
		(*crypt.CAESCipher).Encrypt(nil, aesKey, bufferList[4:firstEnd], dwKey)
		// Encrypt Remain Block
		for i := firstEnd; i < bufferListLen; i += blockSize {
			remainEnd := min(i+blockSize, bufferListLen)
			(*crypt.CAESCipher).Encrypt(nil, aesKey, bufferList[i:remainEnd], dwKey)
		}
	case XORCipher:
		// Encrypt packet header
		p.EncryptHeader(bufferList, dataLen, dwKey)
		// Encrypt packet data
		(*crypt.XORCipher).Encrypt(nil, bufferList[HEADER_LENGTH:], dwKey)
	case LinearCipher:
		// Encrypt packet header
		p.EncryptHeader(bufferList, dataLen, dwKey)
		// Encrypt packet data
		key := dwKey[0]
		for i := HEADER_LENGTH; i < len(bufferList); i++ {
			bufferList[i] += key
		}
	case NullCipher:
		// Encode packet header for CClientSocket::OnConnect
		binary.LittleEndian.PutUint16(bufferList, uint16(dataLen+2)) // +2 for MSVersion
		binary.LittleEndian.PutUint16(bufferList[2:4], gSetting.MSVersion)
	default:
		panic("Unknown cipher type")
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
