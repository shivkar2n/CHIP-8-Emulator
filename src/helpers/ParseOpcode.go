package helpers

import (
	"encoding/binary"
)

// Parse opcode and get arguements
func ParseOpcode(opCode [2]byte) map[string]int {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	n := int(opCode[1]) & 0x0f
	nn := int(opCode[1])
	nnn := int(binary.BigEndian.Uint16(opCode[:]) & 0x0fff)
	return map[string]int{"x": x, "y": y, "n": n, "nn": nn, "nnn": nnn}
}
