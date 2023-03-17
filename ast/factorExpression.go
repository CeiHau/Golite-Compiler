package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type FactorExpression struct {
	*token.Token
	expression Expression
	Usable     bool
}

func NewFactorExpression(expression Expression, token *token.Token) *FactorExpression {
	return &FactorExpression{token, expression, true}
}

func (fe *FactorExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(fe.expression.String())
	out.WriteString(")")
	return out.String()
}

func (fe *FactorExpression) GetType(symTable *st.SymbolTables, scope string) string {
	return fe.expression.GetType(symTable, scope)
}

func (fe *FactorExpression) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	return fe.expression.TypeCheck(errors, tables, scope)
}

func (fe *FactorExpression) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	return fe.expression.TranslateToLLVM(tables, block, scope)
}
