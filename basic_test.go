package asmx86

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMe(t *testing.T) {
	cases := []struct {
		input  []byte
		output string
	}{
		{
			input:  []byte{0x0, 0xc0},
			output: "add al, al",
		},
	}

	for _, cse := range cases {
		result := disasm(cse.input)

		assert.Equal(t, cse.output, result)
	}
}

func disasm(in []byte) string {
	firstByte := opcode(in[0])
	secondByte := modRegRM(in[1])

	if secondByte.Mod() != Mod11 {
		panic(fmt.Errorf("only register addressing mode supported for now (%s)", secondByte))
	}

	registers := [...]string{
		"al", "cl", "dl", "bl",
		"ah", "ch", "dh", "bh",
	}

	reg1 := secondByte.Reg()
	reg2 := secondByte.RM()
	reg1Name, reg2name := registers[reg1], registers[reg2]

	return fmt.Sprintf(
		"%s %s, %s",
		firstByte.Mnemonic(),
		reg1Name,
		reg2name,
	)
}
