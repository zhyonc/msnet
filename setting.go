package msnet

type Setting struct {
	MSRegion       uint8
	MSVersion      uint16
	MSMinorVersion string
	RecvXOR        uint8
	SendXOR        uint8
	IsCycleAESKey  bool
	AESKeyDecrypt  [32]byte
	AESKeyEncrypt  [32]byte
}
