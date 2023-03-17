package ast

import (
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type Variable struct {
	*token.Token
	Identifier string // The variable name
	ty         tp.Type
}

func NewVariable(identifier string, token *token.Token) *Variable {
	return &Variable{token, identifier, nil}
}
func (v *Variable) String() string { return v.Identifier }

func (v *Variable) GetType() tp.Type {
	return v.ty
}
func (v *Variable) TypeCheck(errors []string, tables *st.SymbolTables) []string {

	if varEntry := tables.Globals.Contains(v.Identifier); varEntry != nil {
		//Your found the type so you want to save it in the type for the expression
		v.ty = varEntry.Ty
	} else {
		//Add an error to the errors array because you did not find the type
		//You might want to assign the v.ty to an unknwon type so you know that it's not equal to an actual type.
		errors = append(errors, "type not found")
	}
	return errors
}
