package CPU

import (
	"encoding/binary"

	"github.com/shivkar2n/Chip8-Emulator/Display"
)

func CLS(s *State) {
	Display.ClrScrn()
}

func RET(s *State) {
	s.PC = s.Stack[int(s.SP)]
	sp := int(s.SP) - 1
	s.SP = byte(sp)
}

func JUMP(s *State, opCode [2]byte) {
	s.PC[0] = byte(int(opCode[0]) & 0x0f)
	s.PC[1] = opCode[1]
}

func ADD(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	nn := int(opCode[1])
	s.V[x] = byte(int(s.V[x]) + nn)
}

func LD(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	nn := opCode[1]
	s.V[x] = nn
}

func SETIR(s *State, opCode [2]byte) {
	s.IR[0] = byte(int(opCode[0]) & 0x0f)
	s.IR[1] = opCode[1]
}

func DISPLAY(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := int((binary.BigEndian.Uint16(opCode[:]) >> 4) & 0x00f)
	n := int(binary.BigEndian.Uint16(opCode[:]) & 0x000f)
	vx := int(s.V[x])
	vy := int(s.V[y])
	loc := int(binary.BigEndian.Uint16(s.IR[:]))
	sprite := s.Memory[loc : loc+n]
	Display.Display(vx, vy, n, sprite)
}

func CALL(s *State, opCode [2]byte) {
	nnn := binary.BigEndian.Uint16(opCode[:]) & 0x0fff
	sp := int(s.SP) + 1
	s.Stack[sp] = s.PC
	s.SP = byte(sp)
	s.PC[0] = byte((nnn >> 8) & 0x0f)
	s.PC[1] = byte(nnn & 0xff)
}
