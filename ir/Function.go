package ir

import (
	"bytes"
	"golite/codegen"
	st "golite/symboltable"
)

type Function struct {
	functionName string
	retTy        string
	parameters   []Declaration
}

// define %struct.Point2D* @Init(i64 %initVal)
//
//	{
//	 ; Fill in with basic blocks from CFG
//	}
//
// define i64 @main()
//
//	{
//	 ; Fill in with basic blocks from CFG }
func NewFunction(functionName string, retTy string, parameters []Declaration) *Function {
	return &Function{functionName, retTy, parameters}
}

func (instr *Function) String() string {
	var out bytes.Buffer
	out.WriteString("define ")
	out.WriteString(instr.retTy)
	out.WriteString(" @")
	out.WriteString(instr.functionName)
	out.WriteString("(")
	for index, param := range instr.parameters {
		out.WriteString(param.ty)
		out.WriteString(" %")
		out.WriteString(param.variableName)
		if index != len(instr.parameters)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")

	// if function.statements != nil {
	// 	out.WriteString(function.statements.String())
	// }

	out.WriteString("}\n")
	return out.String()
}

func (instr *Function) TranslateToAssembly(tables st.SymbolTables, armTable map[string]string) []codegen.Instruction {
	return nil
}
