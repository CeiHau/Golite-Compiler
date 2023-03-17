package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// Add represents a ADD instruction in ILOC
type Ne struct {
	result int // The target register for the instruction
	op     *Operand
}

func NewNe(result int, op *Operand) *Ne {
	return &Ne{result, op}
}

func (instr *Ne) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("%%r%d", instr.result)
	out.WriteString(targetReg)
	out.WriteString(" = fneg i64 ")
	out.WriteString(instr.op.String())

	return out.String()

}

func (instr *Ne) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	targetReg := fmt.Sprintf("r%d", instr.result)

	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	reg1 := codegen.NewRegister()
	reg2 := codegen.NewRegister()
	if instr.op.register {
		res = codegen.WriteCode(res, fmt.Sprintf("ldr %v, [sp, #%v]", reg1.Str, armTable[instr.op.String()[1:]]))
	} else {
		res = codegen.WriteCode(res, fmt.Sprintf("mov %v, #%v", reg1.Str, instr.op.String()))
	}

	res = codegen.WriteCode(res, fmt.Sprintf("vneg %v, %v", reg2.Str, reg1.Str))
	res = codegen.WriteCode(res, fmt.Sprintf("str %v, [sp, #%v]", reg2.Str, armTable[targetReg]))
	reg1.FreeRegister()
	reg2.FreeRegister()
	return res
}
