package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

type FunctionDefinition struct {
	funcName  string
	tys       []string
	parmNames []string
	ret_ty    string
}

func NewFunctionDefinition(funcName string, tys []string, parmNames []string, ret_ty string) *FunctionDefinition {
	return &FunctionDefinition{funcName, tys, parmNames, ret_ty}
}

func (instr *FunctionDefinition) String() string {
	var out bytes.Buffer
	out.WriteString("\ndefine " + instr.ret_ty + " @" + instr.funcName + "(")
	for index, ty := range instr.tys {
		out.WriteString(ty + "% " + instr.parmNames[index])
		if index != len(instr.tys)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(")\n{")

	return out.String()
}

func (instr *FunctionDefinition) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	// todo: offset find
	offset := (len(armTable)-len(instr.parmNames))*8 + 16
	if offset%16 != 0 {
		offset += 8
	}
	res := make([]codegen.Instruction, 0)
	// res = codegen.WriteCode(res, fmt.Sprintf(".type %v, %%function", instr.funcName))
	// res = codegen.WriteCode(res, ".p2align 2")
	// res = append(res, codegen.Instruction(fmt.Sprintf("%v:\n", instr.funcName)))
	// Prologue
	res = codegen.WriteCode(res, "// Prologue")
	res = codegen.WriteCode(res, fmt.Sprintf("sub sp, sp, #%d", offset))
	res = codegen.WriteCode(res, "stp x29, x30, [sp]")
	res = codegen.WriteCode(res, "mov x29, sp")

	return res
}
