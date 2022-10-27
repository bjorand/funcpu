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
	OpCode    int16  // operation code (10 bits)
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
	RI Inst    // instruction registers
}

func (p *PSW) dumpCPU() {
	interrupts := []string{
		"NONE", "SEGV", "INST", "TRACE", "SYSC", "KEYB",
	}
	fmt.Printf("PC = %6d | ", p.PC)
	if p.IN < 6 {
		fmt.Printf("IN = %6s\n", interrupts[p.IN])
	} else {
		fmt.Printf("IN = %6d\n", p.IN)
	}
	fmt.Printf("SB = %6d | SE = %6d\n", p.SB, p.SE)
	for i := 0; i < 8; i = i + 2 {
		fmt.Printf("R%d = %6d | R%d = %6d\n", i, p.DR[i], i+1, p.DR[i+1])
	}
	var name string
	instructions := []string{
		"ADD", "HALT", "IFGT", "IFGE", "IFLT", "IFLE",
		"JUMP", "LOAD",
		"NOP", "SET", "STORE", "SUB", "SYSC",
	}
	if p.RI.OpCode < INST_SYSC {
		name = instructions[p.RI.OpCode]
	} else {
		name = "?"
	}
	fmt.Printf("RI  = %s R%d, R%d, %d \n", name, p.RI.Register1, p.RI.Register2, p.RI.Arg)
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

func systemInit() *PSW {
	fmt.Println("Booting...")
	psw := &PSW{
		PC: 20,
		SB: 20,
		SE: 30,
		DR: make([]int32, 8),
	}
	assemble(psw.SB, os.Args[1])
	return psw
}

func ReadLogicalMem(logicalAddr int32, psw *PSW) int32 {
	if !isLogicalAddr(logicalAddr, psw) {
		psw.IN = INT_SEGV
		return 0
	}
	return readMem(logicalAddr)

}

func (psw *PSW) CPU() {
	initCPU()

	psw.IN = INT_NONE

	//  m = keyboard_event(m);
	// if (m.IN) return (m);

	value := ReadLogicalMem(psw.PC, psw)
	if psw.IN == INT_NONE {
		return
	}
	fmt.Println(value)
	// psw.RI = InstructionDecode(value)
}

func main() {
	usedLabels = make([]*UsedLabel, 0)
	labels = make([]*Label, 0)
	psw := systemInit()
	psw.CPU()
	psw.dumpCPU()
}
