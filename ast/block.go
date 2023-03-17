package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Block struct {
	*token.Token
	statements Statements
}

func NewBlock(statements Statements, token *token.Token) *Block {
	return &Block{token, statements}
}

func (block *Block) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	out.WriteString(block.statements.String())
	out.WriteString("}")
	return out.String()
}

func (block *Block) CheckReturn(symTable *st.SymbolTables) bool {
	if len(block.statements.statements) == 0 {
		return false
	}
	return block.statements.statements[len(block.statements.statements)-1].CheckReturn(symTable)
}

func (block *Block) TypeCheck(errors []string, symTable *st.SymbolTables, scope string) []string {
	for _, block := range block.statements.statements {
		errors = block.TypeCheck(errors, symTable, scope)
	}
	return errors
}

func (b *Block) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {

	curBlock := block
	for _, stmt := range b.statements.statements {
		// curBlock.Instructions = append(curBlock.Instructions, ir.NewComment(stmt.String()))
		curBlock = stmt.TranslateToLLVM(tables, curBlock, scope)
	}
	return curBlock
}
