package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

type Malloc struct {
	mem int
}

func NewMalloc(mem int) *Malloc {
	return &Malloc{mem}
}

// i8* @malloc(i32 16)
func (instr *Malloc) String() string {

	var out bytes.Buffer
	out.WriteString(" i8* @malloc(i32")
	if instr.mem != 0 {
		memory := fmt.Sprintf(" %d", instr.mem)
		out.WriteString(memory)
	}
	out.WriteString(")")

	return out.String()

}

func (instr *Malloc) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {

	return nil
}
