package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// <result> = getelementptr <ty>, <ty>* <ptrval>, i32 0, i32 <index>
type StrAcc struct {
	result int // The target register for the instruction
	ty     string
	ptrval int
	index  int
}

func NewStrAcc(result int, ty string, ptrval int, index int) *StrAcc {
	return &StrAcc{result, ty, ptrval, index}
}

func (instr *StrAcc) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("%%r%d", instr.result)
	out.WriteString(targetReg)
	out.WriteString(" = getelementptr ")
	out.WriteString(instr.ty[:len(instr.ty)-1] + ", ")
	out.WriteString(instr.ty + " %r")
	out.WriteString(fmt.Sprintf("%d", instr.ptrval))
	out.WriteString(", i32 0, i32 ")
	out.WriteString(fmt.Sprintf("%d", instr.index))

	return out.String()

}

func (instr *StrAcc) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
