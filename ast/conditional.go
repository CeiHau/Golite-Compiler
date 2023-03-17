// modified
package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Conditional struct {
	*token.Token
	condition Expression
	ifBlock   Block
	withElse  bool
	elseBlock Block
}

func NewConditional(cond Expression, ifBlock Block, withElse bool, elseBlock Block, token *token.Token) *Conditional {
	return &Conditional{token, cond, ifBlock, withElse, elseBlock}
}

func (con *Conditional) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString("(")
	out.WriteString((con.condition).String())
	out.WriteString(") ")
	out.WriteString((con.ifBlock).String())

	if con.withElse == true { //可能有错
		out.WriteString("else\n")
		out.WriteString((con.elseBlock).String())
	}
	return out.String()
}

func (con *Conditional) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}

func (con *Conditional) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	errors = con.condition.TypeCheck(errors, symTable, scope)
	if con.condition.GetType(symTable, scope) != "bool" {
		errors = append(errors, "If Expression does not evaluate to type bool")
	} else {
		errors = con.ifBlock.TypeCheck(errors, symTable, scope)
		if con.withElse {
			errors = con.elseBlock.TypeCheck(errors, symTable, scope)
		}
	}
	return errors
}

func (cond *Conditional) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {

	// Generate the CFG of condtional statement
	condBlock := &ir.BasicBlock{ir.NewLabel(), make([]ir.Instruction, 0), nil, nil, block.FuncName}
	ifBlock := &ir.BasicBlock{ir.NewLabel(), make([]ir.Instruction, 0), nil, nil, block.FuncName}
	elseBlock := &ir.BasicBlock{ir.NewLabel(), make([]ir.Instruction, 0), nil, nil, block.FuncName}
	endBlock := &ir.BasicBlock{ir.NewLabel(), make([]ir.Instruction, 0), nil, nil, block.FuncName}

	// Connectall the blocks
	block.Succs = append(block.Succs, condBlock)

	condBlock.Preds = append(condBlock.Preds, block)
	condBlock.Succs = append(condBlock.Succs, ifBlock)
	condBlock.Succs = append(condBlock.Succs, elseBlock)

	ifBlock.Preds = append(ifBlock.Preds, condBlock)
	// ifBlock.Succs = append(ifBlock.Succs, endBlock)

	elseBlock.Preds = append(elseBlock.Preds, condBlock)
	// elseBlock.Succs = append(elseBlock.Succs, endBlock)

	endBlock.Preds = append(endBlock.Preds, ifBlock)
	endBlock.Preds = append(endBlock.Preds, elseBlock)

	// Condition Translation
	reg1 := cond.condition.TranslateToLLVM(tables, condBlock, scope)
	var condCode string
	if cmp, ok := condBlock.Instructions[len(condBlock.Instructions)-1].(*ir.Cmp); ok {
		condCode = cmp.GetCondCode()
	}
	condBlock.Instructions = append(condBlock.Instructions, ir.NewBr(reg1, ifBlock.Label, elseBlock.Label, condCode))

	// Ifblock  and elseblock Translation
	endIfBlock := cond.ifBlock.TranslateToLLVM(tables, ifBlock, scope)
	endIfBlock.Succs = append(endIfBlock.Succs, endBlock)

	endElseBlock := cond.elseBlock.TranslateToLLVM(tables, elseBlock, scope)
	endElseBlock.Succs = append(endElseBlock.Succs, endBlock)

	// Add br instruction on previous block
	block.Instructions = append(block.Instructions, ir.NewBr(nil, condBlock.Label, "", ""))
	for _, preBlock := range endBlock.Preds {
		n := len(preBlock.Instructions)
		if n == 0 {
			preBlock.Instructions = append(preBlock.Instructions, ir.NewBr(nil, endBlock.Label, "", ""))
		} else {
			_, isRet := preBlock.Instructions[n-1].(*ir.Ret)
			_, isBr := preBlock.Instructions[n-1].(*ir.Br)
			if !isRet && !isBr { // If the last instruction is not ret instruction or br instruction
				preBlock.Instructions = append(preBlock.Instructions, ir.NewBr(nil, endBlock.Label, "", ""))
			}
		}

	}

	return endBlock
}
