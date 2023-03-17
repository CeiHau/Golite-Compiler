package ast

import (
	"bytes"
	"golite/ir"
	st "golite/symboltable"
	"golite/token"
	tp "golite/type"
)

type SelectorTerm struct {
	*token.Token
	factor    Factor
	variables []string
}

func NewSelectorTerm(factor Factor, variables []string, token *token.Token) *SelectorTerm {
	return &SelectorTerm{token, factor, variables}

}

func (selectorTerm *SelectorTerm) String() string {
	var out bytes.Buffer
	out.WriteString(selectorTerm.factor.String())
	for _, variable := range selectorTerm.variables {
		out.WriteString("." + variable)
	}
	return out.String()
}

func (selectorTerm *SelectorTerm) GetType(symTable *st.SymbolTables, scope string) string {

	if len(selectorTerm.variables) == 0 {
		return selectorTerm.factor.GetType(symTable, scope)
	} else {
		selectSlice := selectorTerm.variables
		factorType := selectorTerm.factor.GetType(symTable, scope)
		structEntry := symTable.Structs.Contains(factorType)
		varEntry, _ := structEntry.Fields[selectSlice[0]]
		counter := 0
		for counter < len(selectorTerm.variables) {

			if structEntry == nil {
				return varEntry.Varentry.Ty.String()
			} else {
				varEntry, _ = structEntry.Fields[selectSlice[counter]]
			}
			structEntry = symTable.Structs.Contains(varEntry.Varentry.Ty.String())
			counter++
		}

		return varEntry.Varentry.Ty.String()
	}
}

func (selectorTerm *SelectorTerm) TypeCheck(errors []string, tables *st.SymbolTables, scope string) []string {

	if len(selectorTerm.variables) == 0 {
		return selectorTerm.factor.TypeCheck(errors, tables, scope)
	}
	selectSlice := selectorTerm.variables
	if len(selectSlice) >= 1 {
		factorType := selectorTerm.factor.GetType(tables, scope)
		structEntry := tables.Structs.Contains(factorType)
		// Recursively check function calls
		counter := 0
		for counter < len(selectSlice) {

			varEntry, found := structEntry.Fields[selectSlice[counter]]
			if !found {
				errors = append(errors, "Variable "+selectSlice[counter]+" is not found within struct "+varEntry.Varentry.Ty.String())
				return errors
			}
			counter++
			structEntry = tables.Structs.Contains(varEntry.Varentry.Ty.String())

		}
	}
	return errors
}

func (selectorTerm *SelectorTerm) TranslateToLLVM(tables *st.SymbolTables, block *ir.BasicBlock, scope string) *ir.Operand {
	var result1, result2 *ir.Operand
	if len(selectorTerm.variables) == 0 {
		return selectorTerm.factor.TranslateToLLVM(tables, block, scope)
	} else {
		// Accessing struct field
		ind := 0
		selectSlice := selectorTerm.variables
		factorType := selectorTerm.factor.GetType(tables, scope)
		structEntry := tables.Structs.Contains(factorType)
		varEntry, _ := structEntry.Fields[selectSlice[0]]
		// Recursively check function calls
		counter := 0
		pointer := selectorTerm.factor.String()
		for counter < len(selectSlice) {

			ty := tp.StringToType(structEntry.Name).LLVM()
			if counter != 0 {
				pointer = selectSlice[counter-1]
			}
			if structEntry != nil {
				ve, found := structEntry.Fields[selectSlice[counter]]
				if found {
					varEntry = ve
					ind = ve.Index
				}
			}
			result1 = ir.NewRegister(tables, block.FuncName)
			block.Instructions = append(block.Instructions, ir.NewLoad(result1.GetRegisterNum(), ty, "%"+pointer))
			result2 = ir.NewRegister(tables, block.FuncName)
			block.Instructions = append(block.Instructions, ir.NewStrAcc(result2.GetRegisterNum(), ty, result1.GetRegisterNum(), ind))
			varEntry, _ = structEntry.Fields[selectSlice[counter]]
			structEntry = tables.Structs.Contains(varEntry.Varentry.Ty.String())
			if structEntry == nil {
				return result2
			}

			counter++

		}

	}
	return result2
}
