package main

import "testing"

func TestInstructionEncode(t *testing.T) {
	output := InstructionEncode(&Inst{
		OpCode:    42,
		Register1: 3,
		Register2: 4,
		Arg:       1984,
	})
	var expected int32 = 130059306
	if output != expected {
		t.Errorf("Expected: %d, got: %d", expected, output)
	}
}

func TestInstructionDecode(t *testing.T) {
	output := InstructionDecode(130059306)
	expected := &Inst{
		OpCode:    42,
		Register1: 3,
		Register2: 4,
		Arg:       1984,
	}
	if output.OpCode != expected.OpCode {
		t.Errorf("Expected: %+v, got: %+v", expected, output)
	}
}
