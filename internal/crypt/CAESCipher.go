package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"log/slog"
)

var (
	dwKeyDefault = [16]byte{
		0xF2, 0x53, 0x50, 0xC6, // -229420858
		0x7F, 0x9D, 0x42, 0xA8, // 2141012648
		0x26, 0x1D, 0x09, 0x77, // 639437175
		0x7C, 0x88, 0x53, 0x42, // 2089308994
	}
)

type CAESCipher struct {
	aesKey      [32]byte     // userKey
	aesInitType uint8        //
	chainVar    []byte       // 16 size IV
	block       cipher.Block // Use Block to generate RoundKey
}

func NewCAESCipher(aesKey [32]byte, aesInitType uint8) *CAESCipher {
	c := &CAESCipher{
		aesKey:      aesKey,
		aesInitType: aesInitType,
		chainVar:    make([]byte, aes.BlockSize),
	}
	block, err := aes.NewCipher(c.aesKey[:])
	if err != nil {
		panic(err)
	}
	c.block = block
	return c
}

// void __cdecl CAESCipher::Decrypt(unsigned __int8 *pDest, unsigned __int8 *pSrc, int nLen, unsigned int *pdwKey)
func (c *CAESCipher) Decrypt(buf []byte, dwkey []byte) {
	c.AES_Init(dwkey)
	if len(buf) > 0 {
		// The cipher.NewOFB stream handles the decryption in a single step using XORKeyStream
		c.OFB_Update(buf)
	}
	// No separate completion steps required.
	// OFB_DecFinal(AlgInfo, pDest)
}

// void __cdecl CAESCipher::Encrypt(unsigned __int8 *pDest, unsigned __int8 *pSrc, int nLen, unsigned int *pdwKey)
func (c *CAESCipher) Encrypt(buf []byte, dwkey []byte) {
	c.AES_Init(dwkey)
	if len(buf) > 0 {
		// The cipher.NewOFB stream handles the encryption in a single step using XORKeyStream
		c.OFB_Update(buf)
	}
	// No separate completion steps required.
	// OFB_EncFinal(AlgInfo, pDest)
}

// void __cdecl CAESCipher::AES_DecInit(CAESCipher::AES_ALG_INFO *AlgInfo, unsigned int *pdwKey)
// void __cdecl CAESCipher::AES_EncInit(CAESCipher::AES_ALG_INFO *AlgInfo, unsigned int *pdwKey)
func (c *CAESCipher) AES_Init(dwKey []byte) {
	if len(dwKey) > 0 {
		switch c.aesInitType {
		case 0:
			// Default: Used in versions after about 2008
			for i := range 4 {
				copy(c.chainVar[4*i:], dwKey)
			}
		case 1:
			// Duplicate: Used in versions about 2005~2007 (excluding TMS)
			for i := range c.chainVar {
				c.chainVar[i] = dwKey[0]
			}
		case 2:
			// Shuffle: Used in TMS versions about 2005~2007
			tempKey := make([]byte, 4)
			copy(tempKey, dwKey)
			for i := range 4 {
				(*CIGCipher).Shuffle(nil, tempKey, bShuffle[i])
				copy(c.chainVar[4*i:], tempKey)
			}
		default:
			slog.Warn("Invaild aes init type", "aesInitType", c.aesInitType)
			copy(c.chainVar, dwKeyDefault[:])
		}
	} else {
		// The default key is rarely used so i didn't test it
		copy(c.chainVar, dwKeyDefault[:])
	}
}

// char __cdecl CAESCipher::OFB_DecUpdate(CAESCipher::AES_ALG_INFO *AlgInfo,char *CipherTxt,unsigned int CipherTxtLen,char *PlainTxt,unsigned int *PlainTxtLen)
// char __cdecl CAESCipher::OFB_EncUpdate(CAESCipher::AES_ALG_INFO *AlgInfo,char *PlainTxt,unsigned int PlainTxtLen,char *CipherTxt,unsigned int *CipherTxtLen)
func (c *CAESCipher) OFB_Update(buf []byte) {
	stream := cipher.NewOFB(c.block, c.chainVar)
	stream.XORKeyStream(buf, buf)
}
