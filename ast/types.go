package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Types struct {
	*token.Token
	typeDeclarations []TypeDeclaration
}

func NewTypes(typeDeclarations []TypeDeclaration, token *token.Token) *Types {
	return &Types{token, typeDeclarations}
}
func (tps *Types) BuildSymbolTable(tables *st.SymbolTables) {
	if tps == nil || len(tps.typeDeclarations) == 0 {
		return
	}
	for _, tpDcl := range tps.typeDeclarations {
		tpDcl.BuildSymbolTable(tables)
	}
}
func (tps *Types) DisplayAST(level int) {
	// var out bytes.Buffer
	// out.WriteString(strings.Repeat("   ", level))
	for _, tpDcl := range tps.typeDeclarations {
		tpDcl.DisplayAST(level + 1)
	}
}

func (tps *Types) String() string {
	var out bytes.Buffer
	for _, tp := range tps.typeDeclarations {
		out.WriteString(tp.String())
	}
	return out.String()
}
func (tps *Types) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	errors = append(errors, tables.Structs.Errors...)
	if tps != nil {
		for _, tp := range tps.typeDeclarations {
			errors = tp.TypeCheck(errors, tables, scope)
		}
	}

	return errors
}

func (tps *Types) TranslateToLLVM(st *st.SymbolTables) *ir.BasicBlock {
	// var basicBlock ir.BasicBlock
	var instructions []ir.Instruction
	if tps != nil {
		for _, typdcl := range tps.typeDeclarations {
			instructions = append(instructions, typdcl.TranslateToLLVM(st)...)
		}
	}

	label := ""

	return &ir.BasicBlock{label, instructions, nil, nil, ""}
}
