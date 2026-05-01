package msnet

const (
	SERVER_TYPE     string = "login"
	SERVER_ADDR     string = "127.0.0.1:8484"
	LOG_BACKUP_DIR  string = "./log"
	HEADER_LENGTH   int    = 4
	MAX_DATA_LENGTH int    = 1456
	FT_EPOCH_DIFF   int64  = 116444736000000000 // FileTime epoch is January 1, 1601
	GMSCW_DES_KEY   string = "G0dD@mnN#H@ckEr!"
	KMS_DES_KEY     string = "G0dD@mnN#H@ckEr!"
	JMS_DES_KEY     string = "M@pl3J@p@nH@ck3r"
	CMS_DES_KEY     string = "aVbTpJ5=ZjG&Db3$"
	TMS_DES_KEY     string = "BrN=r54jQp2@yP6G"
	GMS_DES_KEY     string = "N3x@nGLEUH@ckEr!"
)

type Region uint8

const (
	GMSCW Region = 1
	KMS   Region = 1
	KMST  Region = 2
	JMS   Region = 3
	CMS   Region = 4
	CMST  Region = 5
	TMS   Region = 6
	MSEA  Region = 7
	GMS   Region = 8
	EMS   Region = 9
	BMS   Region = 9
)

type CipherType uint8

const (
	AESCipher CipherType = iota
	XORCipher
	LinearCipher
	NullCipher
)

type AESInitType uint8

const (
	Default AESInitType = iota
	Duplicate
	Shuffle
)
