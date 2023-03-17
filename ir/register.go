package ir

import (
	"fmt"
	st "golite/symboltable"
)

var registerNum int

func InitRegister() {
	registerNum = 0
}

func NewRegister(tables *st.SymbolTables, funcName string) *Operand {
	funcEntry := tables.Funcs.Contains(funcName)

	op := Operand{registerNum, true}
	if funcEntry == nil {
		fmt.Println(funcName)
		fmt.Println("nil")
		fmt.Println(registerNum)
	}

	funcEntry.Register = append(funcEntry.Register, op.String())

	registerNum++
	return &op
}
