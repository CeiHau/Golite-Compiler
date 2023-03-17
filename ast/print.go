// modified
package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
	"strings"
)

type Print struct {
	*token.Token
	str   string
	exprs []Expression
}

func NewPrint(str string, expr []Expression, token *token.Token) *Print {
	return &Print{token, str, expr}
}

func (print *Print) String() string {
	var out bytes.Buffer
	out.WriteString("printf (")
	out.WriteString(print.str)
	for _, expr := range print.exprs {
		out.WriteString(", ")
		out.WriteString(expr.String())
	}
	out.WriteString(");\n")

	return out.String()
}
func (print *Print) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {
	var args []string
	for _, exp := range print.exprs {
		ty := tp.StringToType(exp.GetType(tables, scope)).LLVM()
		originalInstructionLength := len(block.Instructions)
		value := exp.TranslateToLLVM(tables, block, scope)
		if originalInstructionLength+1 < len(block.Instructions) {
			rslt := ir.NewRegister(tables, block.FuncName)
			block.Instructions = append(block.Instructions, ir.NewLoad(rslt.GetRegisterNum(), ty, value.String()))
			args = append(args, ty+" "+rslt.String())
		} else {
			args = append(args, ty+" "+value.String())
		}
	}
	printread := ir.NewPR("printf", len(print.str)-2+1, args)
	block.Instructions = append(block.Instructions, printread)
	ir.GlobalDeclare.Instructions = append(ir.GlobalDeclare.Instructions, ir.NewPRLL(printread, len(print.str)-2+1, print.str))

	return block
}
func (print *Print) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}
func (print *Print) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	dCount := strings.Count(print.str, "%d")
	if dCount != len(print.exprs) {
		errors = append(errors, "The number of parameters does not meet the printf format")
		return errors
	}
	// for _, expr := range print.exprs {
	// 	if expr.GetType(symTable, scope) != "int" {
	// 		errors = append(errors, "Only int can be printed in a format.")
	// 	}
	// }

	return errors
}
