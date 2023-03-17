package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type EqualTerm struct {
	*token.Token
	operators     []string
	relationTerms []RelationTerm
}

func NewEqualTerm(operators []string, relationTerms []RelationTerm, token *token.Token) *EqualTerm {
	return &EqualTerm{token, operators, relationTerms}

}
func (equalTerm *EqualTerm) String() string {
	var out bytes.Buffer
	for index, relationTerm := range equalTerm.relationTerms {
		out.WriteString(relationTerm.String())
		if index != len(equalTerm.relationTerms)-1 {
			out.WriteString(equalTerm.operators[index])
		}
	}
	return out.String()
}

func (equalTerm *EqualTerm) GetType(symTable *st.SymbolTables, scope string) string {
	if len(equalTerm.relationTerms) <= 1 {
		return equalTerm.relationTerms[0].GetType(symTable, scope)
	} else {
		return "bool"
	}
}

func (equalTerm *EqualTerm) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	for _, relationt := range equalTerm.relationTerms {
		errors = relationt.TypeCheck(errors, tables, scope)
	}
	if len(equalTerm.relationTerms) <= 1 {
		return errors
	} else {

		for index, t := range equalTerm.relationTerms {
			if index == 0 {
				continue
			}
			relationtermType := equalTerm.relationTerms[index-1].GetType(tables, scope)
			if t.GetType(tables, scope) != relationtermType && t.String() != "nil" {
				print(t.String(), relationtermType, equalTerm.relationTerms[0].String(), "\n")
				errors = append(errors, "different types cannot be operated by EQUALS|NOTEQUALS")
				return errors
			}
		}
	}
	return errors
}

func (equalTerm *EqualTerm) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	result := equalTerm.relationTerms[0].TranslateToLLVM(tables, block, scope)
	// Only one RelationTerms don't need to new register
	if len(equalTerm.relationTerms) != 1 {
		for index, op := range equalTerm.operators {
			op1 := equalTerm.relationTerms[index+1].TranslateToLLVM(tables, block, scope)
			temp := ir.NewRegister(tables, block.FuncName)
			switch op {
			case "==":
				block.Instructions = append(block.Instructions, ir.NewCmp(result.GetRegisterNum(), "eq", result, op1))
			case "!=":
				block.Instructions = append(block.Instructions, ir.NewCmp(result.GetRegisterNum(), "ne", result, op1))
			}
			result = temp
		}
	}
	return result
}
