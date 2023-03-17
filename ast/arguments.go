// modified
package ast

import (
	"bytes"
	"golite/token"
)

type Arguments struct {
	*token.Token
	expressions []Expression
}

func NewArguments(expressions []Expression, token *token.Token) *Arguments {
	return &Arguments{
		Token:       token,
		expressions: expressions,
	}
}

func (arguments Arguments) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	for index, expr := range arguments.expressions {
		out.WriteString(expr.String())
		if index != len(arguments.expressions)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")
	return out.String()
}
