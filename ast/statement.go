package ast

import (
	"golite/ir"
	st "golite/symboltable"
)

type Statement interface {
	String() string
	// BuildSymbolTable(tables *st.SymbolTables)
	TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string
	CheckReturn(symTable *st.SymbolTables) bool
	TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock
}
