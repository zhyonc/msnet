package msnet

import "github.com/zhyonc/msnet/enum"

type Setting struct {
	MSRegion       enum.Region
	MSVersion      uint16
	MSMinorVersion string
	RecvXOR        uint8
	SendXOR        uint8
	IsXORCipher    bool
	IsCycleAESKey  bool
	AESKeyDecrypt  [32]byte
	AESKeyEncrypt  [32]byte
	AESInitType    enum.AESInitType
}
