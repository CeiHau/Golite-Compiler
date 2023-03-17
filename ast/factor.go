package ast

import (
	"golite/ir"
	st "golite/symboltable"
)

type Factor interface {
	String() string
	GetType(symTable *st.SymbolTables, scope string) string
	TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string
	TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand
}
