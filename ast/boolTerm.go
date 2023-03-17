package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type BoolTerm struct {
	*token.Token
	equalTerms []EqualTerm
}

func NewBoolTerm(equalTerms []EqualTerm, token *token.Token) *BoolTerm {
	return &BoolTerm{token, equalTerms}
}

func (boolTerm *BoolTerm) String() string {
	var out bytes.Buffer
	for index, equalTerm := range boolTerm.equalTerms {
		out.WriteString(equalTerm.String())
		if index != len(boolTerm.equalTerms)-1 {
			out.WriteString("&&")
		}
	}
	return out.String()
}
func (boolTerm *BoolTerm) GetType(symTable *st.SymbolTables, scope string) string {
	if len(boolTerm.equalTerms) <= 1 {
		return boolTerm.equalTerms[0].GetType(symTable, scope)
	} else {
		return "bool"
	}
}

func (boolTerm *BoolTerm) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	for _, equal := range boolTerm.equalTerms {
		errors = equal.TypeCheck(errors, tables, scope)
	}
	if len(boolTerm.equalTerms) <= 1 {
		return errors
	} else {
		for _, equalt := range boolTerm.equalTerms {
			if equalt.GetType(tables, scope) != "bool" {
				errors = append(errors, "different types cannot be operated by AND")
				return errors
			}
		}
	}
	return errors
}

func (boolTerm *BoolTerm) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	result := boolTerm.equalTerms[0].TranslateToLLVM(tables, block, scope)
	if len(boolTerm.equalTerms) != 1 {
		// Generate the operation
		for index, _ := range boolTerm.equalTerms {
			temp := ir.NewRegister(tables, block.FuncName)
			op1 := boolTerm.equalTerms[index+1].TranslateToLLVM(tables, block, scope)
			result = ir.NewRegister(tables, block.FuncName)
			block.Instructions = append(block.Instructions, ir.NewAnd(result.GetRegisterNum(), result, op1))
			result = temp
		}
	}
	return result
}
