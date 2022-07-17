package asmx86

var registers = [...]string{
	// data size is 8-bits
	"al", "cl", "dl", "bl",
	"ah", "ch", "dh", "bh",
	// data size is 16-bits
	"ax", "cx", "dx", "bx",
	"sp", "bp", "si", "di",
	// data size is 32-bits
	"eax", "ecx", "edx", "ebx",
	"esp", "ebp", "esi", "edi",
}

// todo range checks and maybe change regValue type to byte

func reg8(regValue int) string {
	return registers[regValue]
}

func reg16(regValue int) string {
	return registers[regValue+8]
}

func reg32(regValue int) string {
	return registers[regValue+16]
}
