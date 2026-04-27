package crypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

type TripleDESCipher struct {
	key   []byte
	block cipher.Block
}

func NewTripleDESCipher(desKey string) (*TripleDESCipher, error) {
	finalKey := make([]byte, 24)
	if len(desKey) == 16 {
		copy(finalKey[0:16], desKey)
		copy(finalKey[16:24], desKey[0:8])
	} else if len(desKey) == 24 {
		copy(finalKey, desKey)
	} else {
		return nil, fmt.Errorf("des key length must be 16 or 24 byte")
	}
	temp, err := des.NewTripleDESCipher(finalKey)
	if err != nil {
		return nil, err
	}

	return &TripleDESCipher{
		key:   finalKey,
		block: temp,
	}, nil
}

func (c *TripleDESCipher) GetBlockSize() int32 {
	return int32(c.block.BlockSize())
}

func (c *TripleDESCipher) Encrypt(content string) ([]byte, error) {
	if c.block == nil {
		return nil, fmt.Errorf("cipher block is empty")
	}
	buf := []byte(content)
	blockSize := c.block.BlockSize()
	if len(buf)%blockSize != 0 {
		paddingLen := blockSize - (len(buf) % blockSize)
		paddingText := bytes.Repeat([]byte{' '}, paddingLen)
		buf = append(buf, paddingText...)
	}
	for i := 0; i < len(buf); i += blockSize {
		c.block.Encrypt(buf[i:i+blockSize], buf[i:i+blockSize])
	}
	return buf, nil
}

func (c *TripleDESCipher) Decrypt(buf []byte) (string, error) {
	if c.block == nil {
		return "", fmt.Errorf("cipher block is empty")
	}
	blockSize := c.block.BlockSize()
	if len(buf)%blockSize != 0 {
		paddingLen := blockSize - (len(buf) % blockSize)
		paddingText := bytes.Repeat([]byte{' '}, paddingLen)
		buf = append(buf, paddingText...)
	}

	for i := 0; i < len(buf); i += blockSize {
		c.block.Decrypt(buf[i:i+blockSize], buf[i:i+blockSize])
	}
	content := string(bytes.TrimRight(buf, " "))
	return content, nil
}
