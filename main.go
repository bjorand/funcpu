package main

import (
	"fmt"
	"os"
)

// setup
const (
	MAX_MEM = 128
)

var (
	word int32
)

// interruption codes
const (
	INT_NONE  = iota // no interrupt
	INT_SEGV         // memory violation
	INT_INST         // unknow instruction
	INT_TRACE        // trace between instructions
	INT_SYSC         // system call
)

// CPU instruction codes
const (
	INST_ADD = iota + 1
	INST_HALT
	INST_IFGT
	INST_IFGE
	INST_IFLT
	INST_IFLE
	INST_JUMP
	INST_LOAD
	INST_NOP
	INST_SET
	INST_STORE
	INST_SUB
	INST_SYSC
	INST_DATA   = -1
	INST_DEFINE = -2
)

// Instruction encoding in 32 bits
type Inst struct {
	OpCode    uint16 // operation code (10 bits)
	Register1 uint8  // 1st register number (3 bits)
	Register2 uint8  // 2nd register number (3 bits)
	Arg       uint16 // argument (16 bits)
}

// PSW defines the processor status word
type PSW struct {
	PC int32   // program counter
	SB int32   // segment begin
	SE int32   // segment end
	IN int32   // interrupt number
	DR []int32 // data registers
	IR Inst    // instruction registers
}

func isPhysicalAddr(a int32) bool {
	return 0 <= a && a < MAX_MEM
}

func isLogicalAddr(a int32, psw *PSW) bool {
	return psw.SB <= a && a <= psw.SE
}

func cpuInit() {
	// ensure structure size are equal
	// word == Inst.Size
}

func systemInit() {
	fmt.Println("Booting...")
	psw := PSW{
		PC: 20,
		SB: 20,
		SE: 30,
	}
	assemble(psw.SB, os.Args[1])
}

func main() {
	usedLabels = make([]*UsedLabel, 0)
	labels = make([]*Label, 0)
	systemInit()
	fmt.Println(mem)
	// for {
	//
	// }
}
