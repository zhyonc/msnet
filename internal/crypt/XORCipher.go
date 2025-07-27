package crypt

import "encoding/binary"

type XORCipher struct{}

func (static *XORCipher) Decrypt(buf []byte, pdwKey []byte) {
	inputKey := pdwKey[0]
	for i := range buf {
		b := inputKey ^ buf[i]
		buf[i] = (b << 4) | (b >> 4)
	}
}

func (static *XORCipher) Encrypt(buf []byte, pdwKey []byte) {
	inputKey := pdwKey[0]
	for i := range buf {
		x := (buf[i] << 4) | (buf[i] >> 4)
		buf[i] = inputKey ^ x
	}
}

func (static *XORCipher) Shuffle(pdwKey []byte) {
	// crtRand
	seed := binary.LittleEndian.Uint32(pdwKey)
	binary.LittleEndian.PutUint32(pdwKey, 214013*seed+2531011)
}
