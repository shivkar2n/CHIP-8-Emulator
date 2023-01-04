package CPU

import (
	"encoding/binary"
	//"fmt"
)

func (s *State) InstructionFetch() [2]byte { // Read bytes slice from CPU memory
	pc := int(binary.BigEndian.Uint16(s.PC[:])) & 0xffff
	buffer := [2]byte{}
	for i := 0; i < INSTRUCTION_SIZE; i++ {
		buffer[i] = s.Memory[pc+i]
	}
	return buffer
}
