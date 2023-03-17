// modified
package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Loop struct {
	*token.Token
	condition Expression
	loopBlock Block
}

func NewLoop(cond Expression, block Block, token *token.Token) *Loop {
	return &Loop{token, cond, block}
}

func (loo *Loop) String() string {
	var out bytes.Buffer

	out.WriteString("for (")
	out.WriteString(loo.condition.String())
	out.WriteString(")")
	out.WriteString((loo.loopBlock).String())
	return out.String()
}

func (loo *Loop) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	errors = loo.condition.TypeCheck(errors, symTable, scope)
	if loo.condition.GetType(symTable, scope) != "bool" {
		errors = append(errors, "Expression in For Loop does not evaluate to type bool")
	} else {
		errors = loo.loopBlock.TypeCheck(errors, symTable, scope)
	}
	return errors
}
func (loop *Loop) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {
	// Generate the CFG of for statement
	condBlock := &ir.BasicBlock{ir.NewLabel(), make([]ir.Instruction, 0), nil, nil, block.FuncName}
	bodyBlock := &ir.BasicBlock{ir.NewLabel(), make([]ir.Instruction, 0), nil, nil, block.FuncName}
	exitBlock := &ir.BasicBlock{ir.NewLabel(), make([]ir.Instruction, 0), nil, nil, block.FuncName}

	// Connect the blocks
	block.Succs = append(block.Succs, condBlock)

	condBlock.Preds = append(condBlock.Preds, block)
	condBlock.Succs = append(condBlock.Succs, bodyBlock)
	condBlock.Succs = append(condBlock.Succs, exitBlock)

	bodyBlock.Preds = append(bodyBlock.Preds, condBlock)

	exitBlock.Preds = append(exitBlock.Preds, condBlock)

	// Conditional Translation
	reg1 := loop.condition.TranslateToLLVM(tables, condBlock, scope)
	var condCode string
	if cmp, ok := condBlock.Instructions[len(condBlock.Instructions)-1].(*ir.Cmp); ok {
		condCode = cmp.GetCondCode()
	}
	condBlock.Instructions = append(condBlock.Instructions, ir.NewBr(reg1, bodyBlock.Label, exitBlock.Label, condCode))

	// Body Translation
	endBodyBlock := loop.loopBlock.TranslateToLLVM(tables, bodyBlock, scope)

	// Connect the end of body block back to conditional block
	endBodyBlock.Succs = append(endBodyBlock.Succs, condBlock)

	block.Instructions = append(block.Instructions, ir.NewBr(nil, condBlock.Label, "", ""))
	endBodyBlock.Instructions = append(endBodyBlock.Instructions, ir.NewBr(nil, condBlock.Label, "", ""))

	return exitBlock
}

func (loo *Loop) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}
