package opcode_test

import (
	"testing"

	"github.com/zhyonc/msnet/internal/opcode"
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
	t.Parallel()
	opcode.GenOpcodeMap(cpPath, cpGenPath, cpMapName)
	opcode.GenOpcodeMap(lpPath, lpGenPath, lpMapName)
}
