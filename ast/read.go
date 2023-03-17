package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Read struct {
	*token.Token
	lValue LValue
}

func NewRead(lValue LValue, token *token.Token) *Read {
	return &Read{token, lValue}
}

func (reads *Read) String() string {
	var out bytes.Buffer
	out.WriteString("scan ")
	out.WriteString(reads.lValue.String())
	out.WriteString(";\n")
	return out.String()
}

func (reads *Read) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	return errors
}
func (reads *Read) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}
func (reads *Read) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {
	var args []string
	args = append(args, "i64* @"+reads.lValue.id[0])
	printread := ir.NewPR("scanf", 4, args)
	block.Instructions = append(block.Instructions, printread)
	ir.GlobalDeclare.Instructions = append(ir.GlobalDeclare.Instructions, ir.NewPRLL(printread, 4, "%d"))
	return block
}
