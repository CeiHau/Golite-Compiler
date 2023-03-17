package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

type Br struct {
	cond     *Operand
	iftrue   string
	iffalse  string
	condCode string
}

func NewBr(cond *Operand, iftrue string, iffalse string, condCode string) *Br {
	return &Br{cond, iftrue, iffalse, condCode}
}

// br label <dest>
// br label %L1

// br i1 <cond>, label <iftrue>, label <iffalse>
// br i1 %r28, label %L4, label %L5
func (instr *Br) String() string {
	var out bytes.Buffer
	if instr.cond == nil {
		out.WriteString("br label %" + instr.iftrue)
	} else {
		out.WriteString("br i1 " + instr.cond.String() + " label %" + instr.iftrue + ", label %" + instr.iffalse)
	}
	return out.String()
}

func (instr *Br) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	if instr.cond == nil {
		res = codegen.WriteCode(res, fmt.Sprintf("b .%v", instr.iftrue))
	} else {
		fmt.Println(instr.String(), "cond", instr.condCode)
		res = codegen.WriteCode(res, fmt.Sprintf("b.%v .%v", codegen.CondTable[instr.condCode], instr.iftrue))
		res = codegen.WriteCode(res, fmt.Sprintf("b .%v", instr.iffalse))
	}

	return res
}
