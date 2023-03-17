package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
)

type Ids struct {
	*token.Token
	variables []string
}

func NewIds(variables []string, token *token.Token) *Ids {
	return &Ids{token, variables}

}

func (ids *Ids) String() string {
	var out bytes.Buffer
	for index, variable := range ids.variables {
		out.WriteString(variable)
		if index != len(ids.variables)-1 {
			out.WriteString(", ")
		}
	}
	return out.String()
}
func (ast *Ids) GetType(symTable *st.SymbolTables, scope string) string {
	curScope := symTable.Funcs.Contains(scope)
	varEntry := curScope.Variables.Contains(ast.variables[len(ast.variables)-1])
	return varEntry.Ty.String()
}
