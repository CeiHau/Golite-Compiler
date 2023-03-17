package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type RelationTerm struct {
	*token.Token
	operators   []string
	simpleTerms []SimpleTerm
}

func NewRelationTerm(operators []string, simpleTerms []SimpleTerm, token *token.Token) *RelationTerm {
	return &RelationTerm{token, operators, simpleTerms}
}

func (relationTerm *RelationTerm) String() string {
	var out bytes.Buffer
	for index, simpleTerm := range relationTerm.simpleTerms {
		out.WriteString(simpleTerm.String())
		if index != len(relationTerm.simpleTerms)-1 {
			out.WriteString(relationTerm.operators[index])
		}
	}

	return out.String()
}

func (relationTerm *RelationTerm) GetType(symTable *st.SymbolTables, scope string) string {
	if len(relationTerm.simpleTerms) == 1 {
		return relationTerm.simpleTerms[0].GetType(symTable, scope)
	} else {
		return "bool"
	}
}

func (relationTerm *RelationTerm) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	for _, simplet := range relationTerm.simpleTerms {
		errors = simplet.TypeCheck(errors, tables, scope)
	}
	if len(relationTerm.simpleTerms) <= 1 {
		return errors
	} else {
		for _, t := range relationTerm.simpleTerms {
			if t.GetType(tables, scope) != "int" {
				errors = append(errors, "bool cannot be operated by MORETHAN|LESSTHAN|LEQ|GEQ")
				return errors
			}
		}
	}
	return errors
}

func (relationTerm *RelationTerm) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	result := relationTerm.simpleTerms[0].TranslateToLLVM(tables, block, scope)
	// Only one SimpleTerm don't need to new register
	if len(relationTerm.simpleTerms) != 1 {
		for index, op := range relationTerm.operators {
			op1 := relationTerm.simpleTerms[index+1].TranslateToLLVM(tables, block, scope)
			temp := ir.NewRegister(tables, block.FuncName)
			switch op {
			case ">":
				block.Instructions = append(block.Instructions, ir.NewCmp(result.GetRegisterNum(), "sgt", result, op1))
			case "<":
				block.Instructions = append(block.Instructions, ir.NewCmp(result.GetRegisterNum(), "slt", result, op1))
			case "<=":
				block.Instructions = append(block.Instructions, ir.NewCmp(result.GetRegisterNum(), "sle", result, op1))
			case ">=":
				block.Instructions = append(block.Instructions, ir.NewCmp(result.GetRegisterNum(), "sge", result, op1))
			}
			result = temp
		}
	}
	// Return the new Register
	return result
}
