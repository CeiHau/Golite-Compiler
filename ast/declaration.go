package ast

import (
	"bytes"
	"fmt"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
	"strings"
)

type Declaration struct {
	*token.Token
	ids Ids
	ty  tp.Type
}

func NewDeclaration(ids Ids, ty tp.Type, token *token.Token) *Declaration {
	return &Declaration{token, ids, ty}
}

func (decl *Declaration) BuildSymbolTable(tables *st.SymbolTables, scope st.VarScope) {
	// for _, id := range decl.ids.variables {
	// 	table.Insert(id, &st.VarEntry{id, decl.ty, scope})
	// }
}

func (decl *Declaration) DisplayAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Declaration (Variable Name: ")

	// decl.ids.DisplayAST(level)
	out.WriteString(" Variable Type: ")

	if decl.ty.String() != "int" || decl.ty.String() != "bool" {
		out.WriteString(" *" + decl.ty.String())
	} else {
		out.WriteString(" " + decl.ty.String())
	}

	out.WriteString(")")
	fmt.Println(out.String())
}
func (decl *Declaration) String() string {
	var out bytes.Buffer
	out.WriteString("var ")
	out.WriteString(decl.ids.String())
	// if decl.ty.String() != "int" || decl.ty.String() != "bool" {
	// 	out.WriteString(" *" + decl.ty.String())
	// } else {
	// 	out.WriteString(" " + decl.ty.String())
	// }
	out.WriteString(" " + decl.ty.String())
	out.WriteString(";\n")
	return out.String()
}

func (decl *Declaration) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	// not exists in symboltable
	if decl != nil {
		fieldType := decl.ty.String()
		if symTable.Structs.Contains(fieldType) == nil && fieldType != "bool" && fieldType != "int" {
			errors = append(errors, fieldType+" structs not declared")
		}
	}
	return errors

}

// @VARIABLE_NAME = common global <ty> <constant value>
// @globalInit = common global i64 0
// @ptr = common global %struct.Point2D* null
func (decl *Declaration) TranslateToLLVM(st *st.SymbolTables, scope string) []ir.Instruction {
	res := make([]ir.Instruction, 0)
	if decl != nil {
		fieldType := decl.ty.String()
		ty := decl.ty.LLVM()

		for _, id := range decl.ids.variables {
			if scope == "global" {
				if fieldType == "bool" || fieldType == "int" {
					res = append(res, ir.NewDeclaration(id, ty, "0"))
				} else {
					res = append(res, ir.NewDeclaration(id, ty, "null"))
				}
			} else {
				res = append(res, ir.NewAlloca(id, ty))
			}
		}
	}
	return res
}
