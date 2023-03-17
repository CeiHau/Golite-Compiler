package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// Sdiv represents a SDIV instruction in ILOC
type Sdiv struct {
	result int      // The target register for the instruction
	op1    *Operand // op1, operands from registers or immediate integer values
	op2    *Operand // op2, operands from registers or immediate integer values
}

// NewSdiv is a constructor and initialization function for a new Sdiv instruction
func NewSdiv(result int, op1 *Operand, op2 *Operand) *Sdiv {
	return &Sdiv{result, op1, op2}
}

func (instr *Sdiv) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("%%r%d", instr.result)
	out.WriteString(targetReg)
	out.WriteString(" = sdiv i64 ")
	out.WriteString(instr.op1.String())
	out.WriteString(", ")
	out.WriteString(instr.op2.String())

	return out.String()

}

func (instr *Sdiv) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	targetReg := fmt.Sprintf("r%d", instr.result)
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	reg1 := codegen.NewRegister()
	reg2 := codegen.NewRegister()
	reg3 := codegen.NewRegister()
	if instr.op1.register {
		res = codegen.WriteCode(res, fmt.Sprintf("ldr %v, [sp, #%v]", reg1.Str, armTable[instr.op1.String()[1:]]))
	} else {
		res = codegen.WriteCode(res, fmt.Sprintf("mov %v, #%v", reg1.Str, instr.op1.String()))
	}

	if instr.op2.register {
		res = codegen.WriteCode(res, fmt.Sprintf("ldr %v, [sp, #%v]", reg2.Str, armTable[instr.op2.String()[1:]]))
	} else {
		res = codegen.WriteCode(res, fmt.Sprintf("mov %v, #%v", reg2.Str, instr.op2.String()))
	}

	res = codegen.WriteCode(res, fmt.Sprintf("sdiv %v, %v, %v", reg3.Str, reg2.Str, reg1.Str))
	res = codegen.WriteCode(res, fmt.Sprintf("str %v, [sp, #%v]", reg3.Str, armTable[targetReg]))
	reg1.FreeRegister()
	reg2.FreeRegister()
	reg3.FreeRegister()
	return res
}
