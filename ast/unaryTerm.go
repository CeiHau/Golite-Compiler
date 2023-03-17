package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type UnaryTerm struct {
	*token.Token
	operator     string
	selectorTerm SelectorTerm
}

func NewUnaryTerm(selectorTerm SelectorTerm, operator string, token *token.Token) *UnaryTerm {
	return &UnaryTerm{token, operator, selectorTerm}

}

func (unaryTerm *UnaryTerm) String() string {
	var out bytes.Buffer
	if len(unaryTerm.operator) > 0 {
		out.WriteString(unaryTerm.operator)
		out.WriteString(" ")
	}
	out.WriteString(unaryTerm.selectorTerm.String())
	return out.String()
}

func (unaryTerm *UnaryTerm) GetType(symTable *st.SymbolTables, scope string) string {

	if unaryTerm.operator != "!" {
		return unaryTerm.selectorTerm.GetType(symTable, scope)
	} else {
		return "bool"
	}
}

func (unaryTerm *UnaryTerm) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	errors = unaryTerm.selectorTerm.TypeCheck(errors, tables, scope)
	if errors != nil {
		return errors
	}
	if unaryTerm.GetType(tables, scope) == "int" && unaryTerm.operator == "!" {
		errors = append(errors, "int can not have NOT operation")
	}
	if unaryTerm.GetType(tables, scope) == "bool" && unaryTerm.operator == "-" {
		errors = append(errors, "bool can not have MINUS operation")
	}
	return errors
}

func (unaryTerm *UnaryTerm) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	// Only one UnaryTerm don't need to new register
	if unaryTerm.operator == "!" || unaryTerm.operator == "-" {
		// Use fneg instruction
		result := ir.NewRegister(tables, block.FuncName)
		switch unaryTerm.operator {
		case "!":
			block.Instructions = append(block.Instructions, ir.NewXor(result.GetRegisterNum(), unaryTerm.selectorTerm.TranslateToLLVM(tables, block, scope), NewBoolLit(true, nil).TranslateToLLVM(tables, block, scope)))
		case "-":
			block.Instructions = append(block.Instructions, ir.NewNe(result.GetRegisterNum(), unaryTerm.selectorTerm.TranslateToLLVM(tables, block, scope)))

		}

		return result
	} else {

		return unaryTerm.selectorTerm.TranslateToLLVM(tables, block, scope)
	}
}
