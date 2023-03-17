package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Term struct {
	*token.Token
	operators  []string
	unaryTerms []UnaryTerm
}

func NewTerm(unaryTerms []UnaryTerm, operators []string, token *token.Token) *Term {
	return &Term{token, operators, unaryTerms}
}

func (term *Term) String() string {
	var out bytes.Buffer
	for index, ut := range term.unaryTerms {
		out.WriteString(ut.String())
		if index != len(term.unaryTerms)-1 {
			out.WriteString(term.operators[index])
		}
	}
	return out.String()
}
func (term *Term) GetType(symTable *st.SymbolTables, scope string) string {
	if len(term.unaryTerms) == 1 {
		return term.unaryTerms[0].GetType(symTable, scope)
	} else {
		return "int"
	}
}
func (term *Term) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	for _, unaryt := range term.unaryTerms {
		errors = unaryt.TypeCheck(errors, tables, scope)
	}
	if len(term.unaryTerms) <= 1 {
		return errors
	} else {
		for _, unaryt := range term.unaryTerms {
			if unaryt.GetType(tables, scope) != "int" {
				errors = append(errors, "bool cannot be operated by * or /")
				return errors
			}
		}
	}
	return errors

}

func (term *Term) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {

	result := term.unaryTerms[0].TranslateToLLVM(tables, block, scope)
	// Only one UnaryTerm don't need to new register
	if len(term.unaryTerms) != 1 {

		for index, op := range term.operators {

			op1 := term.unaryTerms[index+1].TranslateToLLVM(tables, block, scope)
			temp := ir.NewRegister(tables, block.FuncName)
			switch op {
			case "*":
				block.Instructions = append(block.Instructions, ir.NewMul(temp.GetRegisterNum(), result, op1))
			case "/":
				block.Instructions = append(block.Instructions, ir.NewSdiv(temp.GetRegisterNum(), result, op1))
			}
			result = temp
		}
	}
	return result
}
