package main

import (
	"testing"
)

// func TestStrTok(t *testing.T) {
// 	output := strtok(" set R2, 1234, hello  // Comment", " ,:\t\n")
// 	expected := []string{"set", "R2", "1234", "hello", "//", "Comment"}
// 	if len(output) != len(expected) {
// 		t.Errorf("Got %d, expecting %d", len(output), len(expected))
// 	}
// 	for i := range output {
// 		if output[i] != expected[i] {
// 			t.Errorf("Got %s, expecting %s", output[i], expected[i])
// 		}
// 	}

// }

func TestTokenize(t *testing.T) {
	data := "  set R2, 1234, hello  // Comment"
	token := tokenize(&data)
	if token.Type != TOKEN_TYPE_INST {
		t.Errorf("expected token type to be: %d, got %d", TOKEN_TYPE_INST, token.Type)
	}
	if token.Value != INST_SET {
		t.Errorf("expected token value to be: %d, got %d", INST_SET, token.Value)
	}
	token = tokenize(&data)
	if token.Type != TOKEN_TYPE_REGISTER {
		t.Errorf("expected token type to be: %d, got %d", TOKEN_TYPE_REGISTER, token.Type)
	}
	if token.Value != 2 {
		t.Errorf("expected token value to be: %d, got %d", 2, token.Value)
	}
	token = tokenize(&data)
	if token.Type != TOKEN_TYPE_VALUE {
		t.Errorf("expected token type to be: %d, got %d", TOKEN_TYPE_VALUE, token.Type)
	}
	if token.Value != 1234 {
		t.Errorf("expected token value to be: %d, got %d", 1234, token.Value)
	}
	token = tokenize(&data)
	if token.Type != TOKEN_TYPE_VALUE {
		t.Errorf("expected token type to be: %d, got %d", TOKEN_TYPE_VALUE, token.Type)
	}
	if token.Value != 0 {
		t.Errorf("expected token value to be: %d, got %d", 0, token.Value)
	}
	if len(labels) != 1 {
		t.Errorf("expected label count to be: %d, got %d", 1, len(labels))
	}
	if len(labels) > 0 && labels[0].Name != "hello" {
		t.Errorf("expected label name type to be: %s, got %s", "hello", labels[0].Name)
	}

	token = tokenize(&data)
	if token.Type != TOKEN_TYPE_NULL {
		t.Errorf("expected token type to be: %d, got %d", TOKEN_TYPE_NULL, token.Type)
	}
}
