package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// <result> = load <ty>, <ty>* <pointer>
type Load struct {
	result  int // The target register for the instruction
	ty      string
	pointer string // pointer
}

// NewAdd is a constructor and initialization function for a new Add instruction
func NewLoad(result int, ty string, pointer string) *Load {
	return &Load{result, ty, pointer}
}

func (instr *Load) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("%%r%d", instr.result)
	out.WriteString(targetReg)
	out.WriteString(" = load " + instr.ty + ", " + instr.ty + "* ")
	out.WriteString(instr.pointer)

	return out.String()

}

func (instr *Load) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {

	targetReg := fmt.Sprintf("r%d", instr.result)

	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	var pointerLocation string

	res, pointerLocation, _ = codegen.GetVariableAddress(res, instr.pointer, armTable)
	reg := codegen.NewRegister()
	res = codegen.WriteCode(res, fmt.Sprintf("ldr %v, %v", reg.Str, pointerLocation))

	res = codegen.WriteCode(res, fmt.Sprintf("str %v, [sp, #%v]", reg.Str, armTable[targetReg]))

	reg.FreeRegister()
	return res
}
