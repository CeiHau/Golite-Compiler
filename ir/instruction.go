package ir

import (
	"bytes"
	"fmt"
	"golite/codegen"
	st "golite/symboltable"
	"os"
	"strconv"
	"strings"
)

type Instruction interface {
	// GetTargets() []int // Get the registers targeted by this instruction

	// GetSources() []int // Get the source registers for this instruction

	// GetImmediate() *int // Get the immediate value (i.e., constant) of this instruction

	// GetLabel() string // Get the label for this instruction

	// SetLabel(newLabel string) //Set the label for this instruction

	String() string // Return a string representation of this instruction

	TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction
}

type BasicBlock struct {
	Label        string        // represents the label for
	Instructions []Instruction // represents the instructions for this block
	Preds        []*BasicBlock
	Succs        []*BasicBlock
	FuncName     string
}

func (block *BasicBlock) DisplayLLVM() (string, []Instruction) {
	// BFS Search
	var out bytes.Buffer
	var instructions []Instruction
	blockStack := []*BasicBlock{}
	blockStack = append(blockStack, block)
	visited := make(map[string]bool)
	for len(blockStack) != 0 {
		b := blockStack[0]
		blockStack = blockStack[1:]
		if _, visit := visited[b.Label]; visit && b.Label != "" && b.Label != "functionDef" {
			continue
		}
		visited[b.Label] = true

		if len(b.Label) != 0 && b.Label != "functionDef" {
			instructions = append(instructions, LabelInstruction())
			out.WriteString(b.Label + ":\n")
		}

		for _, instruction := range b.Instructions {
			if len(b.Label) != 0 && b.Label != "functionDef" {

				out.WriteString(strings.Repeat("    ", 1))
			}
			instructions = append(instructions, instruction)
			out.WriteString(instruction.String() + "\n")
		}
		if b.Succs != nil {
			blockStack = append(blockStack, b.Succs...)
		}
	}
	return out.String(), instructions
}

func TranslateToAssembly(blocks []*BasicBlock, tables *st.SymbolTables, filename string) {
	f, err := os.Create(filename + ".s")
	if err != nil {
		panic(err)
	}
	codegen.InitRegister()
	codegen.InitCondTable()
	// var out bytes.Buffer
	var out *bytes.Buffer
	var arch, globals bytes.Buffer

	// architecture
	arch.WriteString("        .arch armv8-a\n")
	f.WriteString(arch.String())

	// globals
	for _, entry := range tables.Globals.GetTable() {
		globals.WriteString("        .comm " + entry.Name + ",8,8\n")
	}
	globals.WriteString("        .text\n")
	f.WriteString(globals.String())

	for _, block := range blocks {
		// fmt.Print(index)
		out = block.TranslateToAssembly(tables)
		f.WriteString(out.String())
	}
	res := make([]codegen.Instruction, 0)
	for _, instr := range GlobalDeclare.Instructions {

		switch instr.(type) {
		case *PRLL:
			prll := instr.(*PRLL)
			if prll.pr == "printf" {
				label := fmt.Sprintf(".FMT%v", strings.ToUpper(prll.name))
				res = append(res, codegen.Instruction(label+":\n"))
				res = codegen.WriteCode(res, fmt.Sprintf(".asciz %v", prll.args))
				res = codegen.WriteCode(res, fmt.Sprintf(".size %v, %d", label, prll.len))
			} else {
				res = append(res, ".READ:\n")
				res = codegen.WriteCode(res, ".asciz \"%ld\\00\"")
				res = codegen.WriteCode(res, ".size .READ, 4")
			}

		}

	}
	for _, str := range res {
		out.WriteString(string(str))
	}
	f.WriteString(out.String())
}

func (block *BasicBlock) TranslateToAssembly(tables *st.SymbolTables) *bytes.Buffer {
	var out bytes.Buffer
	fmt.Println(block.FuncName)
	blockStack := []*BasicBlock{}
	blockStack = append(blockStack, block)
	visited := make(map[string]bool)
	var offset int
	var armTable map[string]string
	var firstLabel string
	if block.Label == "functionDef" { // The begin block of new function, create table
		/************ Create Arm Table ************/
		funcName := block.Instructions[0].(*FunctionDefinition).funcName
		funcEntry := tables.Funcs.Contains(funcName)
		// Initialize the arm table and offset
		armTable = make(map[string]string)
		armTable["_retval"] = "16"
		offset = 24

		// Map local variable
		for _, v := range funcEntry.Variables.GetTable() {
			armTable[v.Name] = strconv.Itoa(offset)
			offset += 8
		}
		// Epilogue = (len(armTable)-len(funcEntry.Parameters))*8 + 16
		// fmt.Println("offset", Epilogue, len(armTable), len(funcEntry.Parameters))

		// Map arguments
		for argReg, arg := range funcEntry.Parameters {
			if argReg < 8 { // The first 8 arguments to a function must be passed to registers.
				armTable[arg.Name] = "x" + strconv.Itoa(argReg)
				argReg++
			} else { //Additional arguments are passed on the stack.
				fmt.Println("not implemented yet")
			}

		}

		// Map temporary variable
		for _, r := range funcEntry.Register {
			armTable[r[1:]] = strconv.Itoa(offset)
			offset += 8
		}

		// Map parameter
		for _, para := range funcEntry.Parameters {
			armTable["P_"+para.Name] = strconv.Itoa(offset)
			offset += 8
		}

		/************ Generate arm code head ************/
		firstLabel = block.Succs[0].Label
		head := make([]codegen.Instruction, 0)
		head = codegen.WriteCode(head, fmt.Sprintf(".type %v, %%function", funcName))
		head = codegen.WriteCode(head, fmt.Sprintf(".global %v", funcName))
		head = codegen.WriteCode(head, ".p2align 2")
		head = append(head, codegen.Instruction(fmt.Sprintf("%v:\n", funcName)))
		head = append(head, codegen.Instruction(fmt.Sprintf(".%v:\n", firstLabel)))
		for _, instr := range head {
			out.WriteString(string(instr))
		}

	}
	// Display the armTbale, temporary use
	for k, v := range armTable {
		fmt.Println(k, ": ", v)
	}

	for len(blockStack) != 0 {
		b := blockStack[0]
		blockStack = blockStack[1:]
		if _, visit := visited[b.Label]; visit && b.Label != "" {
			continue
		}
		visited[b.Label] = true
		if b.Label != "functionDef" && b.Label != "" && b.Label != firstLabel {
			out.WriteString("." + b.Label + ":\n")
		}

		for _, instruction := range b.Instructions {
			for _, code := range instruction.TranslateToAssembly(*tables, armTable) {
				out.WriteString(string(code))
			}

		}
		if b.Succs != nil {
			blockStack = append(blockStack, b.Succs...)
		}
	}
	if block.Label == "functionDef" {
		// Epilogue
		if offset%16 != 0 {
			offset += 8
		}
		out.WriteString("// Epilogue\n")
		out.WriteString("        ldp x29, x30, [sp]\n")
		out.WriteString(fmt.Sprintf("        add sp, sp, #%d\n", offset))
		out.WriteString("        ret\n")
		out.WriteString("        .size " + block.FuncName + ", (.-" + block.FuncName + ")\n")
	}
	return &out
}

var GlobalDeclare = &BasicBlock{"", nil, nil, nil, ""}

func InitGlobalDeclare() {
	GlobalDeclare.Instructions = append(GlobalDeclare.Instructions, NewGeneral(""))
	GlobalDeclare.Instructions = append(GlobalDeclare.Instructions, NewGeneral("declare i8* @malloc(i32)"))
	GlobalDeclare.Instructions = append(GlobalDeclare.Instructions, NewGeneral("declare void @free(i8*)"))
	GlobalDeclare.Instructions = append(GlobalDeclare.Instructions, NewGeneral("declare i32 @printf(i8*, ...)"))
	GlobalDeclare.Instructions = append(GlobalDeclare.Instructions, NewGeneral("declare i32 @scanf(i8*, ...)"))

}

type Operand struct {
	val      int
	register bool
}

func NewOperand(val int, register bool) *Operand {
	return &Operand{val, register}
}
func (operand Operand) String() string {

	if operand.register {
		return fmt.Sprintf("%%r%d", operand.val)
	} else {
		return fmt.Sprintf("%d", operand.val)
	}
}

func (operand Operand) GetRegisterNum() int {
	return operand.val

}

type RightBrace struct {
}

func (rb *RightBrace) String() string {
	return "}"
}

func (instr *RightBrace) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
