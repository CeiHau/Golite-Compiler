package ir

import (
	"bytes"
	"golite/codegen"
	st "golite/symboltable"
)

type Alloca struct {
	result string
	ty     string
}

func NewAlloca(result string, ty string) *Alloca {
	return &Alloca{result, ty}
}

func (instr *Alloca) String() string {
	var out bytes.Buffer
	out.WriteString("%")
	out.WriteString(instr.result)
	out.WriteString(" = alloca ")
	out.WriteString(instr.ty)
	return out.String()
}

func (instr *Alloca) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
