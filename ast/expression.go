package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Expression struct {
	*token.Token
	boolTerms []BoolTerm
}

func NewExpression(boolTerms []BoolTerm, token *token.Token) *Expression {
	return &Expression{token, boolTerms}
}

func (expression *Expression) String() string {
	var out bytes.Buffer
	for index, boolTerm := range expression.boolTerms {
		out.WriteString(boolTerm.String())
		if index != len(expression.boolTerms)-1 {
			out.WriteString("||")
		}
	}
	return out.String()
}
func (expression *Expression) GetType(symTable *st.SymbolTables, scope string) string {
	if len(expression.boolTerms) <= 1 {
		return expression.boolTerms[0].GetType(symTable, scope)
	} else {
		return "bool"
	}

}

func (expression *Expression) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	for _, boolt := range expression.boolTerms {
		errors = boolt.TypeCheck(errors, tables, scope)
	}
	if len(expression.boolTerms) <= 1 {
		return errors
	} else {
		for _, boolt := range expression.boolTerms {
			if boolt.GetType(tables, scope) != "bool" {
				errors = append(errors, "different types cannot be operated by OR")
				return errors
			}
		}
	}
	return errors
}

func (expression *Expression) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	result := expression.boolTerms[0].TranslateToLLVM(tables, block, scope)
	// Only one BoolTerm don't need to new register
	if len(expression.boolTerms) != 1 {
		for index, _ := range expression.boolTerms {
			op1 := expression.boolTerms[index+1].TranslateToLLVM(tables, block, scope)
			temp := ir.NewRegister(tables, block.FuncName)
			block.Instructions = append(block.Instructions, ir.NewOr(result.GetRegisterNum(), result, op1))
			result = temp
		}
	}
	return result

}
