package crypt

import (
	"crypto/aes"
	"crypto/cipher"
)

var (
	userKeyDefault = [32]byte{
		0x13, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00,
		0x06, 0x00, 0x00, 0x00,
		0xB4, 0x00, 0x00, 0x00,
		0x1B, 0x00, 0x00, 0x00,
		0x0F, 0x00, 0x00, 0x00,
		0x33, 0x00, 0x00, 0x00,
		0x52, 0x00, 0x00, 0x00,
	}

	pdwKeyDefault = [16]byte{
		0xF2, 0x53, 0x50, 0xC6, // -229420858
		0x7F, 0x9D, 0x42, 0xA8, // 2141012648
		0x26, 0x1D, 0x09, 0x77, // 639437175
		0x7C, 0x88, 0x53, 0x42, // 2089308994
	}
)

type AES_ALG_INFO struct {
	ChainVar []byte       // IV
	Block    cipher.Block // Use Block to generate RoundKey
}

type CAESCipher struct{}

// void __cdecl CAESCipher::Decrypt(unsigned __int8 *pDest, unsigned __int8 *pSrc, int nLen, unsigned int *pdwKey)
func (static *CAESCipher) Decrypt(userKey [32]byte, buf []byte, pdwkey []byte) {
	AlgInfo := &AES_ALG_INFO{
		ChainVar: make([]byte, aes.BlockSize),
	}
	RIJNDAEL_KeySchedule(userKey[:], AlgInfo)
	AES_Init(AlgInfo, pdwkey)
	if len(buf) > 0 {
		// The cipher.NewOFB stream handles the decryption in a single step using XORKeyStream
		OFB_Update(AlgInfo, buf)
	}
	// No separate completion steps required.
	// OFB_DecFinal(AlgInfo, pDest)
}

// void __cdecl CAESCipher::Encrypt(unsigned __int8 *pDest, unsigned __int8 *pSrc, int nLen, unsigned int *pdwKey)
func (static *CAESCipher) Encrypt(userKey [32]byte, buf []byte, pdwkey []byte) {
	AlgInfo := &AES_ALG_INFO{
		ChainVar: make([]byte, aes.BlockSize),
	}
	RIJNDAEL_KeySchedule(userKey[:], AlgInfo)
	AES_Init(AlgInfo, pdwkey)
	if len(buf) > 0 {
		// The cipher.NewOFB stream handles the encryption in a single step using XORKeyStream
		OFB_Update(AlgInfo, buf)
	}
	// No separate completion steps required.
	// OFB_EncFinal(AlgInfo, pDest)
}

// void __cdecl CAESCipher::RIJNDAEL_KeySchedule(unsigned int *UserKey, unsigned int *e_key)
func RIJNDAEL_KeySchedule(userKey []byte, info *AES_ALG_INFO) {
	if userKey[0] == 0 {
		userKey = userKeyDefault[:]
	}
	block, err := aes.NewCipher(userKey)
	if err != nil {
		panic(err)
	}
	info.Block = block
}

// void __cdecl CAESCipher::AES_DecInit(CAESCipher::AES_ALG_INFO *AlgInfo, unsigned int *pdwKey)
// void __cdecl CAESCipher::AES_EncInit(CAESCipher::AES_ALG_INFO *AlgInfo, unsigned int *pdwKey)
func AES_Init(info *AES_ALG_INFO, pdwKey []byte) {
	if len(pdwKey) > 0 {
		for i := range 4 {
			copy(info.ChainVar[4*i:], pdwKey[:])
		}
	} else {
		// The default key is rarely used so i didn't test it
		copy(info.ChainVar, pdwKeyDefault[:])
	}
}

// char __cdecl CAESCipher::OFB_DecUpdate(CAESCipher::AES_ALG_INFO *AlgInfo,char *CipherTxt,unsigned int CipherTxtLen,char *PlainTxt,unsigned int *PlainTxtLen)
// char __cdecl CAESCipher::OFB_EncUpdate(CAESCipher::AES_ALG_INFO *AlgInfo,char *PlainTxt,unsigned int PlainTxtLen,char *CipherTxt,unsigned int *CipherTxtLen)
func OFB_Update(info *AES_ALG_INFO, buf []byte) {
	stream := cipher.NewOFB(info.Block, info.ChainVar)
	stream.XORKeyStream(buf, buf)
}
