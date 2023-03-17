package ast

import (
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type NilLiteral struct {
	*token.Token
	ty tp.Type
}

func NewNilLit(token *token.Token) *NilLiteral {
	return &NilLiteral{token, nil}
}

func (n *NilLiteral) String() string {
	return fmt.Sprintf("nil")
}

func (n *NilLiteral) GetType(symTable *st.SymbolTables, scope string) string {
	return "nil"
}
func (n *NilLiteral) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	return errors
}

func (n *NilLiteral) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	res := ir.NewRegister(tables, block.FuncName)
	block.Instructions = append(block.Instructions, ir.NewLoad(res.GetRegisterNum(), "%struct."+scope+"*", "@.nil"+scope))
	return res
}
