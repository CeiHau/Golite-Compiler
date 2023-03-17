package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type NewID struct {
	*token.Token
	varibale string
	Usable   bool
}

func NewNewID(variable string, token *token.Token) *NewID {
	return &NewID{token, variable, true}
}

func (ni *NewID) String() string {
	var out bytes.Buffer
	out.WriteString("new ")
	out.WriteString(ni.varibale)
	return out.String()
}
func (ni *NewID) GetType(symTable *st.SymbolTables, scope string) string {
	return ni.varibale
}

func (fi *NewID) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	structEntry := tables.Structs.Contains(fi.varibale)
	if structEntry == nil {
		errors = append(errors, fi.varibale+" struct not defined")
	}
	return errors
}

// %r3 = call i8* @malloc(i32 16)
// %r4 = bitcast i8* %r3 to %struct.Point2D*
func (fi *NewID) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	rslt1 := ir.NewRegister(tables, block.FuncName)
	// Add malloc instruction to the block
	structEntry := tables.Structs.Contains(fi.varibale)
	block.Instructions = append(block.Instructions, ir.NewNewAlloc(rslt1.GetRegisterNum(), len(structEntry.Fields)*8))
	// Add bitcast instruction to the block
	rslt2 := ir.NewRegister(tables, block.FuncName)
	block.Instructions = append(block.Instructions, ir.NewBitcast(rslt2.GetRegisterNum(), "i8*", rslt1.GetRegisterNum(), tp.StringToType(fi.GetType(tables, scope)).LLVM()))

	// Return the Operands
	return rslt2
}
