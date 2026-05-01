package msnet

type Setting struct {
	MSRegion          Region
	MSVersion         uint16
	MSMinorVersion    string
	CipherType        CipherType
	DESKey            string
	IsCycleAESKey     bool
	AESKeyDecrypt     [32]byte
	AESKeyEncrypt     [32]byte
	RecvXOR           uint8
	SendXOR           uint8
	AliveAckMins      uint8
	IsTypeHeader1Byte bool
	AESInitType       AESInitType
}
