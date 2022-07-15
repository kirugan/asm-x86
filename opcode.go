package asmx86

import "fmt"

type opcode byte

func (op opcode) Mnemonic() string {
	switch {
	case op <= 5:
		return "add"
	default:
		panic(fmt.Errorf("un"))
	}
}
