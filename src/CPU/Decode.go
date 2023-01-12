package CPU

//package main

import (
	"encoding/binary"
	"fmt"
	_ "fmt"
	//"os"
	//"github.com/shivkar2n/Chip8-Emulator/helpers"
)

func (s *State) InstructionDecode(opCodeSlice [2]byte, keyboardState []uint8) {
	opcodeBytes := binary.BigEndian.Uint16(opCodeSlice[:])

	n := int(opcodeBytes) & 0x000f
	nn := int(opcodeBytes) & 0x00ff
	nnn := int(opcodeBytes) & 0x0fff
	//x := int(opcodeBytes) & 0x0f00 >> 8
	//y := int(opcodeBytes) & 0x00f0 >> 4
	opCode0 := int(opcodeBytes) & 0xf000
	opCode := int(opcodeBytes)

	switch opCode0 {
	case 0x0000:
		if opCode == 0x00e0 {
			CLS(s)
		} else if opCode == 0x00ee {
			RET(s)
		} else {
			fmt.Printf("SYS %x\n", nnn)
		}

	case 0x1000:
		JUMP(s, opCodeSlice)

	case 0x2000:
		CALL(s, opCodeSlice)

	case 0x3000:
		SE(s, opCodeSlice)

	case 0x4000:
		SNE(s, opCodeSlice)

	case 0x5000:
		SEVxVy(s, opCodeSlice)

	case 0x6000:
		LD(s, opCodeSlice)

	case 0x7000:
		ADD(s, opCodeSlice)

	case 0x8000:
		operandCode := n
		if operandCode == 0x0 {
			LDVxVy(s, opCodeSlice)

		} else if operandCode == 0x1 {
			OR(s, opCodeSlice)

		} else if operandCode == 0x2 {
			AND(s, opCodeSlice)

		} else if operandCode == 0x3 {
			XOR(s, opCodeSlice)

		} else if operandCode == 0x4 {
			ADDVxVy(s, opCodeSlice)

		} else if operandCode == 0x5 {
			SUBVxVy(s, opCodeSlice)

		} else if operandCode == 0x6 {
			SHR(s, opCodeSlice)

		} else if operandCode == 0x7 {
			SUBVyVx(s, opCodeSlice)

		} else if operandCode == 0xe {
			SHL(s, opCodeSlice)
		}

	case 0x9000:
		SNEVxVy(s, opCodeSlice)

	case 0xa000:
		SETIR(s, opCodeSlice)

	case 0xb000:
		JUMPVx(s, opCodeSlice)

	case 0xc000:
		RND(s, opCodeSlice)

	case 0xd000:
		DISPLAY(s, opCodeSlice)

	case 0xe000:
		if nn == 0x009e {
			SKP(s, opCodeSlice, keyboardState)

		} else if nn == 0x00a1 {
			SKNP(s, opCodeSlice, keyboardState)
		}

	case 0xf000:
		if nn == 0x07 {
			LDVxDT(s, opCodeSlice)

		} else if nn == 0x0a {
			LDVxK(s, opCodeSlice, keyboardState)

		} else if nn == 0x15 {
			LDDTVx(s, opCodeSlice)

		} else if nn == 0x18 {
			LDSTVx(s, opCodeSlice)

		} else if nn == 0x1e {
			ADDI(s, opCodeSlice)

		} else if nn == 0x29 {
			LDFVx(s, opCodeSlice)

		} else if nn == 0x33 {
			LDBVx(s, opCodeSlice)

		} else if nn == 0x55 {
			LDIVx(s, opCodeSlice)

		} else if nn == 0x65 {
			LDVxI(s, opCodeSlice)
		}
	}
}
