package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
	"strings"
)

var printreadNum int

type PR struct {
	pr   string //printf, read, scan
	name string
	len  int
	args []string
}

func InitPR() {
	printreadNum = 0
}

func NewPR(pr string, len int, args []string) *PR {
	res := &PR{pr, fmt.Sprintf("STR%d", printreadNum), len, args}
	if pr == "printf" {
		printreadNum++
	}
	return res
}

func (instr *PR) String() string {

	var out bytes.Buffer
	out.WriteString("call i32 (i8*, ...) ")
	if instr.pr == "printf" {
		out.WriteString(fmt.Sprintf("@printf(i8* getelementptr inbounds ([%d x i8], [%d x i8]* @.f%v, i32 0, i32 0)", instr.len, instr.len, instr.name))
	} else {
		out.WriteString("@scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0)")
	}
	for _, arg := range instr.args {
		out.WriteString(", ")
		out.WriteString(arg)

	}
	out.WriteString(")")

	return out.String()

}

func (instr *PR) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	res := make([]codegen.Instruction, 0)
	res = codegen.WriteCode(res, "// "+instr.String())
	reg1 := codegen.NewRegister()
	if instr.pr == "printf" {

		res = codegen.WriteCode(res, fmt.Sprintf("adrp %v, .FMT%v", reg1.Str, instr.name))
		res = codegen.WriteCode(res, fmt.Sprintf("add %v, %v, :lo12:.FMT%v", reg1.Str, reg1.Str, instr.name))
		res = codegen.WriteCode(res, "mov x0, "+reg1.Str)
		for i, arg := range instr.args {
			loc := strings.Split(arg, " ")[1]
			if strings.HasPrefix(loc, "%") {
				res = codegen.WriteCode(res, fmt.Sprintf("ldr x%d, [sp, #%v]", i+1, armTable[loc[1:]]))
			} else {
				res = codegen.WriteCode(res, fmt.Sprintf("mov x%d, #%v", i+1, loc))
			}
		}

		res = codegen.WriteCode(res, "bl printf")
	} else {
		var loc string
		var global bool
		arg := strings.Split(instr.args[0], " ")[1]
		res = codegen.WriteCode(res, fmt.Sprintf("adrp %v, .READ", reg1.Str))
		res = codegen.WriteCode(res, fmt.Sprintf("add %v, %v, :lo12:.READ", reg1.Str, reg1.Str))
		res = codegen.WriteCode(res, fmt.Sprintf("mov x0, %v", reg1.Str))
		res, loc, global = codegen.GetVariableAddress(res, arg, armTable)
		if global {
			res = codegen.WriteCode(res, fmt.Sprintf("add x1, %v, #0", loc[1:len(loc)-1]))
		} else {
			res = codegen.WriteCode(res, fmt.Sprintf("add x1, %v", loc[1:len(loc)-1]))
		}

		res = codegen.WriteCode(res, "bl scanf")
	}
	reg1.FreeRegister()
	return res
}

// @.fstr1 = private unnamed_addr constant [32 x i8] c"offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00", align 1
type PRLL struct {
	pr   string
	name string //printf, read, scan
	len  int
	args string
}

func NewPRLL(printread *PR, l int, args string) *PRLL {
	name := ""
	if printread.pr == "printf" {
		name = printread.name
	}
	args = strings.Replace(args, "%d", "%ld", -1)
	args = strings.Replace(args, "\\n", "\\0A", -1)
	args = args[:len(args)-1] + "\\00\""
	return &PRLL{printread.pr, name, l, args}
}

func (instr *PRLL) String() string {

	var out bytes.Buffer
	if instr.pr == "printf" {
		out.WriteString(fmt.Sprintf("@.f%v = private unnamed_addr constant [%d x i8] c%v, align 1", instr.name, instr.len, instr.args))
	} else {
		out.WriteString(fmt.Sprintf("@.read = private unnamed_addr constant [%d x i8] c%v, align 1", instr.len, instr.args))
	}
	return out.String()

}

func (instr *PRLL) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {

	return nil
}
