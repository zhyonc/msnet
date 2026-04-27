package msnet

import "github.com/zhyonc/msnet/def"

type Setting struct {
	MSRegion          def.Region
	MSVersion         uint16
	MSMinorVersion    string
	CipherType        def.CipherType
	DESKey            string
	IsCycleAESKey     bool
	AESKeyDecrypt     [32]byte
	AESKeyEncrypt     [32]byte
	RecvXOR           uint8
	SendXOR           uint8
	IsTypeHeader1Byte bool
	AESInitType       def.AESInitType
}
