package asmx86

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		input  []byte
		output string
	}{
		{
			input:  []byte{0x0, 0xc0},
			output: "add al, al",
		},
		{
			input:  []byte{0x0, 0xc9},
			output: "add cl, cl",
		},
		{
			input:  []byte{0x0, 0xd2},
			output: "add dl, dl",
		},
		{
			input:  []byte{0x0, 0xdb},
			output: "add bl, bl",
		},
		{
			input:  []byte{0x0, 0xe4},
			output: "add ah, ah",
		},
		{
			input:  []byte{0x0, 0xed},
			output: "add ch, ch",
		},
		{
			input:  []byte{0x0, 0xf6},
			output: "add dh, dh",
		},
		{
			input:  []byte{0x0, 0xff},
			output: "add bh, bh",
		},
		{
			input:  []byte{0x3, 0x7},
			output: "...",
		},
	}

	for _, cse := range cases {
		result := disasm(cse.input)

		assert.Equal(t, cse.output, result)
	}
}

func TestInt(t *testing.T) {
	cases := []struct {
		input  []byte
		output string
	}{
		{
			input:  []byte{0xcd, 0x80},
			output: "int 0x80",
		},
	}

	for _, cse := range cases {
		result := disasm(cse.input)

		assert.Equal(t, cse.output, result)
	}
}

func disasm(in []byte) string {
	if len(in) != 2 {
		panic("only input with 2 bytes supported for now")
	}

	firstByte := opcode(in[0])
	secondByte := modRegRM(in[1])

	switch secondByte.Mod() {
	case Mod11, Mod00:
	// ok
	default:
		panic(fmt.Errorf("only register addressing mode supported for now (firstByte=%s)", secondByte))
	}

	regName := reg8
	if firstByte.Operands32() {
		regName = reg32
	}

	var operands []string

	reg := secondByte.Reg()
	regValue := regName(reg)
	operands = append(operands, regValue)

	rm := secondByte.RM()
	rmValue := regName(rm)
	operands = append(operands, rmValue)

	operandsStr := strings.Join(operands, ", ")
	if operandsStr != "" {
		// prepend with whitespace if there is some meaningful
		operandsStr = " " + operandsStr
	}

	return firstByte.Mnemonic() + operandsStr
}
