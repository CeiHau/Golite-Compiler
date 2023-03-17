package main

import (
	"flag"
	"fmt"
	"golite/ir"
	"golite/lexer"
	"golite/parser"
	"golite/sa"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	lexFlag := flag.Bool("lex", false, "Print tokens")
	astFlag := flag.Bool("ast", false, "Print AST")
	llvmFlag := flag.Bool("llvm", false, "Print LLVM")
	armFlag := flag.Bool("s", false, "Print arm code")
	programFlag := flag.String("o", "", "program name input")

	flag.Parse()
	argLen := len(os.Args[1:])
	var arguments []string
	arguments = os.Args[:]
	inputSourcePath := arguments[argLen]
	file := filepath.Base(inputSourcePath)
	filename := strings.TrimSuffix(file, filepath.Ext(file))
	lexer := lexer.NewLexer(inputSourcePath)
	if *lexFlag {
		lexer.DisplayTokenStream()
		return
	}

	parser := parser.NewParser(lexer)
	ast := parser.Parse()

	// fmt.Println(ast.String())
	if *astFlag {
		parser.DisplayAST()
	}
	if tables := sa.Execute(ast); tables != nil {
		fmt.Println("passed")
		// tables.DisplaySymbolTables()
		fmt.Println("\n-----------------------------\n")
		basicBlocks := ast.TranslateToLLVM(tables)
		output, _ := ast.DisplayLLVM(basicBlocks, filename)
		if *llvmFlag {
			// fmt.Println(output)
			f, err := os.Create(filename + ".ll")
			if err != nil {
				panic(err)
			}
			f.WriteString(output)
			// return
		}
		if *armFlag {
			ir.TranslateToAssembly(basicBlocks, tables, filename)
			fmt.Println("\n-----------------------------\n")
			outName := "a.out"
			if *programFlag != "" {
				outName = *programFlag
			}
			cmd := exec.Command("clang", "-o", outName, filename+".s")
			cmd.Run()
			// return
		}

	} else {
		fmt.Println("failed")
	}
	fmt.Println("\n-----------------------------\n")

	fmt.Println("end")
}
