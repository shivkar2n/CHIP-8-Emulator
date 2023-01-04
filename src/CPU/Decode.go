package CPU

//package main

import (
	"encoding/binary"
	"fmt"
	_ "fmt"
	//"os"
	//"github.com/shivkar2n/Chip8-Emulator/helpers"
)

func (s *State) InstructionDecode(opCodeSlice [2]byte) {
	opcodeBytes := binary.BigEndian.Uint16(opCodeSlice[:])

	n := int(opcodeBytes) & 0x000f
	nn := int(opcodeBytes) & 0x00ff
	nnn := int(opcodeBytes) & 0x0fff
	x := int(opcodeBytes) & 0x0f00 >> 8
	y := int(opcodeBytes) & 0x00f0 >> 4
	opCode0 := int(opcodeBytes) & 0xf000
	opCode := int(opcodeBytes)

	switch opCode0 {
	case 0x0000:
		if opCode == 0x00e0 {
			fmt.Printf("CLS\n")
			CLS(s)
		} else if opCode == 0x00ee {
			fmt.Printf("RET\n")
			//RET(s)
		} else {
			fmt.Printf("SYS %x\n", nnn)
		}

	case 0x1000:
		fmt.Printf("JP %x\n", nnn)
		JUMP(s, opCodeSlice)

	case 0x2000:
		fmt.Printf("CALL %x\n", nnn)
		//CALL(s, opCodeSlice)

	case 0x3000:
		fmt.Printf("SE V%x, %x\n", x, nn)

	case 0x4000:
		fmt.Printf("SNE V%x, %x\n", x, nn)

	case 0x5000:
		fmt.Printf("SE V%x, V%x\n", x, y)

	case 0x6000:
		fmt.Printf("LD V%x, %x\n", x, nn)
		LD(s, opCodeSlice)

	case 0x7000:
		fmt.Printf("ADD V%x, %x\n", x, nn)
		ADD(s, opCodeSlice)

	case 0x8000:
		operandCode := n
		if operandCode == 0x0000 {
			fmt.Printf("LD V%x, V%x\n", x, y)

		} else if operandCode == 0x1 {
			fmt.Printf("OR V%x, V%x\n", x, y)

		} else if operandCode == 0x2 {
			fmt.Printf("AND V%x, V%x\n", x, y)

		} else if operandCode == 0x3 {
			fmt.Printf("XOR V%x, V%x\n", x, y)

		} else if operandCode == 0x4 {
			fmt.Printf("ADD V%x, V%x\n", x, y)

		} else if operandCode == 0x5 {
			fmt.Printf("SUB V%x, V%x\n", x, y)

		} else if operandCode == 0x6 {
			fmt.Printf("SHR V%x {, V%x}\n", x, y)

		} else if operandCode == 0x7 {
			fmt.Printf("SUBN V%x, V%x\n", x, y)

		} else if operandCode == 0xe {
			fmt.Printf("SHL V%x {, V%x}\n", x, y)
		}

	case 0x9000:
		fmt.Printf("SNE V%x, V%x\n", x, y)

	case 0xa000:
		fmt.Printf("LD I, %x\n", nnn)
		SETIR(s, opCodeSlice)

	case 0xb000:
		fmt.Printf("JP V0, %x\n", nnn)

	case 0xc000:
		fmt.Printf("RND V%x, %x\n", x, nn)

	case 0xd000:
		fmt.Printf("DRW V%x, V%x, %x\n", x, y, n)
		DISPLAY(s, opCodeSlice)

	case 0xe000:
		if nn == 0x009e {
			fmt.Printf("SKP V%x\n", x)

		} else if nn == 0x00a1 {
			fmt.Printf("SKNP V%s\n", x)
		}

	case 0xf000:
		if nn == 0x07 {
			fmt.Printf("LD V%x, DT\n", x)

		} else if nn == 0x0a {
			fmt.Printf("LD V%x, K\n", x)

		} else if nn == 0x15 {
			fmt.Printf("LD DT, V%x\n", x)

		} else if nn == 0x18 {
			fmt.Printf("LD ST, V%x\n", x)

		} else if nn == 0x1e {
			fmt.Printf("ADD I, V%x\n", x)

		} else if nn == 0x29 {
			fmt.Printf("LD F, V%x\n", x)

		} else if nn == 0x33 {
			fmt.Printf("LD B, V%x\n", x)

		} else if nn == 0x55 {
			fmt.Printf("LD [I], V%x\n", x)

		} else if nn == 0x65 {
			fmt.Printf("LD V%x, [I]\n", x)
		}
	}
}
