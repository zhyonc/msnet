package msnet

type Setting struct {
	MSRegion       uint8
	MSVersion      uint16
	MSMinorVersion string
	AESKeyDecrypt  [32]byte
	AESKeyEncrypt  [32]byte
}
