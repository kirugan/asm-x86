package asmx86

import "fmt"

const (
	Mod00 = 0b00
	Mod01 = 0b01
	Mod10 = 0b10
	Mod11 = 0b11
)

const (
	Reg000 = 0b000
	Reg001 = 0b001
	Reg010 = 0b010
	Reg011 = 0b011
	Reg100 = 0b100
	Reg101 = 0b101
	Reg110 = 0b110
	Reg111 = 0b111
)

type modRegRM byte

func (m modRegRM) String() string {
	return fmt.Sprintf("%08b", m)
}

func (m modRegRM) Mod() int {
	return int(m & 0b11000000 >> 6)
}

func (m modRegRM) Reg() int {
	return int(m & 0b00111000 >> 3)
}

func (m modRegRM) RM() int {
	return int(m & 0b000000111)
}
