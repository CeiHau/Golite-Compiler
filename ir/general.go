package ir

import (
	"golite/codegen"
	st "golite/symboltable"
)

type General struct {
	content string
}

func NewGeneral(content string) *General {
	return &General{content}
}

func (instr *General) String() string {
	return instr.content
}

func (instr *General) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
