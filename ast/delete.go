package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type Delete struct {
	*token.Token
	expression Expression
}

func NewDelete(expression Expression, token *token.Token) *Delete {
	return &Delete{token, expression}
}

func (delete *Delete) String() string {
	var out bytes.Buffer
	out.WriteString("delete")
	out.WriteString(delete.expression.String())
	out.WriteString(";")
	return out.String()
}

func (delete *Delete) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {

	argType := delete.expression.boolTerms[0].String()
	curScope := tables.Funcs.Contains(scope)
	varEntry := curScope.Variables.Contains(argType)
	isParameter := false
	if varEntry == nil {
		for _, para := range curScope.Parameters {
			if para.Name == argType {
				varEntry = para
				isParameter = true
				break
			}
		}
	}
	if varEntry == nil && !isParameter {
		errors = append(errors, "invalid function call to delete: (variable "+argType+" not found)")
	}
	return errors

}

// %r24 = load %struct.Point2D*, %struct.Point2D** %pt2
// %r25 = bitcast %struct.Point2D* %r24 to i8*
// call void @free(i8* %r25)
func (delete *Delete) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {
	rslt1 := ir.NewRegister(tables, block.FuncName)
	// Add malloc instruction to the block
	ty := tp.StringToType(delete.expression.GetType(tables, scope)).LLVM()
	rslt := delete.expression.TranslateToLLVM(tables, block, block.FuncName)
	pointer := rslt.String()
	print(pointer)
	block.Instructions = append(block.Instructions, ir.NewLoad(rslt1.GetRegisterNum(), ty, pointer))
	// Add bitcast instruction to the block
	rslt2 := ir.NewRegister(tables, block.FuncName)
	block.Instructions = append(block.Instructions, ir.NewBitcast(rslt2.GetRegisterNum(), ty, rslt1.GetRegisterNum(), "i8*"))
	block.Instructions = append(block.Instructions, ir.NewDelete(rslt2.GetRegisterNum()))
	return block
}

func (delete *Delete) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}
