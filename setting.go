package msnet

type Setting struct {
	MSRegion       uint8
	MSVersion      uint16
	MSMinorVersion string
	IsCycleAESKey  bool
	AESKeyDecrypt  [32]byte
	AESKeyEncrypt  [32]byte
}
