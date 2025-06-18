package opcode

import (
	"testing"
)

const (
	cpPath    string = "./CP.go"
	cpGenPath string = "./CP_gen.go"
	cpMapName string = "CPMap"
	lpPath    string = "./LP.go"
	lpGenPath string = "./LP_gen.go"
	lpMapName string = "LPMap"
)

func TestGenOpcodeMap(t *testing.T) {
	GenOpcodeMap(cpPath, cpGenPath, cpMapName)
	GenOpcodeMap(lpPath, lpGenPath, lpMapName)
}
