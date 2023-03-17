package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

type Delete struct {
	value int
}

func NewDelete(value int) *Delete {
	return &Delete{value}
}

// call void @free(i8* %r23)
func (instr *Delete) String() string {

	var out bytes.Buffer
	out.WriteString("call void @free(i8* ")
	sourceReg := fmt.Sprintf("%%r%d", instr.value)
	out.WriteString(sourceReg)
	out.WriteString(")")

	return out.String()

}

func (instr *Delete) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	res := make([]codegen.Instruction, 0)
	valueLocation := armTable[fmt.Sprintf("r%d", instr.value)]
	res = codegen.WriteCode(res, fmt.Sprintf("str x0, [sp, #%v]", valueLocation))
	res = codegen.WriteCode(res, "bl free")

	return res
}
