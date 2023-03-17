# Milestone 4
Successfully and fully design and implement arm code. 

The representation of arm assembly under the `ir` folder, **TranslateToAssembly** function.

To run and print out the Arm code, using the following command:
```
go run golite.go -s  ../benchmarks/simple/simple.golite
```
Our implementation can pass simple tests with simple assignment and invocation.
# Milestone 3
Successfully and fully design and implement LLVM IR. 

The representation of LLVM IR under the `ir` folder.

To run and print out the LLVM IR, using the following command:
```
go run golite.go -llvm  ../benchmarks/simple/simple.golite
```
# Milestone 2
Successfully and fully design and construct the abstract syntax tree. Partially completed the Semantic Analysis.

The representation of AST under the `ast` folder, the AST generate process in `parser.go` below comment  `/******************* Implementation of the Listeners **************************/`.

To run and print out the AST, using the following command:
```
go run golite.go -ast  ../benchmarks/simple/simple.golite
```
# Milestone1
## Compile and Run

We use Golang to build the compiler. To run the compiler on the CS Linux Servers:

Step 1: Clone this repo to your CS Linux Server

Step 2: cd golite/

Step 3: Run commands like below:

- For the parser, the compiler must read in a Golite program and parse through the input source code. E.g.
```
go run golite.go ../benchmarks/simple/simple.golite
```
- Add a flag -lex to command line arguments to print out each token on a separate line for a given program. E.g.
```
go run golite.go -lex  ../benchmarks/simple/simple.golite
```

## Testing

- Lexer results: Please see the result for **../benchmarks/simple/simple.golite**.

<img src="benchmarks/result/lexer.png" width="200"/>

- Parser results: Please see the result for **../benchmarks/bad/bad.golite**.

<img src="benchmarks/result/bad_parser.png" width="600"/>

