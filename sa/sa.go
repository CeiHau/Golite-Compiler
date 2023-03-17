package sa

import (
	"golite/ast"
	st "golite/symboltable"
)

func reportErrors(errors []string) bool {
	if len(errors) == 0 {
		return false
	}
	for _, error := range errors {
		print(error + "\n")
	}
	return true
}

func Execute(program *ast.Program) *st.SymbolTables {

	//Define the compiler symbol tables
	tables := st.NewSymbolTables()

	//First Build the Symbol Table(s) for all global declarations
	program.BuildSymbolTable(tables)

	//Second perform type checking.
	errors := make([]string, 0)
	errors = program.TypeCheck(errors, tables)
	if !reportErrors(errors) {
		return tables
	}

	return nil
}
