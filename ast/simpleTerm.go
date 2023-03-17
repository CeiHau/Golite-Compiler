package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type SimpleTerm struct {
	*token.Token
	operators []string
	terms     []Term
}

func NewSimpleTerm(terms []Term, operators []string, token *token.Token) *SimpleTerm {
	return &SimpleTerm{token, operators, terms}
}

func (simpleTerm *SimpleTerm) String() string {
	var out bytes.Buffer
	for index, term := range simpleTerm.terms {
		out.WriteString(term.String())
		if index != len(simpleTerm.terms)-1 {
			out.WriteString(simpleTerm.operators[index])
		}
	}
	return out.String()
}

func (simpleTerm *SimpleTerm) GetType(symTable *st.SymbolTables, scope string) string {
	if len(simpleTerm.terms) == 1 {
		return simpleTerm.terms[0].GetType(symTable, scope)
	} else {
		return "int"
	}
}

func (simpleTerm *SimpleTerm) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	for _, term := range simpleTerm.terms {
		errors = term.TypeCheck(errors, tables, scope)
	}

	if len(simpleTerm.terms) <= 1 {
		return errors
	} else {
		for _, t := range simpleTerm.terms {
			if t.GetType(tables, scope) != "int" {
				errors = append(errors, "bool cannot be operated by + or -")
				return errors
			}
		}
	}
	return errors
}

func (simpleTerm *SimpleTerm) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	result := simpleTerm.terms[0].TranslateToLLVM(tables, block, scope)
	// Only one Term don't need to new register
	if len(simpleTerm.terms) != 1 {

		for index, op := range simpleTerm.operators {
			op1 := simpleTerm.terms[index+1].TranslateToLLVM(tables, block, scope)
			temp := ir.NewRegister(tables, block.FuncName)
			switch op {
			case "+":
				block.Instructions = append(block.Instructions, ir.NewAdd(temp.GetRegisterNum(), result, op1))
			case "-":
				block.Instructions = append(block.Instructions, ir.NewSub(temp.GetRegisterNum(), result, op1))
			}
			result = temp
		}
	}
	return result
}
