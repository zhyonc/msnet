package crypt

import "encoding/binary"

type XORCipher struct{}

func NewXORCipher() *XORCipher {
	return &XORCipher{}
}

func (c *XORCipher) Decrypt(buf []byte, dwKey []byte) {
	inputKey := dwKey[0]
	for i := range buf {
		b := inputKey ^ buf[i]
		buf[i] = (b << 4) | (b >> 4)
	}
}

func (c *XORCipher) Encrypt(buf []byte, dwKey []byte) {
	inputKey := dwKey[0]
	for i := range buf {
		x := (buf[i] << 4) | (buf[i] >> 4)
		buf[i] = inputKey ^ x
	}
}

func (c *XORCipher) Shuffle(dwKey []byte) {
	// crtRand
	seed := binary.LittleEndian.Uint32(dwKey)
	binary.LittleEndian.PutUint32(dwKey, 214013*seed+2531011)
}
