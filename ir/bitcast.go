package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// bitcast represents a ADD instruction in ILOC
type Bitcast struct {
	result int // The target register for the instruction
	ty     string
	value  int
	ty2    string
}

func NewBitcast(result int, ty string, value int, ty2 string) *Bitcast {
	return &Bitcast{result, ty, value, ty2}
}

// %r4 = bitcast i8* %r3 to %struct.Point2D*
// <result> = bitcast <ty> <value> to <ty2>
func (instr *Bitcast) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("%%r%d", instr.result)
	out.WriteString(targetReg)
	out.WriteString(" = bitcast ")
	out.WriteString(instr.ty)
	sourceReg := fmt.Sprintf(" %%r%d", instr.value)
	out.WriteString(sourceReg)
	out.WriteString(" to ")
	out.WriteString(instr.ty2)

	return out.String()

}

func (instr *Bitcast) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	return res
}
