package CPU

import (
	"encoding/binary"
	"fmt"
	_ "fmt"
	"math/rand"

	"github.com/shivkar2n/Chip8-Emulator/Display"
	"github.com/shivkar2n/Chip8-Emulator/helpers"
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
	x := helpers.ParseOpcode(opCode)["x"]
	s.PC[0] = byte(x)
	s.PC[1] = opCode[1]
}

func SE(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	nn := byte(helpers.ParseOpcode(opCode)["nn"])
	if s.V[x] == nn {
		s.IncrementPC()
	}
}

func SNE(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	nn := byte(helpers.ParseOpcode(opCode)["nn"])
	if s.V[x] != nn {
		s.IncrementPC()
	}
}

func SEVxVy(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	if s.V[x] == s.V[y] {
		s.IncrementPC()
	}
}

func SNEVxVy(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	if s.V[x] != s.V[y] {
		s.IncrementPC()
	}
}

func LD(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	nn := byte(helpers.ParseOpcode(opCode)["nn"])
	s.V[x] = nn
}

func ADD(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	nn := helpers.ParseOpcode(opCode)["nn"]
	s.V[x] = byte(int(s.V[x]) + nn)
}

func SETIR(s *State, opCode [2]byte) {
	s.IR[0] = byte(int(opCode[0]) & 0x0f)
	s.IR[1] = opCode[1]
}

func JUMPVx(s *State, opCode [2]byte) {
	//x := helpers.ParseOpcode(opCode)["x"]
	//pc := nnn + s.V[x]

	nnn := helpers.ParseOpcode(opCode)["nnn"]
	pc := nnn + int(s.V[0])
	s.PC[0] = byte(pc >> 8)
	s.PC[1] = byte(pc & 0xff)
}

func CALL(s *State, opCode [2]byte) {
	nnn := helpers.ParseOpcode(opCode)["nnn"]
	sp := int(s.SP) + 1
	s.Stack[sp] = s.PC
	s.SP = byte(sp)
	s.PC[0] = byte((nnn >> 8) & 0x0f)
	s.PC[1] = byte(nnn & 0xff)
}

func LDVxVy(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	s.V[x] = byte(s.V[y])
}

func OR(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	s.V[x] = byte(int(s.V[x]) | int(s.V[y]))
}

func AND(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	s.V[x] = byte(int(s.V[x]) & int(s.V[y]))
}

func XOR(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	s.V[x] = byte(int(s.V[x]) ^ int(s.V[y]))
}

func ADDVxVy(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	sum := int(s.V[x]) + int(s.V[y])
	if sum > 0xff {
		sum = sum & 0xff
		s.V[0xf] = byte(0x01)
	} else {
		s.V[0xf] = byte(0x00)
	}
	s.V[x] = byte(sum)
}

func SUBVxVy(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	diff := int(s.V[x]) - int(s.V[y])
	if diff < 0x00 {
		s.V[0xf] = byte(0x00)
	} else {
		s.V[0xf] = byte(0x01)
	}
	s.V[x] = byte(diff)
}

func SUBVyVx(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	diff := int(s.V[y]) - int(s.V[x])
	if diff < 0x00 {
		diff = 0xff
		s.V[0xf] = byte(0x00)
	} else {
		s.V[0xf] = byte(0x01)
	}
	s.V[x] = byte(diff)
}

func SHR(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	if int(s.V[x])&0x01 == 0x01 {
		s.V[0xf] = byte(0x01)
	} else {
		s.V[0xf] = byte(0x00)
	}
	s.V[0xf] = s.V[x] & 0x01
	s.V[x] = s.V[x] >> 1
}

func SHL(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	if int(s.V[x])&0x80 == 0x80 {
		s.V[0xf] = byte(0x01)
	} else {
		s.V[0xf] = byte(0x00)
	}
	s.V[x] = s.V[x] << 1
}

func RND(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	rn := rand.Intn(0xff)
	nn := helpers.ParseOpcode(opCode)["nn"]
	s.V[x] = byte(rn & nn)
}

func DISPLAY(s *State, opCode [2]byte) {
	var res bool = false
	x := helpers.ParseOpcode(opCode)["x"]
	y := helpers.ParseOpcode(opCode)["y"]
	n := helpers.ParseOpcode(opCode)["n"]
	vx := int(s.V[x])
	vy := int(s.V[y])
	loc := int(binary.BigEndian.Uint16(s.IR[:]))
	sprite := s.Memory[loc : loc+n]
	Display.Display(vx, vy, n, sprite, &res)
	if res {
		s.V[0xf] = byte(0x01)
	} else {
		s.V[0xf] = byte(0x00)
	}
}

func LDVxDT(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	s.V[x] = s.DelayTimer
}

func LDVxK(s *State, opCode [2]byte, keyboardState []uint8) {
	x := helpers.ParseOpcode(opCode)["x"]
	s.DecrementPC()
	loop := true
	for loop {
		val := helpers.KeyPressed(keyboardState)
		if val != byte(0xff) {
			loop = false
			s.V[x] = val

		}
	}
}

func LDDTVx(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	s.DelayTimer = s.V[x]
}

func LDSTVx(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	s.SoundTimer = s.V[x]
}

func ADDI(s *State, opCode [2]byte) {
	I := int(binary.BigEndian.Uint16(s.IR[:]))
	x := helpers.ParseOpcode(opCode)["x"]
	I = (I + int(s.V[x])) & 0xffff
	s.IR[0] = byte(I >> 8)
	s.IR[1] = byte(I & 0xff)
}

func SKP(s *State, opCode [2]byte, keyboardState []uint8) {
	x := helpers.ParseOpcode(opCode)["x"]
	if helpers.KeyPressed(keyboardState) == byte(0xff) {
		fmt.Println("Nothing pressed!")
		return
	} else if helpers.KeyPressed(keyboardState) == s.V[x] {
		fmt.Println(s.V[x], " pressed!")
		s.IncrementPC()
	}
}

func SKNP(s *State, opCode [2]byte, keyboardState []uint8) {
	x := helpers.ParseOpcode(opCode)["x"]
	if helpers.KeyPressed(keyboardState) != s.V[x] {
		s.IncrementPC()
	}
}

func LDFVx(s *State, opCode [2]byte) {
	x := helpers.ParseOpcode(opCode)["x"]
	char := int(s.V[x]) & 0x0f
	addr := 80 + 5*(char)
	s.IR[0] = byte(addr >> 8)
	s.IR[1] = byte(addr & 0xff)
}

func LDBVx(s *State, opCode [2]byte) {
	IR := int(binary.BigEndian.Uint16(s.IR[:]))
	x := helpers.ParseOpcode(opCode)["x"]
	num := int(s.V[x])
	s.Memory[IR] = byte(num / 100)
	s.Memory[IR+1] = byte((num / 10) % 10)
	s.Memory[IR+2] = byte(num % 10)
}

func LDIVx(s *State, opCode [2]byte) {
	IR := int(binary.BigEndian.Uint16(s.IR[:]))
	x := helpers.ParseOpcode(opCode)["x"]
	for i := 0x0; i <= x; i++ {
		s.Memory[IR+i] = s.V[i]
	}
}

func LDVxI(s *State, opCode [2]byte) {
	IR := int(binary.BigEndian.Uint16(s.IR[:]))
	x := helpers.ParseOpcode(opCode)["x"]
	for i := 0; i <= x; i++ {
		s.V[i] = s.Memory[IR+i]
	}
}
