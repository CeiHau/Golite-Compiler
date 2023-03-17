// modified
package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type LValue struct {
	*token.Token
	id []string
}

func NewLValue(id []string, token *token.Token) *LValue {
	return &LValue{token, id}

}

func (lvs *LValue) String() string {
	var out bytes.Buffer
	for index, variable := range lvs.id {
		out.WriteString(variable)
		if index != len(lvs.id)-1 {
			out.WriteString(".")
		}
	}
	return out.String()
}

func (lv *LValue) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) string {
	curScope := tables.Funcs.Contains(scope)
	varEntry := curScope.Variables.Contains(lv.id[0])
	scopeSign := "@"
	if varEntry == nil {
		for _, para := range curScope.Parameters {
			if para.Name == lv.id[0] {
				varEntry = para
				scopeSign = "%"
				break
			}
		}
	}
	for _, v := range curScope.Variables.GetTable() {
		if v.Name == lv.id[0] {
			scopeSign = "%"
			break
		}
	}

	if len(lv.id) == 1 && (varEntry.Ty.String() == "int" || varEntry.Ty.String() == "bool") {
		return scopeSign + lv.id[0]
	} else {
		// Accessing struct field

		ind := 0
		structEntry := tables.Structs.Contains(varEntry.Ty.String())
		var result1, result2 *ir.Operand
		for index, _ := range lv.id {
			if index == 0 {
				continue
			}
			ty := tp.StringToType(structEntry.Name).LLVM()
			if structEntry != nil {
				ve, found := structEntry.Fields[lv.id[index]]
				if found {
					varEntry = &ve.Varentry
					ind = ve.Index
				}
			}

			result1 = ir.NewRegister(tables, block.FuncName)
			block.Instructions = append(block.Instructions, ir.NewLoad(result1.GetRegisterNum(), ty, scopeSign+lv.id[index-1]))
			result2 = ir.NewRegister(tables, block.FuncName)
			block.Instructions = append(block.Instructions, ir.NewStrAcc(result2.GetRegisterNum(), ty, result1.GetRegisterNum(), ind))
			structEntry = tables.Structs.Contains(varEntry.Ty.String())
		}
		if result2 == nil {
			return scopeSign + lv.id[0]
		}
		return result2.String()
	}
}
