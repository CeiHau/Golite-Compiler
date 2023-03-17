package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

type NewAlloc struct {
	result int
	mem    int
}

func NewNewAlloc(result int, mem int) *NewAlloc {
	return &NewAlloc{result, mem}
}

// %r3 = call i8* @malloc(i32 16)
func (instr *NewAlloc) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("%%r%d", instr.result)
	out.WriteString(targetReg)
	out.WriteString(" = call")
	out.WriteString(NewMalloc(instr.mem).String())

	return out.String()

}

func (instr *NewAlloc) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	res = codegen.WriteCode(res, fmt.Sprintf("mov x0, #%d", instr.mem))
	res = codegen.WriteCode(res, "bl malloc")
	valueLocation := armTable[fmt.Sprintf("r%d", instr.result+1)]
	res = codegen.WriteCode(res, fmt.Sprintf("str x0, [sp, #%v]", valueLocation))
	return res
}
