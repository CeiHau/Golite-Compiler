package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// ret void
// ret <ty> <value>
// ret i64 %r41
// ret %struct.Point2D* %r40
// ret represents a ret instruction in ILOC
type Ret struct {
	value int // The source register for the instruction
	ty    string
}

// NewAdd is a constructor and initialization function for a new Add instruction
func NewRet(value int, ty string) *Ret {
	return &Ret{value, ty}
}

func (instr *Ret) String() string {

	var out bytes.Buffer
	out.WriteString("ret ")
	out.WriteString(instr.ty)
	if instr.value != -1 {
		sourceReg := fmt.Sprintf(" %%r%d", instr.value)
		out.WriteString(sourceReg)
	}
	return out.String()

}

func (instr *Ret) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	rsltReg := fmt.Sprintf("r%d", instr.value)
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	res = codegen.WriteCode(res, fmt.Sprintf("ldr x0, [sp, #%v]", armTable[rsltReg]))
	return res
}
