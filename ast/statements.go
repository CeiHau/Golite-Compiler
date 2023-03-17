package ast

import (
	"bytes"
	"golite/token"
)

type Statements struct {
	*token.Token
	statements []Statement
}

func NewStatements(statements []Statement, token *token.Token) *Statements {
	return &Statements{token, statements}
}

func (statements Statements) String() string {
	var out bytes.Buffer
	for _, statement := range statements.statements {
		out.WriteString(statement.String())
	}
	return out.String()
}
