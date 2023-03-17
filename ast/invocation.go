// modified
package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type Invocation struct {
	*token.Token
	id        string
	arguments Arguments
}

func NewInvocation(id string, arguments Arguments, token *token.Token) *Invocation {
	return &Invocation{token, id, arguments}
}

func (inv *Invocation) String() string {
	var out bytes.Buffer

	out.WriteString(inv.id)
	out.WriteString(inv.arguments.String())
	out.WriteString(");\n")
	return out.String()
}
func (inv *Invocation) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	funcEntry := tables.Funcs.Contains(inv.id)
	if funcEntry == nil {
		errors = append(errors, "Invalid undefined function call to "+inv.id)
		return errors
	}
	if len(funcEntry.Parameters) != len(inv.arguments.expressions) {
		errors = append(errors, "Invalid function call to "+inv.id+": unmatched number of arguments")
		return errors
	}

	// Checks Input Parameters
	for index, expr := range inv.arguments.expressions {
		exprType := expr.GetType(tables, scope)
		funcType := funcEntry.Parameters[index].Ty
		if exprType != funcType.String() {
			errors = append(errors, "invalid function call to "+inv.id+": (mismatched types"+exprType+" and "+funcType.String()+")")
		}
	}
	// Check return types
	// if funcEntry.RetTy.String() == "int" || funcEntry.RetTy.String() == "bool" {
	// 	errors = append(errors, inv.String()+" evaluated but not used")
	// 	return errors
	// }
	return errors
}

func (inv *Invocation) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}

func (inv *Invocation) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {

	// rslt := ir.NewRegister(tables, block.FuncName)
	var args []string
	for _, exp := range inv.arguments.expressions {
		ty := tp.StringToType(exp.GetType(tables, scope)).LLVM()
		rslt := exp.TranslateToLLVM(tables, block, block.FuncName)
		// rslt1 := ir.NewRegister(tables, block.FuncName)
		// block.Instructions = append(block.Instructions, ir.NewLoad(rslt1.GetRegisterNum(), ty, rslt.String()))
		args = append(args, ty+" "+rslt.String())
	}
	rslt := ir.NewRegister(tables, block.FuncName)
	funcEntry := tables.Funcs.Contains(inv.id)

	block.Instructions = append(block.Instructions, ir.NewInvocation(rslt.GetRegisterNum(), funcEntry.RetTy.LLVM(), inv.id, args))
	return block
}
