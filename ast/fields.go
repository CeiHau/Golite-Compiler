package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

type Fields struct {
	*token.Token
	decls []Decl
}

func NewFields(decls []Decl, token *token.Token) *Fields {
	return &Fields{token, decls}
}
func (fields *Fields) DisplayAST(level int) {
	for index, decl := range fields.decls {
		decl.DisplayAST(index, level+1)
	}
}
func (fields *Fields) String() string {
	var out bytes.Buffer
	for _, decl := range fields.decls {
		out.WriteString(decl.String())
		out.WriteString(";\n")
	}

	return out.String()
}
func (ast *Fields) TypeCheck(errors []string, st *st.SymbolTables, scope string) []string {
	return errors
}
