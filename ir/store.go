package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
)

// store <ty> <value>, <ty>* <pointer>
type Store struct {
	ty      string
	value   string
	pointer string // pointer
}

// NewAdd is a constructor and initialization function for a new Add instruction
func NewStore(ty string, value string, pointer string) *Store {
	return &Store{ty, value, pointer}
}

func (instr *Store) String() string {

	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("store %v %v, %v* %v ", instr.ty, instr.value, instr.ty, instr.pointer))

	return out.String()

}

func (instr *Store) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	valueLocation := armTable[instr.value[1:]]
	var pointerLocation string
	reg := codegen.NewRegister()
	if instr.value[0] == '%' { // Store in register or stack
		if valueLocation[0] == 'x' { // Store from register
			pointerLocation = armTable[instr.pointer[1:]]
			res = codegen.WriteCode(res, fmt.Sprintf("str %v, [sp, #%v]", valueLocation, pointerLocation))
			reg.FreeRegister()
			return res
		} else { // Store from stack, need temporary register, use ldr
			res = codegen.WriteCode(res, fmt.Sprintf("ldr %v, [sp, #%v]", reg.Str, valueLocation))
		}

	} else { // Store immediate number, use mov
		res = codegen.WriteCode(res, fmt.Sprintf("mov %v, %v", reg.Str, instr.value))

	}
	res, pointerLocation, _ = codegen.GetVariableAddress(res, instr.pointer, armTable)
	res = codegen.WriteCode(res, fmt.Sprintf("str %v, %v", reg.Str, pointerLocation))
	reg.FreeRegister()

	return res
}
