package ast

import (
	"bytes"
	"fmt"
	"golite/token"
	tp "golite/type"
	"strconv"
	"strings"
)

type Decl struct {
	*token.Token
	variable string
	ty       tp.Type
}

func NewDecl(variable string, ty tp.Type, token *token.Token) *Decl {
	return &Decl{token, variable, ty}
}

func (decl *Decl) DisplayAST(index int, level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Field (Struct Variable ")
	out.WriteString(strconv.Itoa(index))
	out.WriteString(") (Varaiable Name: ")
	out.WriteString(decl.variable)
	out.WriteString(", Variable Type: ")
	out.WriteString(decl.ty.String())
	out.WriteString(")")
	fmt.Println(out.String())
}

func (decl *Decl) String() string {
	var out bytes.Buffer
	out.WriteString(decl.variable)
	out.WriteString(" ")
	out.WriteString(decl.ty.String())
	// out.WriteString(";\n")
	return out.String()
}
