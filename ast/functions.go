package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

type Functions struct {
	*token.Token
	functions []Function
}

func NewFunctions(functions []Function, token *token.Token) *Functions {
	return &Functions{token, functions}
}

func (functions *Functions) BuildSymbolTable(tables *st.SymbolTables) {
	if functions == nil || len(functions.functions) == 0 {
		return
	}
	for _, f := range functions.functions {
		f.BuildSymbolTable(tables)
	}
}

func (functions *Functions) String() string {
	var out bytes.Buffer
	for _, f := range functions.functions {
		out.WriteString(f.String())
	}
	return out.String()
}

func (functions *Functions) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	flag := false
	errors = append(errors, tables.Funcs.Errors...)
	if functions != nil {
		for _, fun := range functions.functions {
			errors = fun.TypeCheck(errors, tables, fun.variable)
			if fun.variable == "main" {
				flag = true
			}
		}
	}
	if !flag {
		errors = append(errors, "missing main function")
	}

	return errors
}

// func (functions *Functions) TranslateToLLVM(st *st.SymbolTables) []*ir.BasicBlock {
// 	// var basicBlock ir.BasicBlock
// 	var bb []*ir.BasicBlock
// 	if functions != nil {
// 		label := "function"
// 		for _, fun := range functions.functions {
// 			var instructions []ir.Instruction
// 			instructions, childs := fun.TranslateToLLVM(st)
// 			functionBlock := &ir.BasicBlock{label, instructions, childs}
// 			bb = append(bb, functionBlock)
// 		}

// 	}

// 	return bb
// }
