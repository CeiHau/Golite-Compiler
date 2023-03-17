package symboltable

import (
	"bytes"
	"fmt"
	tp "golite/type"
	"strings"
)

type VarScope int

const (
	GLOBAL VarScope = iota
	LOCAL
)

type VarEntry struct {
	Name   string
	Ty     tp.Type
	Scope  VarScope
	Usable bool
}

func (entry *VarEntry) String() string {
	if entry != nil {
		return entry.Name + " " + entry.Ty.String() + " " + ScopeToString(entry.Scope) + " VarEntry\n"
	}
	return ""
}

type FuncEntry struct {
	Name       string
	RetTy      tp.Type
	Parameters []*VarEntry
	Variables  *SymbolTable[*VarEntry]
	Register   []string
}

func (entry *FuncEntry) String() string {
	var out bytes.Buffer
	out.WriteString(entry.Name)
	out.WriteString("\nparameters:\n")
	for _, p := range entry.Parameters {
		out.WriteString(strings.Repeat(" ", 4) + p.String())
	}
	out.WriteString("return type: " + entry.RetTy.String() + "\n")
	out.WriteString("variables:\n")
	out.WriteString(strings.Repeat(" ", 2) + "LOCAL\n")
	for _, elem := range entry.Variables.table {
		out.WriteString(strings.Repeat(" ", 4) + elem.String())
	}
	out.WriteString(strings.Repeat(" ", 2) + "GLOBAL\n")
	for _, elem := range entry.Variables.parent.table {
		out.WriteString(strings.Repeat(" ", 4) + elem.String())
	}

	return out.String()
}

type Pair struct {
	Varentry VarEntry
	Index    int
}

type StructEntry struct {
	Name   string
	Fields map[string]*Pair
	Scope  VarScope
}

func (entry *StructEntry) String() string {
	var out bytes.Buffer
	out.WriteString(entry.Name + " " + "StrutEntry\n")
	for _, f := range entry.Fields {
		out.WriteString(strings.Repeat(" ", 4) + f.Varentry.String())
	}
	return out.String()

}

type SymbolTables struct {
	Funcs     *SymbolTable[*FuncEntry]
	Globals   *SymbolTable[*VarEntry]
	Structs   *SymbolTable[*StructEntry]
	Localtion map[string]map[string]string
}

func NewSymbolTables() *SymbolTables {
	return &SymbolTables{NewSymbolTable[*FuncEntry](nil), NewSymbolTable[*VarEntry](nil), NewSymbolTable[*StructEntry](nil), map[string]map[string]string{}}
}

type SymbolTable[T fmt.Stringer] struct {
	parent *SymbolTable[T]
	table  map[string]T
	Errors []string
}

func NewSymbolTable[T fmt.Stringer](parent *SymbolTable[T]) *SymbolTable[T] {
	return &SymbolTable[T]{parent, map[string]T{}, []string{}}
}
func (st *SymbolTable[T]) GetTable() map[string]T {
	return st.table
}
func (st *SymbolTable[T]) Insert(id string, entry T) {
	_, found := st.table[id]
	if found {
		st.Errors = append(st.Errors, id+" redeclared")
	}
	st.table[id] = entry
}

func (st *SymbolTable[T]) Contains(id string) T {
	value, found := st.table[id]
	if !found {
		if st.parent != nil {
			return st.parent.Contains(id)
		} else {
			return value
		}
	}
	return value
}

func ScopeToString(scope VarScope) string {
	switch scope {
	case GLOBAL:
		return "GLOBAL"
	case LOCAL:
		return "LOCAL"
	}
	panic("Error: Could not determine Scope")
}
func (st *SymbolTable[T]) DisplaySymbolTable() {
	for _, elem := range st.table {
		fmt.Print(elem)
	}
	if st.parent != nil {
		st.parent.DisplaySymbolTable()
	}
}

func (sts *SymbolTables) DisplaySymbolTables() {
	fmt.Println("Structs:")
	sts.Structs.DisplaySymbolTable()
	fmt.Println("\nGlobals:")
	sts.Globals.DisplaySymbolTable()
	fmt.Println("\nFunctions:")
	sts.Funcs.DisplaySymbolTable()
}
