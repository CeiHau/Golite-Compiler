package codegen

import (
	"fmt"
	"strconv"
)

type Instruction string

// type Instruction interface {
// }

// func TranslateToAssembly(blocks []*ir.BasicBlock, tables *st.SymbolTables) {

// }

func WriteCode(ins []Instruction, s string) []Instruction {
	if s[0] == '/' && s[1] == '/' {
		s = "\n" + s
	}
	ins = append(ins, Instruction("        "+s+"\n"))
	return ins
}

/*************** Register code ***************/
var registerTable map[int]bool

type Register struct {
	Num int
	Str string
}

func InitRegister() {
	registerTable = map[int]bool{}
	for r := 8; r <= 15; r++ {
		registerTable[r] = true
	}
}

func NewRegister() *Register {
	for r := 8; r <= 15; r++ {
		if registerTable[r] == true {
			registerTable[r] = false

			return &Register{r, "x" + strconv.Itoa(r)}
		}
	}
	return &Register{-1, "don't have enough register"}
}

func (r *Register) FreeRegister() {
	registerTable[r.Num] = true
}

/*************** Loading and Storing Global Variables ***************/

func GetVariableAddress(res []Instruction, variable string, armTable map[string]string) ([]Instruction, string, bool) {
	if variableLocation, ok := armTable[variable[1:]]; ok { // Local variable
		return res, fmt.Sprintf("[sp, #%v]", variableLocation), false
	} else { // Global variable
		reg := NewRegister()
		res = WriteCode(res, fmt.Sprintf("adrp %v, %v", reg.Str, variable[1:]))
		res = WriteCode(res, fmt.Sprintf("add %v, %v, :lo12:%v", reg.Str, reg.Str, variable[1:]))
		reg.FreeRegister()
		return res, fmt.Sprintf("[%v]", reg.Str), true
	}
}

var CondTable map[string]string

func InitCondTable() {
	CondTable = map[string]string{
		"eq":  "eq",
		"ne":  "ne",
		"sgt": "gt",
		"sge": "ge",
		"slt": "lt",
		"sle": "le",
	}
}
