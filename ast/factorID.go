package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type FactorID struct {
	*token.Token
	variable  string
	withArgs  bool
	arguments Arguments
	Usable    bool
}

func NewFactorID(variable string, withArgs bool, arguments Arguments, token *token.Token) *FactorID {
	return &FactorID{token, variable, withArgs, arguments, false}
}

func (fi *FactorID) String() string {
	var out bytes.Buffer
	out.WriteString(fi.variable)
	if fi.withArgs {
		out.WriteString(fi.arguments.String())
	}
	return out.String()
}

func (fi *FactorID) GetType(symTable *st.SymbolTables, scope string) string {
	if fi.withArgs {
		varEntry := symTable.Funcs.Contains(fi.variable)
		if varEntry == nil {
			return "nil"
		}
		return varEntry.RetTy.String()
	}
	curScope := symTable.Funcs.Contains(scope)
	varEntry := curScope.Variables.Contains(fi.variable)
	if varEntry == nil {
		for _, para := range curScope.Parameters {
			if para.Name == fi.variable {
				varEntry = para
				return varEntry.Ty.String()
			}
		}

	}
	if varEntry != nil {
		return varEntry.Ty.String()
	}
	return "nil"
}

func (fi *FactorID) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {
	// check function
	if fi.withArgs {
		funcEntry := tables.Funcs.Contains(fi.variable)
		if funcEntry == nil {
			errors = append(errors, "Function "+fi.variable+" is not declared")
		}
		if len(fi.arguments.expressions) > 0 {
			for _, arg := range fi.arguments.expressions {
				arg.TypeCheck(errors, tables, scope)
			}
		}
		return errors
	}

	curScope := tables.Funcs.Contains(scope)
	varEntry := curScope.Variables.Contains(fi.variable)
	// isParameter := false
	if varEntry == nil {
		for _, para := range curScope.Parameters {
			if para.Name == fi.variable {
				varEntry = para
				// isParameter = true
				break
			}
		}
	}

	if varEntry == nil {
		errors = append(errors, "Variable "+fi.variable+" is not declared")
	}
	// else {
	// 	if !varEntry.Usable && !isParameter && varEntry.Ty.String() != "int" && varEntry.Ty.String() != "bool" {
	// 		print(fi.Token.Line, fi.variable)
	// 		errors = append(errors, "Variable "+fi.variable+" is not initialized")
	// 	}
	// }

	return errors
}

func (fi *FactorID) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	if fi.withArgs { // It's a function

		var args []string
		for _, exp := range fi.arguments.expressions {

			ty := tp.StringToType(exp.GetType(tables, scope)).LLVM()
			rslt := exp.TranslateToLLVM(tables, block, block.FuncName)
			// rslt1 := ir.NewRegister(tables, block.FuncName)
			// block.Instructions = append(block.Instructions, ir.NewLoad(rslt1.GetRegisterNum(), ty, rslt.String()))
			args = append(args, ty+" "+rslt.String())
		}
		rslt := ir.NewRegister(tables, block.FuncName)

		block.Instructions = append(block.Instructions, ir.NewInvocation(rslt.GetRegisterNum(), tp.StringToType(fi.GetType(tables, scope)).LLVM(), fi.variable, args))
		// Return the Operands
		return rslt
	} else {
		ifPara := ""
		scopeSign := "@"
		funcEntry := tables.Funcs.Contains(scope)
		for _, p := range funcEntry.Parameters {
			if p.Name == fi.variable {
				ifPara = "P_"
				scopeSign = "%"
				break
			}
		}
		for _, v := range funcEntry.Variables.GetTable() {
			if v.Name == fi.variable {
				scopeSign = "%"
				break
			}
		}
		// new <result> register
		rslt := ir.NewRegister(tables, block.FuncName)
		// Add load instruction to the block
		ty := tp.StringToType(fi.GetType(tables, scope)).LLVM()
		block.Instructions = append(block.Instructions, ir.NewLoad(rslt.GetRegisterNum(), ty, scopeSign+ifPara+fi.variable))

		// Return the Operands
		return rslt

	}
}
