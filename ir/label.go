package ir

import (
	"golite/codegen"
	st "golite/symboltable"
	"strconv"
)

var labelNum int

type Label struct {
}

func InitLable() {
	labelNum = 0
}

func NewLabel() string {
	rslt := "L" + strconv.Itoa(labelNum)
	labelNum++
	return rslt
}

func LabelInstruction() *Label {
	return &Label{}
}

func (lable *Label) String() string {
	return "L" + strconv.Itoa(labelNum)
}

// func (labe)

func (instr *Label) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
