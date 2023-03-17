package ir

import (
	"bytes"
	"golite/codegen"
	st "golite/symboltable"
)

// @VARIABLE_NAME = common global <ty> <constant value>
// @globalInit = common global i64 0
type Declaration struct {
	variableName string
	ty           string
	value        string
}

func NewDeclaration(variableName string, ty string, value string) *Declaration {
	return &Declaration{variableName, ty, value}
}

func (instr *Declaration) String() string {
	var out bytes.Buffer
	out.WriteString("@")
	out.WriteString(instr.variableName)
	out.WriteString(" = common global ")
	out.WriteString(instr.ty)
	out.WriteString(" ")
	out.WriteString(instr.value)
	return out.String()
}

func (instr *Declaration) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
