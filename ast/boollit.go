package ast

import (
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type BoolLiteral struct {
	*token.Token      // The token for the integer literal
	Value        bool // The value for the integer literal
	ty           tp.Type
}

func NewBoolLit(value bool, token *token.Token) *BoolLiteral {
	return &BoolLiteral{token, value, tp.StringToType("bool")}
}
func (bl *BoolLiteral) String() string {
	return fmt.Sprintf("%t", bl.Value)
}
func (bl *BoolLiteral) GetType(symTable *st.SymbolTables, scope string) string {
	//Exercise: You should think about what should be returned here.
	return "bool"
}

func (bl *BoolLiteral) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	return errors
}

func (bl *BoolLiteral) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	if bl.Value {
		return ir.NewOperand(1, false)
	}
	return ir.NewOperand(0, false)

}
