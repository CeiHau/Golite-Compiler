package ir

import (
	"bytes"
	"golite/codegen"
	st "golite/symboltable"
)

type TypeDeclaration struct {
	typeName string
	tys      []string
}

func NewTypeDeclaration(typeName string, tys []string) *TypeDeclaration {
	return &TypeDeclaration{typeName, tys}
}

func (instr *TypeDeclaration) String() string {
	var out bytes.Buffer
	out.WriteString(instr.typeName)
	out.WriteString(" = type{")
	for index, ty := range instr.tys {
		out.WriteString(ty)
		if index != len(instr.tys)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString("}")
	return out.String()
}

func (instr *TypeDeclaration) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
