package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	nbLabels int
	// InstructionDictory *[]Instruction
	InstructionDictory = map[string]int{
		"set":    INST_SET,
		"add":    INST_ADD,
		"sub":    INST_SUB,
		"nop":    INST_NOP,
		"jump":   INST_JUMP,
		"ifgt":   INST_IFGT,
		"ifge":   INST_IFGE,
		"iflt":   INST_IFLT,
		"ifle":   INST_IFLE,
		"sysc":   INST_SYSC,
		"store":  INST_STORE,
		"load":   INST_LOAD,
		"halt":   INST_HALT,
		"data":   INST_DATA,
		"define": INST_DEFINE,
	}
	RegisterDictory = map[string]int{
		"R0": 0,
		"R1": 1,
		"R2": 2,
		"R3": 3,
		"R4": 4,
		"R5": 5,
		"R6": 6,
		"R7": 7,
		"r0": 0,
		"r1": 1,
		"r2": 2,
		"r3": 3,
		"r4": 4,
		"r5": 5,
		"r6": 6,
		"r7": 7,
	}
	labels      []*Label
	usedLabels  []*UsedLabel
	logical_adr = 0
)

type Label struct {
	Name string
	Addr int
}

type UsedLabel struct {
	Label int
	Addr  int
}

type Instruction struct {
	Name string
	Code int
}

const (
	TOKEN_TYPE_NULL = iota
	TOKEN_TYPE_REGISTER
	TOKEN_TYPE_INST
	TOKEN_TYPE_VALUE

	MAX_LABELS       = 30
	MAX_LABEL_LENGTH = 60
)

type Token struct {
	Type    int
	Value   int
	IsIdent bool
}

func MakeInst(addr int, inst *Inst) {
	w := InstructionEncode(inst)
	writeMem(addr, w)
}

func strtok(data *string, separators string) string {
	re := regexp.MustCompile(fmt.Sprintf("[%s]+", separators))
	split := re.Split(*data, -1)
	set := []string{}
	for i := range split {
		if split[i] != "" {
			set = append(set, split[i])
		}
	}
	if len(set) == 0 {
		return ""
	}
	*data = strings.Join(set[1:], " ")
	return set[0]
}

func tokenize(data *string) *Token {
	s := strtok(data, " ,:\t\n")
	fmt.Println("-->", s)
	null := &Token{
		Type: TOKEN_TYPE_NULL,
	}
	if s == "" {
		return null
	}

	if strings.HasPrefix(s, "//") {
		return null
	}

	v, err := strconv.Atoi(s)
	if err == nil {
		return &Token{
			Type:  TOKEN_TYPE_VALUE,
			Value: v,
		}
	}
	inst := InstructionDictory[s]
	if inst != 0 {
		return &Token{
			Type:  TOKEN_TYPE_INST,
			Value: inst,
		}
	}
	reg := RegisterDictory[s]
	if reg != 0 {
		return &Token{
			Type:  TOKEN_TYPE_REGISTER,
			Value: reg,
		}
	}

	labelIndex := 0
	for i := 0; i < nbLabels; i++ {

		if labels[i].Name == s {
			break
		}
		labelIndex = i + 1
	}

	if labelIndex == nbLabels {
		if !(nbLabels < MAX_LABELS) {
			panic("Too many labels")
		}
		if !(len(s) < MAX_LABEL_LENGTH) {
			panic("Label name is too long")
		}
		labels = append(labels, &Label{
			Name: s,
			Addr: -1,
		})
		// labels[nbLabels] = &Label{
		// 	Name: s,
		// 	Addr: -1,
		// }
		nbLabels++
	}
	if !(len(usedLabels) < MAX_LABELS) {
		panic("Too many used labels")
	}

	usedLabels = append(usedLabels, &UsedLabel{
		Label: labelIndex,
		Addr:  logical_adr,
	})
	return &Token{
		Type:    TOKEN_TYPE_VALUE,
		Value:   labelIndex,
		IsIdent: true,
	}
}

func assembleDefine(line *string) {
	token := tokenize(line)
	if !(token.Type == TOKEN_TYPE_VALUE && token.IsIdent) {
		panic("instruction define incorrecte")
	}
	label := token.Value
	token = tokenize(line)
	if !(token.Type == TOKEN_TYPE_VALUE && !token.IsIdent) {
		panic("instruction define incorrecte")
	}
	labels[label].Addr = token.Value
}

func assembleLine(line string) {
	inst := &Inst{}
	token := tokenize(&line)
	if token.Type == TOKEN_TYPE_NULL {
		return
	}
	if token.Type == TOKEN_TYPE_VALUE && token.IsIdent {
		fmt.Printf("%+v\n", labels[0])
		if labels[token.Value].Addr >= 0 {
			panic(fmt.Sprintf("Already used label: %s", labels[token.Value].Name))
		}
		labels[token.Value].Addr = logical_adr
		token = tokenize(&line)

	}
	if token.Type == TOKEN_TYPE_NULL {
		return
	}
	if token.Type != TOKEN_TYPE_INST {
		panic(fmt.Sprintf("Incorrect instruction %+v", token))
	}
	if token.Type == TOKEN_TYPE_INST && token.Value == INST_DEFINE {
		fmt.Println("DEFINEEE")
		assembleDefine(&line)
		return

	}
	inst.OpCode = uint16(token.Value)
	token = tokenize(&line)
	if token.Type == TOKEN_TYPE_NULL {
		MakeInst(logical_adr, inst)
		logical_adr++
		return
	}
	if token.Type == TOKEN_TYPE_REGISTER {
		inst.Register1 = uint8(token.Value)
		token = tokenize(&line)
	}

	if token.Type == TOKEN_TYPE_REGISTER {
		inst.Register2 = uint8(token.Value)
		token = tokenize(&line)
	}
	if token.Type == TOKEN_TYPE_VALUE {
		inst.Arg = uint16(token.Value)
		token = tokenize(&line)
	}
	if token.Type != TOKEN_TYPE_NULL {
		panic("instruction incorrecte")
	}
	if inst.OpCode == 1023 {
		writeMem(logical_adr, 0)
	} else {
		MakeInst(logical_adr, inst)
	}
	logical_adr++
}

func assemble(addr int32, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		assembleLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
