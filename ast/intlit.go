package ast

import (
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type IntLiteral struct {
	*token.Token       // The token for the integer literal
	Value        int64 // The value for the integer literal
	ty           tp.Type
}

func NewIntLit(value int64, token *token.Token) *IntLiteral {
	return &IntLiteral{token, value, tp.StringToType("int")}
}
func (il *IntLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}
func (il *IntLiteral) GetType(symTable *st.SymbolTables, scope string) string {
	//Exercise: You should think about what should be returned here.
	return "int"
}

func (il *IntLiteral) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	return errors
}

func (il *IntLiteral) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	return ir.NewOperand(int(il.Value), false)
}
