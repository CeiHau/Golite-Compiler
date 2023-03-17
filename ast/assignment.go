// modified
package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
	"strings"
)

type Assignment struct {
	*token.Token
	lVal LValue
	expr Expression
}

func NewAssignment(LVal LValue, expr Expression, token *token.Token) *Assignment {

	return &Assignment{
		Token: token,
		lVal:  LVal,
		expr:  expr,
	}
}

func (ass *Assignment) String() string {

	var out bytes.Buffer
	out.WriteString(ass.lVal.String())
	out.WriteString(" = ")
	out.WriteString(ass.expr.String())
	out.WriteString(";\n")
	return out.String()
	// return (ass.lVal).String() + " = " + (ass.expr).String() + ";\n"
}

func (ast *Assignment) TypeCheck(errors []string, symTables *st.SymbolTables, scope string) []string {
	curScope := symTables.Funcs.Contains(scope)
	varEntry := curScope.Variables.Contains(ast.lVal.id[0])
	isParameter := false
	if varEntry == nil {
		for _, para := range curScope.Parameters {
			if para.Name == ast.lVal.id[0] {
				varEntry = para
				isParameter = true
				break
			}
		}
	}
	if varEntry == nil && !isParameter {
		errors = append(errors, "Variable "+ast.lVal.id[0]+" is not declared")
		return errors
	}
	varType := varEntry.Ty

	// Leftvalue is a value in some struct pointer
	if len(ast.lVal.id) > 1 {
		// Recursively check function calls
		counter := 1
		for counter < len(ast.lVal.id) {

			structEntry := symTables.Structs.Contains(varType.String())
			if structEntry == nil {
				errors = append(errors, "Struct "+varType.String()+" is not found ")
				return errors
			}
			varEntry, found := structEntry.Fields[ast.lVal.id[counter]]
			if !found {
				errors = append(errors, "Variable "+ast.lVal.id[counter]+" is not found within struct "+varType.String())
				return errors
			}
			varType = varEntry.Varentry.Ty
			counter++
		}
	}

	// Check right hand side
	errors = ast.expr.TypeCheck(errors, symTables, scope)
	right := ast.expr.GetType(symTables, scope)
	if varType.String() != right {
		if strings.Contains(right, "invalid operation") {
			errors = append(errors, strings.Split(right, ",")...)
		} else {
			if !(varType.String() != "bool" && varType.String() != "int" && right == "nil") {
				errors = append(errors, "invalid assignment: "+ast.String()+" (mismatched types "+varType.String()+" and "+right+")")
			}

		}
	}
	varEntry.Usable = true
	return errors
}

func (ast *Assignment) CheckReturn(symTable *st.SymbolTables) bool {
	return false
}

func (assignment *Assignment) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.BasicBlock {
	// Get the LValue type
	if assignment == nil {
		return block
	}
	originalScope := scope
	ty := tp.StringToType(assignment.expr.GetType(tables, scope)).LLVM()
	// Translate expression to LLVM
	if assignment.expr.String() == "nil" {
		curScope := tables.Funcs.Contains(scope)
		varEntry := curScope.Variables.Contains(assignment.lVal.id[0])
		if varEntry == nil {
			for _, para := range curScope.Parameters {
				if para.Name == assignment.lVal.id[0] {
					varEntry = para
					break
				}
			}
		}
		scope = varEntry.Ty.String()
		ty = tp.StringToType(varEntry.Ty.String()).LLVM()
	}
	value := assignment.expr.TranslateToLLVM(tables, block, scope)
	// Translate LValue to LLVM
	pointer := assignment.lVal.TranslateToLLVM(tables, block, originalScope)
	block.Instructions = append(block.Instructions, ir.NewStore(ty, value.String(), pointer))
	return block
}
