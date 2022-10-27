package main

import (
	"fmt"
	"math"
)

var (
	initDone bool
	mem      []int32
)

func initCPU() {
	if initDone {
		return
	}
	mem = make([]int32, MAX_MEM)
	for adr, _ := range mem {
		mem[adr] = -1
	}
	initDone = true
}

func InstructionEncode(inst *Inst) int32 {
	return int32(inst.OpCode) |
		int32(math.Pow(2, 10))*int32(inst.Register1) |
		int32(math.Pow(2, 13))*int32(inst.Register2) |
		int32(math.Pow(2, 16))*int32(inst.Arg)
}

func InstructionDecode(int32) *Inst {

	return &Inst{
		OpCode:    1,
		Register1: 1,
		Register2: 1,
		Arg:       1,
	}

}

func writeMem(physicalAddr int, value int32) {
	initCPU()
	// if (! IS_PHYSICAL_ADR(physical_address)) {
	//     fprintf(stderr, "ERROR: write_mem: bad address %d\n", physical_address);
	//     exit(EXIT_FAILURE);
	// }
	if !isPhysicalAddr(int32(physicalAddr)) {
		panic(fmt.Sprintf("writeMem: bad address %d", physicalAddr))
	}
	mem[physicalAddr] = value
}

func readMem(physicalAddr int32) int32 {
	initCPU()
	if !isPhysicalAddr(int32(physicalAddr)) {
		panic(fmt.Sprintf("readMem: bad address %d", physicalAddr))
	}

	return mem[physicalAddr]
}
