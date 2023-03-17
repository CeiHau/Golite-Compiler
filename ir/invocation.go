package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
	"strings"
)

// <result> = call <ty> <fnval>(<function args>)
type Invocation struct {
	result int
	ty     string
	fnval  string
	args   []string
}

func NewInvocation(result int, ty string, fnval string, args []string) *Invocation {
	return &Invocation{result, ty, fnval, args}
}

// %r31 = call %struct.Point2D* @Init(i64 %r12)
func (instr *Invocation) String() string {

	var out bytes.Buffer
	targetReg := fmt.Sprintf("%%r%d", instr.result)
	out.WriteString(targetReg)
	out.WriteString(" = call ")
	out.WriteString(instr.ty)
	out.WriteString(" @")
	out.WriteString(instr.fnval)
	out.WriteString("(")
	for index, arg := range instr.args {
		out.WriteString(arg)
		if index != len(instr.args)-1 {
			out.WriteString(",")
		}
	}
	out.WriteString(")")

	return out.String()

}

func (instr *Invocation) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())

	for index, arg := range instr.args {
		argument := strings.Split(arg, " ")[1]
		if index < 8 {
			res = codegen.WriteCode(res, fmt.Sprintf("ldr x%d, [sp, #%v]", index, armTable[argument[1:]]))
		} else {
			fmt.Println("not implemented yet")
		}
	}
	rsltLocation := fmt.Sprintf("r%d", instr.result)
	res = codegen.WriteCode(res, fmt.Sprintf("bl %v", instr.fnval))
	res = codegen.WriteCode(res, fmt.Sprintf("str x0, [sp, #%v]", armTable[rsltLocation]))
	return res
}
