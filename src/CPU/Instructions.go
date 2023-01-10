package CPU

import (
	"encoding/binary"
	"math/rand"

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

func SE(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	nn := opCode[1]
	if s.V[x] == nn {
		s.IncrementPC()
	}
}

func SNE(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	nn := opCode[1]
	if s.V[x] != nn {
		s.IncrementPC()
	}
}

func SEVxVy(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	if s.V[x] == s.V[y] {
		s.IncrementPC()
	}
}

func SNEVxVy(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	if s.V[x] == s.V[y] {
		s.IncrementPC()
	}
}

func LD(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	nn := opCode[1]
	s.V[x] = nn
}

func ADD(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	nn := int(opCode[1])
	s.V[x] = byte(int(s.V[x]) + nn)
}

func SETIR(s *State, opCode [2]byte) {
	s.IR[0] = byte(int(opCode[0]) & 0x0f)
	s.IR[1] = opCode[1]
}

func JUMPVx(s *State, opCode [2]byte) {
	//x := int(opCode[0]) & 0x0f
	//pc := nnn + s.V[x]

	nnn := int(binary.BigEndian.Uint16(opCode[:])) & 0x0fff
	pc := nnn + int(s.V[0])
	s.PC[0] = byte(pc >> 8)
	s.PC[1] = byte(pc & 0xff)
}

func CALL(s *State, opCode [2]byte) {
	nnn := binary.BigEndian.Uint16(opCode[:]) & 0x0fff
	sp := int(s.SP) + 1
	s.Stack[sp] = s.PC
	s.SP = byte(sp)
	s.PC[0] = byte((nnn >> 8) & 0x0f)
	s.PC[1] = byte(nnn & 0xff)
}

func LDVxVy(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	s.V[x] = byte(s.V[y])
}

func OR(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	s.V[x] = byte(s.V[x] | s.V[y])
}

func AND(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	s.V[x] = byte(s.V[x] & s.V[y])
}

func XOR(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	s.V[x] = byte(s.V[x] ^ s.V[y])
}

func ADDVxVy(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	sum := s.V[x] + s.V[y]
	if sum > 0xff {
		sum = sum & 0xff
		s.V[0xf] = byte(0x01)
	} else {
		s.V[0xf] = byte(0x00)
	}
	s.V[x] = byte(sum)
}

func SUBVxVy(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	diff := s.V[x] - s.V[y]
	if diff < 0x00 {
		diff = 0xff
		s.V[0xf] = byte(0x00)
	} else {
		s.V[0xf] = byte(0x01)
	}
	s.V[x] = byte(diff)
}

func SUBVyVx(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	diff := s.V[y] - s.V[x]
	if diff < 0x00 {
		diff = 0xff
		s.V[0xf] = byte(0x00)
	} else {
		s.V[0xf] = byte(0x01)
	}
	s.V[x] = byte(diff)
}

func SHR(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	//y := (int(opCode[1]) >> 4) & 0x0f
	//s.V[x] = s.V[y]
	s.V[0xf] = s.V[x] & 0x01
	s.V[x] = s.V[x] >> 1
}

func SHL(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	//y := (int(opCode[1]) >> 4) & 0x0f
	//s.V[x] = s.V[y]
	s.V[0xf] = s.V[x] & 0x80
	s.V[x] = s.V[x] << 1
}

func RND(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	rn := rand.Intn(0xff)
	nn := int(opCode[1])
	s.V[x] = byte(rn & nn)
}

func DISPLAY(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	y := (int(opCode[1]) >> 4) & 0x0f
	n := int(opCode[1]) & 0x0f
	vx := int(s.V[x])
	vy := int(s.V[y])
	loc := int(binary.BigEndian.Uint16(s.IR[:]))
	sprite := s.Memory[loc : loc+n]
	Display.Display(vx, vy, n, sprite)
}

func LDVxDT(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	s.V[x] = s.DelayTimer
}

//func LDVxK(s *State, opCode [2]byte) {
//    x := int(opCode[0]) & 0x0f
//    s.DecrementPC()
//    loop := true
//    for loop {
//        if KeyPressed {
//            loop = false
//            s.V[x] = byte(KeyPressed.val)
//        }
//    }
//}

func LDDTVx(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	s.DelayTimer = s.V[x]
}

func LDSTVx(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	s.SoundTimer = s.V[x]
}

func ADDI(s *State, opCode [2]byte) {
	I := int(binary.BigEndian.Uint16(s.IR[:]))
	x := int(opCode[0]) & 0x0f
	I = (I + x) & 0xffff
	s.IR[0] = byte(I >> 8)
	s.IR[1] = byte(I & 0xff)
}

//func SKP(s *State, opCode [2]byte) {
//    x := int(opCode[0]) & 0x0f
//    if KeyPressed == s.V[x] {
//        s.IncrementPC()
//    }
//}

//func SKNP(s *State, opCode [2]byte) {
//    x := int(opCode[0]) & 0x0f
//    if KeyPressed != s.V[x] {
//        s.IncrementPC()
//    }
//}

func LDFVx(s *State, opCode [2]byte) {
	x := int(opCode[0]) & 0x0f
	char := int(s.V[x]) & 0x0f
	addr := 80 + 5*(char)
	s.IR[0] = byte(addr >> 8)
	s.IR[1] = byte(addr & 0xff)
}

func LDBVx(s *State, opCode [2]byte) {
	IR := int(binary.BigEndian.Uint16(s.IR[:]))
	x := int(opCode[0]) & 0x0f
	n := int(s.V[x])
	s.Memory[IR] = byte(n / 100)
	s.Memory[IR+1] = byte(n / 10)
	s.Memory[IR+2] = byte(n % 10)
}

func LDIVx(s *State, opCode [2]byte) {
	IR := int(binary.BigEndian.Uint16(s.IR[:]))
	x := int(opCode[0]) & 0x0f
	for i := 0x0; i <= x; i++ {
		s.Memory[IR+i] = s.V[i]
	}
}

func LDVxI(s *State, opCode [2]byte) {
	IR := int(binary.BigEndian.Uint16(s.IR[:]))
	x := int(opCode[0]) & 0x0f
	for i := 0; i <= x; i++ {
		s.V[i] = s.Memory[IR+i]
	}
}
