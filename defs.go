package msnet

import "time"

const (
	LOG_BACKUP_DIR  string = "./log"
	SERVER_TYPE     string = "login"
	SERVER_ADDR     string = "127.0.0.1:8484"
	HEADER_LENGTH   int    = 4
	MAX_DATA_LENGTH int    = 1456
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

const FT_EPOCH_DIFF int64 = 116444736000000000 // (difference between 1601 and 1970 epochs)

var (
	ST_ZERO          time.Time                                                   // ST 0001-01-01 00:00:00 <-> FT -504911232000000000
	ST_ZERO_OVERFLOW time.Time = time.Date(1754, 8, 30, 22, 43, 41, 0, time.UTC) // ST 1754-08-30 22:43:41 <-> FT 48491090210000000
	ST_SQL_MIN       time.Time = time.Date(1753, 1, 1, 0, 0, 0, 0, time.UTC)     // ST 1753-01-01 00:00:00 <-> FT 47966688000000000
	ST_SQL_MIN_DATE  time.Time = time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)     // ST 1000-01-01 00:00:00 <-> FT -189657504000000000
	ST_START         time.Time = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)     // ST 1900-01-01 00:00:00 <-> FT 94354848000000000
	ST_END           time.Time = time.Date(2079, 1, 1, 0, 0, 0, 0, time.UTC)     // ST 2079-01-01 00:00:00 <-> FT 150842304000000000
	ST_EMPTY         time.Time = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)        // ST 0000-00-00 00:00:00 <-> FT -505255104000000000
)
