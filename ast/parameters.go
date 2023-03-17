package ast

import (
	"bytes"
	"golite/token"
)

type Parameters struct {
	*token.Token
	decls []Decl
}

func NewParameters(decls []Decl, token *token.Token) *Parameters {
	return &Parameters{token, decls}
}

func (parameters Parameters) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	for index, decl := range parameters.decls {
		out.WriteString(decl.String())
		if index != len(parameters.decls)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(") ")
	return out.String()
}
