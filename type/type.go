package types

import "strings"

// Type includes information about working on Types
type Type interface {
	String() string
	LLVM() string
}

type StructType struct {
	Name   string
	Fields map[string]Type
	Order  []string
}

func (structType *StructType) String() string {
	return structType.Name
}

func (structType *StructType) LLVM() string {
	return "%struct." + structType.Name + "*"
}

type IntTy struct{}

func (intTy *IntTy) String() string {
	return "int"
}

func (intTy *IntTy) LLVM() string {
	return "i64"
}

type BoolTy struct{}

func (boolTy *BoolTy) String() string {
	return "bool"
}

func (boolTy *BoolTy) LLVM() string {
	return "i64"
}

func StringToType(ty string) Type {
	if ty == "int" {
		return &IntTy{} //TODO: MAKE THIS A SINGLETON!! Do not create a new type each time.
	} else if ty == "bool" {
		return &BoolTy{}
	} else if ty == "nil" {
		return &NilTy{}
	} else { // for temporary use
		if strings.HasPrefix(ty, "*") {
			ty = ty[1:]
		}
		return &StructType{ty, nil, nil}
	}
	panic("Not found type")
}

type NilTy struct{}

func (nilTy *NilTy) String() string {
	return "nil"
}

func (nilTy *NilTy) LLVM() string {
	return "i64"
}
