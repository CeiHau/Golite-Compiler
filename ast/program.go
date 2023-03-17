package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Program struct {
	*token.Token
	types        *Types
	declarations *Declarations
	functions    *Functions
}

func NewProgram(declarations *Declarations, types *Types, functions *Functions, token *token.Token) *Program {
	return &Program{token, types, declarations, functions}

}

func (p *Program) BuildSymbolTable(tables *st.SymbolTables) {
	if p.types != nil {
		p.types.BuildSymbolTable(tables)
	}

	// p.declarations.BuildSymbolTable[*st.VarEntry](tables.Globals, st.GLOBAL)
	if p.declarations != nil {
		for _, declaration := range p.declarations.declarations {
			for _, variable := range declaration.ids.variables {
				tables.Globals.Insert(variable, &st.VarEntry{variable, declaration.ty, st.GLOBAL, false})
			}
		}
	}
	if p.functions != nil {
		p.functions.BuildSymbolTable(tables)
	}
}

func (p *Program) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	errors = p.declarations.TypeCheck(errors, tables, "GLOBAL")
	errors = p.types.TypeCheck(errors, tables, "GLOBAL")
	errors = p.functions.TypeCheck(errors, tables)
	return errors
}

func (p *Program) String() string {
	var out bytes.Buffer
	if p.types != nil {
		out.WriteString(p.types.String())
	}
	if p.declarations != nil {
		out.WriteString(p.declarations.String())
	}
	if p.functions != nil {
		out.WriteString(p.functions.String())
	}
	return out.String()
}

func (p *Program) TranslateToLLVM(tables *st.SymbolTables) []*ir.BasicBlock {
	ir.InitPR()
	ir.InitGlobalDeclare()
	basicBlocks := make([]*ir.BasicBlock, 0)

	basicBlocks = append(basicBlocks, p.types.TranslateToLLVM(tables))
	basicBlocks = append(basicBlocks, p.declarations.TranslateToLLVM(tables, "global"))
	endBrace := &ir.BasicBlock{"", nil, nil, nil, ""}
	endBrace.Instructions = append(endBrace.Instructions, &ir.RightBrace{})

	for _, function := range p.functions.functions {
		basicBlocks = append(basicBlocks, function.TranslateToLLVM(tables))
		basicBlocks = append(basicBlocks, endBrace)
	}
	basicBlocks = append(basicBlocks, ir.GlobalDeclare)
	return basicBlocks
}
func (p *Program) DisplayLLVM(basicBlocks []*ir.BasicBlock, filename string) (string, []ir.Instruction) {
	var out bytes.Buffer
	out.WriteString("source_filename = \"" + filename + "\"\n")
	out.WriteString("target triple = \"x86_64-linux-gnu\"\n")
	var instructions []ir.Instruction
	for _, basicBlock := range basicBlocks {
		str, instr := basicBlock.DisplayLLVM()
		instructions = append(instructions, instr...)
		out.WriteString(str)

	}

	return out.String(), instructions
}

// func (p *Program) DisplayAST(level int) {
// 	var out bytes.Buffer
// 	out.WriteString(strings.Repeat("   ", level))
// 	out.WriteString("*ast.Program")
// 	// TODO Consider if import "fmt"; need to be printed
// 	fmt.Println(out.String())

// 	if p.types != nil {
// 		p.types.DisplayAST(level + 1)
// 	}
// 	if p.declarations != nil {
// 		p.declarations.DisplayAST(level + 1)
// 	}
// 	if p.functions != nil {
// 		// p.functions.DisplayAST(level + 1)
// 	}

// }
