// modified
package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type Return struct {
	*token.Token
	withExpression bool
	expr           Expression
}

func NewReturn(expression Expression, withExpression bool, token *token.Token) *Return {
	return &Return{token, withExpression, expression}
}
func (ret *Return) String() string {
	var out bytes.Buffer

	out.WriteString("return ")
	out.WriteString(ret.expr.String()) //可能有错 因为是expression？
	out.WriteString(";\n")
	return out.String()
}
func (ret *Return) TypeCheck(errors []string, symTables *st.SymbolTables, scope string) []string {
	returTy := symTables.Funcs.Contains(scope).RetTy
	if ret.expr.boolTerms == nil && returTy == nil {
		errors = append(errors, "cannot return nothing in return argument")
	}
	return errors
}

func (ret *Return) BuildSymbolTable(tables *st.SymbolTables) {
}
func (ret *Return) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}

func (ret *Return) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {
	if ret.withExpression {
		rg1 := ret.expr.TranslateToLLVM(tables, block, scope)
		ty := tables.Funcs.Contains(scope).RetTy.LLVM()
		block.Instructions = append(block.Instructions, ir.NewStore(ty, rg1.String(), "%_retval"))
		rg2 := ir.NewRegister(tables, block.FuncName)
		block.Instructions = append(block.Instructions, ir.NewLoad(rg2.GetRegisterNum(), ty, "%_retval"))
		block.Instructions = append(block.Instructions, ir.NewRet(rg2.GetRegisterNum(), tp.StringToType(ret.expr.GetType(tables, scope)).LLVM()))
	} else {
		block.Instructions = append(block.Instructions, ir.NewStore("i64", "0", "%_retval"))
		rg2 := ir.NewRegister(tables, block.FuncName)
		block.Instructions = append(block.Instructions, ir.NewLoad(rg2.GetRegisterNum(), "i64", "%_retval"))
		block.Instructions = append(block.Instructions, ir.NewRet(rg2.GetRegisterNum(), "i64"))
	}
	return block

}
