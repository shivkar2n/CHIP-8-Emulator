package CPU

import (
	"fmt"
)

const INSTRUCTION_SIZE int = 2

type State struct {
	Memory     [4096]byte
	V          [16]byte
	PC         [2]byte
	IR         [2]byte
	DelayTimer byte
	SoundTimer byte
	SP         byte
	Stack      [16][2]byte
}

func (s *State) WriteMemory(buffer []byte, offset int) { // Write bytes slice to CPU memory
	for i := 0; i < len(buffer); i++ {
		s.Memory[offset+i] = buffer[i]
	}
}

func (s *State) ReadMemory(buffer []byte, offset int) { // Read bytes slice from CPU memory
	for i := 0; i < len(buffer); i++ {
		buffer[i] = s.Memory[offset+i]
	}
}

func (s *State) IncrementPC() {
	buf := [2]byte{}
	newPC := int(s.PC[0])*256 + int(s.PC[1]) + 2
	buf[0], buf[1] = byte(newPC>>8), byte(newPC&0x00ff)
	s.PC = buf
}

func (s *State) Print() { // Print contents of CPU State
	fmt.Printf("\n------CPU STATE------\n")
	//fmt.Println("Memory:\n")
	//for i, mem := range s.Memory {
	//    fmt.Printf("Mem[%d]: %x\n", i, mem)
	//}
	//fmt.Println()

	fmt.Printf("Index Register: %x\n", s.IR)
	fmt.Printf("Program Counter: %x\n", s.PC)
	for i := 0; i < len(s.V); i++ {
		fmt.Printf("Register V%x: %x\n", i, s.V[i])
	}

	//fmt.Printf("Delay Timer: %x\n", s.DelayTimer)
	//fmt.Printf("Sound Timer: %x\n", s.SoundTimer)
	//fmt.Printf("Stack Pointer: %x\n", s.SP)
	//fmt.Println("Stack:\n")
	//for _, mem := range s.Stack {
	//    fmt.Printf("%x ", mem)
	//}
	fmt.Println()

}
