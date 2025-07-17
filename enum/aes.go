package enum

type AESInitType int8

const (
	Default AESInitType = iota
	Duplicate
	Shuffle
)
