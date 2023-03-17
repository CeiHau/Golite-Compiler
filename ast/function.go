package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type Function struct {
	*token.Token
	parameters   *Parameters
	variable     string
	returnType   tp.Type
	declarations *Declarations
	statements   *Statements
}

func NewFunction(parameters *Parameters, variable string, returnType tp.Type, declarations *Declarations, statements *Statements, token *token.Token) *Function {
	return &Function{token, parameters, variable, returnType, declarations, statements}
}

func (function *Function) BuildSymbolTable(tables *st.SymbolTables) {
	var parameters []*st.VarEntry
	// var variables *st.SymbolTable[*st.VarEntry]
	variables := st.NewSymbolTable[*st.VarEntry](tables.Globals)
	for _, decl := range function.parameters.decls {
		parameters = append(parameters, &st.VarEntry{decl.variable, decl.ty, st.LOCAL, false})
	}
	if function.declarations != nil {
		for _, declaration := range function.declarations.declarations {
			for _, variable := range declaration.ids.variables {
				variables.Insert(variable, &st.VarEntry{variable, declaration.ty, st.LOCAL, false})
			}
		}
	}
	tables.Funcs.Insert(function.variable, &st.FuncEntry{function.variable, function.returnType, parameters, variables, make([]string, 0)})

}
func (function *Function) String() string {
	var out bytes.Buffer
	out.WriteString("func ")
	out.WriteString(function.variable)
	if function.parameters != nil {
		out.WriteString(function.parameters.String())
	}
	out.WriteString(function.returnType.String())
	out.WriteString(" {\n")
	if function.declarations != nil {
		out.WriteString(function.declarations.String())
	}
	if function.statements != nil {
		out.WriteString(function.statements.String())
	}

	out.WriteString("}\n")

	return out.String()
}
func (function *Function) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	funcEntry := tables.Funcs.Contains(function.variable)
	errors = append(errors, funcEntry.Variables.Errors...)
	if function.declarations != nil {
		for _, decl := range function.declarations.declarations {
			errors = decl.TypeCheck(errors, tables, scope)
		}
	}
	if function.statements != nil {
		for _, state := range function.statements.statements {
			errors = state.TypeCheck(errors, tables, scope)
		}
	}
	if function.returnType == nil && function.variable != "main" {
		if !function.statements.statements[len(function.statements.statements)-1].CheckReturn(tables) {
			errors = append(errors, "missing return at end of function")
		}
	}
	return errors
}

func (function *Function) TranslateToLLVM(tables *st.SymbolTables) *ir.BasicBlock {
	/************  Function Blocks ************/
	funcDefBlock := ir.BasicBlock{"functionDef", nil, nil, nil, function.variable} // Function Definition
	funcBlock := ir.BasicBlock{ir.NewLabel(), nil, nil, nil, function.variable}    // Entry Block
	funcDefBlock.Succs = append(funcDefBlock.Succs, &funcBlock)

	// Get the funcEntry from symboltables
	funcEntry := tables.Funcs.Contains(function.variable)

	paramAlloca := []ir.Instruction{}
	pamNames := []string{}
	tys := []string{}

	for _, parameter := range funcEntry.Parameters {
		pamNames = append(pamNames, parameter.Name)
		tys = append(tys, parameter.Ty.LLVM())
		paramAlloca = append(paramAlloca, ir.NewAlloca("P_"+parameter.Name, parameter.Ty.LLVM()))
	}
	// Function Definition
	funcDefBlock.Instructions = append(funcDefBlock.Instructions, ir.NewFunctionDefinition(funcEntry.Name, tys, pamNames, funcEntry.RetTy.LLVM()))

	/************  Return value allocation ************/
	if funcEntry.RetTy.String() == "nil" {
		funcBlock.Instructions = append(funcBlock.Instructions, ir.NewAlloca("_retval", "i64"))
	} else {

		funcBlock.Instructions = append(funcBlock.Instructions, ir.NewAlloca("_retval", funcEntry.RetTy.LLVM()))
	}
	// fmt.Println("!!!!!")
	/************  Parameters allocation ************/
	funcBlock.Instructions = append(funcBlock.Instructions, paramAlloca...)

	/************ Local variable allocation ************/
	if function.declarations != nil {
		for _, decl := range function.declarations.declarations {
			funcBlock.Instructions = append(funcBlock.Instructions, decl.TranslateToLLVM(tables, "local")...)
		}
	}
	/************ Parameter store ************/
	for _, parameter := range funcEntry.Parameters {
		pamNames = append(pamNames, parameter.Name)
		funcBlock.Instructions = append(funcBlock.Instructions, ir.NewStore(parameter.Ty.LLVM(), "%"+parameter.Name, "%P_"+parameter.Name))
	}

	/************ Traverse all statements ************/
	curBlock := &funcBlock
	var exitBlock *ir.BasicBlock
	for index, stmt := range function.statements.statements {
		// curBlock.Instructions = append(curBlock.Instructions, ir.NewComment(stmt.String()))
		if index == len(function.statements.statements)-1 { // Check If the last statement is return statement

			if _, ok := stmt.(*Return); ok { // with Return statement
				if len(curBlock.Instructions) != 0 { // if the currentBlock is not empty we need new block
					exitBlock = &ir.BasicBlock{ir.NewLabel(), nil, nil, nil, function.variable}
					curBlock.Succs = append(curBlock.Succs, exitBlock)
					exitBlock.Preds = append(exitBlock.Preds, curBlock)
				} else { // exitBlock is current Block
					exitBlock = curBlock
				}
				stmt.TranslateToLLVM(tables, exitBlock, function.variable)
			} else { // without Return statement
				curBlock = stmt.TranslateToLLVM(tables, curBlock, function.variable)

				if len(curBlock.Instructions) != 0 { // if the currentBlock is not empty we need new block
					exitBlock = &ir.BasicBlock{ir.NewLabel(), nil, nil, nil, function.variable}
					curBlock.Succs = append(curBlock.Succs, exitBlock)
					exitBlock.Preds = append(exitBlock.Preds, curBlock)
				} else {
					exitBlock = curBlock
				}
				exitBlock.Instructions = append(exitBlock.Instructions, ir.NewStore("i64", "0", "%_retval"))
				rg2 := ir.NewRegister(tables, function.variable)
				exitBlock.Instructions = append(exitBlock.Instructions, ir.NewLoad(rg2.GetRegisterNum(), "i64", "%_retval"))
				exitBlock.Instructions = append(exitBlock.Instructions, ir.NewRet(rg2.GetRegisterNum(), "i64"))
			}
		} else {

			curBlock = stmt.TranslateToLLVM(tables, curBlock, function.variable)
		}
	}
	// fmt.Println("!!!!!")
	// If it's empty function
	if len(function.statements.statements) == 0 {
		exitBlock = &ir.BasicBlock{ir.NewLabel(), nil, nil, nil, function.variable}
		curBlock.Succs = append(curBlock.Succs, exitBlock)
		exitBlock.Preds = append(exitBlock.Preds, curBlock)
		exitBlock.Instructions = append(exitBlock.Instructions, ir.NewStore("i64", "0", "%_retval"))
		rg2 := ir.NewRegister(tables, function.variable)
		exitBlock.Instructions = append(exitBlock.Instructions, ir.NewLoad(rg2.GetRegisterNum(), "i64", "%_retval"))
		exitBlock.Instructions = append(exitBlock.Instructions, ir.NewRet(rg2.GetRegisterNum(), "i64"))
	}

	for _, prev := range exitBlock.Preds {
		lastInstr := prev.Instructions[len(prev.Instructions)-1]
		_, ifBr := lastInstr.(*ir.Br)
		_, ifRet := lastInstr.(*ir.Ret)
		if !ifBr && !ifRet { // if the last instruction is not br and ret
			prev.Instructions = append(prev.Instructions, ir.NewBr(nil, exitBlock.Label, "", ""))
		}

	}

	// /************ End of right  curly brace ************/
	// endBrace := &ir.BasicBlock{"", nil, nil, nil}
	// endBrace.Instructions = append(endBrace.Instructions, &ir.RightBrace{})
	// endBrace.Preds = append(endBrace.Preds, exitBlock)
	// exitBlock.Succs = append(exitBlock.Succs, endBrace)

	return &funcDefBlock
}
