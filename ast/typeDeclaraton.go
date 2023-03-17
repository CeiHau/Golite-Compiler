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

type TypeDeclaration struct {
	*token.Token
	variable string
	fields   Fields
}

func NewTypeDeclaration(variable string, fields Fields, token *token.Token) *TypeDeclaration {
	return &TypeDeclaration{token, variable, fields}
}

func (typdcl *TypeDeclaration) BuildSymbolTable(tables *st.SymbolTables) {
	fields := make(map[string]*st.Pair)
	for ind, decl := range typdcl.fields.decls {
		fields[decl.variable] = &st.Pair{st.VarEntry{decl.variable, decl.ty, st.LOCAL, false}, ind}
	}
	tables.Structs.Insert(typdcl.variable, &st.StructEntry{typdcl.variable, fields, st.GLOBAL})
}
func (typdcl *TypeDeclaration) DisplayAST(level int) {
	var out bytes.Buffer
	out.WriteString(strings.Repeat("   ", level))
	out.WriteString("*ast.Types (Struct Name :")
	out.WriteString(typdcl.variable)
	out.WriteString(")")
	fmt.Println(out.String())
	typdcl.fields.DisplayAST(level)
}
func (typdcl *TypeDeclaration) String() string {
	var out bytes.Buffer
	out.WriteString("type")
	out.WriteString(typdcl.variable)
	out.WriteString("  struct {\n")
	out.WriteString(typdcl.fields.String())
	out.WriteString("}\n")
	return out.String()
}

func (typdcl *TypeDeclaration) TypeCheck(errors []string, st *st.SymbolTables, scope string) []string {
	// NOT exist
	if typdcl != nil {
		for _, field := range typdcl.fields.decls {
			fieldType := field.ty.String()
			if st.Structs.Contains(fieldType) == nil && fieldType != "bool" && fieldType != "int" {
				errors = append(errors, fieldType+" structs not declared")
			}
		}
	}
	return errors
}

func (typdcl *TypeDeclaration) TranslateToLLVM(tables *st.SymbolTables) []ir.Instruction {
	res := make([]ir.Instruction, 0)
	structEntry := tables.Structs.Contains(typdcl.variable)
	typeName := "%struct." + structEntry.Name
	var tys []string
	for _, field := range structEntry.Fields {
		tys = append(tys, field.Varentry.Ty.LLVM())
	}
	// Struct Declaration
	res = append(res, ir.NewTypeDeclaration(typeName, tys))
	// Struct nil
	res = append(res, ir.NewDeclaration(".nil"+structEntry.Name, typeName+"*", "null"))

	tables.Globals.Insert(".nil"+structEntry.Name, &st.VarEntry{".nil" + structEntry.Name, tp.StringToType(typeName + "*"), st.GLOBAL, false})
	return res
}
