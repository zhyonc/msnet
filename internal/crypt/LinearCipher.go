package crypt

type LinearCipher struct{}

func NewLinearCipher() *LinearCipher {
	return &LinearCipher{}
}

func (c *LinearCipher) Decrypt(buf []byte, dwKey []byte) {
	inputKey := dwKey[0]
	for i := range buf {
		buf[i] -= inputKey
	}
}

func (c *LinearCipher) Encrypt(buf []byte, dwKey []byte) {
	inputKey := dwKey[0]
	for i := range buf {
		buf[i] += inputKey
	}
}
