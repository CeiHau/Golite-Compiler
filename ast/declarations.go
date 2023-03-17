package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Declarations struct {
	*token.Token
	declarations []Declaration
}

func NewDeclarations(declarations []Declaration, token *token.Token) *Declarations {
	return &Declarations{token, declarations}
}

// // Because Declarations in Program and Declarations in Function have different usage in SymbolTable, we use table instead of tables
// func (declarations *Declarations) BuildSymbolTable(tables *st.SymbolTable, scope st.VarScope) {

// }

func (declarations *Declarations) DisplayAST(level int) {
	// var out bytes.Buffer
	// out.WriteString(strings.Repeat("   ", level))
	for _, decl := range declarations.declarations {
		decl.DisplayAST(level + 1)
	}
}
func (declarations *Declarations) String() string {
	var out bytes.Buffer
	for _, decl := range declarations.declarations {
		out.WriteString(decl.String())
	}
	return out.String()
}

func (declarations *Declarations) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	errors = append(errors, symTable.Globals.Errors...)
	if declarations != nil {
		for _, decl := range declarations.declarations {
			errors = decl.TypeCheck(errors, symTable, scope)
		}
	}
	return errors
}

func (declarations *Declarations) TranslateToLLVM(st *st.SymbolTables, scope string) *ir.BasicBlock {
	// var basicBlock ir.BasicBlock
	var instructions []ir.Instruction
	if declarations != nil {
		for _, dcl := range declarations.declarations {
			instructions = append(instructions, dcl.TranslateToLLVM(st, scope)...)
		}
	}

	label := ""
	return &ir.BasicBlock{label, instructions, nil, nil, ""}
}
