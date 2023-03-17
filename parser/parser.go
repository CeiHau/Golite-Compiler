package parser

import (
	"fmt"
	"golite/ast"
	"golite/context"
	"golite/lexer"
	"golite/token"
	tp "golite/type"
	"sort"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Parser interface {
	Parse() *ast.Program
	GetErrors() []*context.CompilerError
	DisplayAST()
}
type Nodes interface {
	String() string
}
type goliteParser struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	*BaseGoliteParserListener
	parser     *GoliteParser
	lexer      lexer.Lexer
	errors     []*context.CompilerError
	programKey string
	nodes      map[string]Nodes
}

func (gParser *goliteParser) GetErrors() []*context.CompilerError {
	return gParser.errors
}

func (gParser *goliteParser) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	gParser.errors = append(gParser.errors, &context.CompilerError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Phase:  context.PARSER,
	})
}

func NewParser(lexer lexer.Lexer) Parser {
	wrappedParser := &goliteParser{antlr.NewDefaultErrorListener(), &BaseGoliteParserListener{}, nil, lexer, nil, "", make(map[string]Nodes)}
	p := NewGoliteParser(lexer.GetTokenStream())
	p.RemoveErrorListeners()
	p.AddErrorListener(wrappedParser)
	wrappedParser.parser = p
	return wrappedParser
}

func (gParser *goliteParser) Parse() *ast.Program {
	ctx := gParser.parser.Program()

	if context.HasErrors(gParser.lexer.GetErrors()) ||
		context.HasErrors(gParser.GetErrors()) {
		return nil
	}
	antlr.ParseTreeWalkerDefault.Walk(gParser, ctx)
	_, _, key := GetTokenInfo(ctx)
	return gParser.nodes[key].(*ast.Program)
}

func keyToInfos(key string) (int, int) {
	infos := strings.Split(key, ",")
	r, _ := strconv.Atoi(infos[0])
	c, _ := strconv.Atoi(infos[1])
	return r, c
}
func (gParser *goliteParser) DisplayAST() {
	var keys []string
	for key, _ := range gParser.nodes {
		keys = append(keys, key)
	}

	sort.Slice(keys,
		func(i, j int) bool {
			ri, ci := keyToInfos(keys[i])
			rj, cj := keyToInfos(keys[j])
			if ri < rj {
				return true
			} else if ri == rj {
				return ci < cj
			} else {
				return false
			}
		})

	for i := len(keys) - 1; i >= 0; i-- {
		fmt.Println(keys[i], gParser.nodes[keys[i]])
	}

	// fmt.Println(gParser.nodes[gParser.programKey].(*ast.Program).String())
}

/******************* Implementation of the Listeners **************************/
// GetTokenInfo gerates a unique key for each node in the ParserTree
func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string) {
	key := fmt.Sprintf("%d,%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), key
}

// ExitProgram is called when production program is exited.
func (gParser *goliteParser) ExitProgram(ctx *ProgramContext) {
	// fmt.Println("Exit Program")
	line, col, key := GetTokenInfo(ctx)

	_, _, typesKey := GetTokenInfo(ctx.Types())

	types, _ := gParser.nodes[typesKey].(*ast.Types)

	_, _, declKey := GetTokenInfo(ctx.Declarations())
	decl, _ := gParser.nodes[declKey].(*ast.Declarations)

	_, _, funcsKey := GetTokenInfo(ctx.Functions())
	funcs, _ := gParser.nodes[funcsKey].(*ast.Functions)

	gParser.nodes[key] = ast.NewProgram(decl, types, funcs, token.NewToken(line, col))
	gParser.programKey = key
	// fmt.Println(gParser.nodes[key].(*ast.Program).String())
}

// ExitTypes is called when production types is exited.
func (gParser *goliteParser) ExitTypes(ctx *TypesContext) {
	// fmt.Println("Exit Types")
	line, col, key := GetTokenInfo(ctx)
	// fmt.Println(key)
	var typeDecls []ast.TypeDeclaration
	for _, typeDeclCtx := range ctx.AllTypeDeclaration() {
		_, _, key := GetTokenInfo(typeDeclCtx)
		astTypeDecl := gParser.nodes[key].(*ast.TypeDeclaration)
		typeDecls = append(typeDecls, *astTypeDecl)
	}
	if len(typeDecls) != 0 {
		gParser.nodes[key] = ast.NewTypes(typeDecls, token.NewToken(line, col))
	}

}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (gParser *goliteParser) ExitTypeDeclaration(ctx *TypeDeclarationContext) {
	line, col, key := GetTokenInfo(ctx)
	_, _, fieldsKey := GetTokenInfo(ctx.Fields())
	fields := gParser.nodes[fieldsKey].(*ast.Fields)
	gParser.nodes[key] = ast.NewTypeDeclaration(ctx.ID().GetText(), *fields, token.NewToken(line, col))
}

// ExitFields is called when production fields is exited.
func (gParser *goliteParser) ExitFields(ctx *FieldsContext) {
	line, col, key := GetTokenInfo(ctx)
	var decls []ast.Decl
	for _, declCtx := range ctx.AllDecl() {
		_, _, declkey := GetTokenInfo(declCtx)
		astDecl := gParser.nodes[declkey].(*ast.Decl)
		decls = append(decls, *astDecl)
	}
	gParser.nodes[key] = ast.NewFields(decls, token.NewToken(line, col))
}

// ExitDecl is called when production decl is exited.
func (gParser *goliteParser) ExitDecl(ctx *DeclContext) {
	line, col, key := GetTokenInfo(ctx)
	gParser.nodes[key] = ast.NewDecl(ctx.ID().GetText(), tp.StringToType(ctx.GetTy().GetText()), token.NewToken(line, col))

}

// ExitDeclarations is called when production declarations is exited.
func (gParser *goliteParser) ExitDeclarations(ctx *DeclarationsContext) {
	line, col, key := GetTokenInfo(ctx)
	var declarations []ast.Declaration
	for _, declCtx := range ctx.AllDeclaration() {
		_, _, declKey := GetTokenInfo(declCtx)
		astDecl := gParser.nodes[declKey].(*ast.Declaration)
		declarations = append(declarations, *astDecl)
	}
	if len(declarations) != 0 {
		gParser.nodes[key] = ast.NewDeclarations(declarations, token.NewToken(line, col))
	}

}

// ExitDeclaration is called when production declaration is exited.
func (gParser *goliteParser) ExitDeclaration(ctx *DeclarationContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, idsKey := GetTokenInfo(ctx.Ids())
	ids := gParser.nodes[idsKey].(*ast.Ids)
	// fmt.Println(ctx.GetTy().GetText())
	gParser.nodes[key] = ast.NewDeclaration(*ids, tp.StringToType(ctx.GetTy().GetText()), token.NewToken(line, col))
}

// ExitIds is called when production ids is exited.
func (gParser *goliteParser) ExitIds(ctx *IdsContext) {
	line, col, key := GetTokenInfo(ctx)
	var variables []string
	for _, variable := range ctx.AllID() {
		variables = append(variables, variable.GetText())
	}
	gParser.nodes[key] = ast.NewIds(variables, token.NewToken(line, col))
}

// ExitFunctions is called when production functions is exited.
func (gParser *goliteParser) ExitFunctions(ctx *FunctionsContext) {
	line, col, key := GetTokenInfo(ctx)
	var functions []ast.Function
	for _, funcCtx := range ctx.AllFunction() {
		_, _, funKey := GetTokenInfo(funcCtx)
		astFunc := gParser.nodes[funKey].(*ast.Function)
		functions = append(functions, *astFunc)
	}
	if len(functions) != 0 {
		gParser.nodes[key] = ast.NewFunctions(functions, token.NewToken(line, col))
	}
}

// ExitFunction is called when production function is exited.
func (gParser *goliteParser) ExitFunction(ctx *FunctionContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, parasKey := GetTokenInfo(ctx.Parameters())
	paras, _ := gParser.nodes[parasKey].(*ast.Parameters)

	_, _, declesKey := GetTokenInfo(ctx.Declarations())
	decles, _ := gParser.nodes[declesKey].(*ast.Declarations)

	_, _, stmtsKey := GetTokenInfo(ctx.Statements())
	stmts, _ := gParser.nodes[stmtsKey].(*ast.Statements)

	returnType := ctx.ReturnType()

	if returnType == nil {
		gParser.nodes[key] = ast.NewFunction(paras, ctx.ID().GetText(), tp.StringToType("nil"), decles, stmts, token.NewToken(line, col))
	} else {
		gParser.nodes[key] = ast.NewFunction(paras, ctx.ID().GetText(), tp.StringToType(returnType.GetTy().GetText()), decles, stmts, token.NewToken(line, col))
	}

}

// ExitParameters is called when production parameters is exited.
func (gParser *goliteParser) ExitParameters(ctx *ParametersContext) {
	line, col, key := GetTokenInfo(ctx)
	var decls []ast.Decl
	for _, declCtx := range ctx.AllDecl() {
		_, _, declKey := GetTokenInfo(declCtx)
		astDecl := gParser.nodes[declKey].(*ast.Decl)
		decls = append(decls, *astDecl)
	}
	gParser.nodes[key] = ast.NewParameters(decls, token.NewToken(line, col))
}

// ExitStatements is called when production statements is exited.
func (gParser *goliteParser) ExitStatements(ctx *StatementsContext) {
	line, col, key := GetTokenInfo(ctx)
	var statements []ast.Statement
	for _, stmtCtx := range ctx.AllStatement() {
		_, _, stmtKey := GetTokenInfo(stmtCtx)
		astStmt := gParser.nodes[stmtKey].(ast.Statement)
		statements = append(statements, astStmt)
	}
	gParser.nodes[key] = ast.NewStatements(statements, token.NewToken(line, col))

}

// ExitRead is called when production read is exited.
func (gParser *goliteParser) ExitRead(ctx *ReadContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, lvKey := GetTokenInfo(ctx.LValue())
	lv := gParser.nodes[lvKey].(*ast.LValue)

	gParser.nodes[key] = ast.NewRead(*lv, token.NewToken(line, col))
}

// ExitBlock is called when production block is exited.
func (gParser *goliteParser) ExitBlock(ctx *BlockContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, stmtsKey := GetTokenInfo(ctx.Statements())
	stmts := gParser.nodes[stmtsKey].(*ast.Statements)

	gParser.nodes[key] = ast.NewBlock(*stmts, token.NewToken(line, col))
}

// ExitDelete is called when production delete is exited.
func (gParser *goliteParser) ExitDelete(ctx *DeleteContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, expKey := GetTokenInfo((ctx.Expression()))
	exp := gParser.nodes[expKey].(*ast.Expression)

	gParser.nodes[key] = ast.NewDelete(*exp, token.NewToken(line, col))
}

// ExitAssignment is called when production assignment is exited.
func (gParser *goliteParser) ExitAssignment(ctx *AssignmentContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, lvKey := GetTokenInfo(ctx.LValue())
	lv := gParser.nodes[lvKey].(*ast.LValue)

	_, _, expKey := GetTokenInfo((ctx.Expression()))
	exp := gParser.nodes[expKey].(*ast.Expression)

	gParser.nodes[key] = ast.NewAssignment(*lv, *exp, token.NewToken(line, col))
}

// ExitPrint is called when production print is exited.
func (gParser *goliteParser) ExitPrint(ctx *PrintContext) {
	line, col, key := GetTokenInfo(ctx)

	var exprs []ast.Expression

	for _, exprCtx := range ctx.AllExpression() {
		_, _, exprKey := GetTokenInfo(exprCtx)
		astExpr := gParser.nodes[exprKey].(*ast.Expression)
		exprs = append(exprs, *astExpr)
	}
	gParser.nodes[key] = ast.NewPrint(ctx.STRING().GetText(), exprs, token.NewToken(line, col))

}

// ExitConditional is called when production conditional is exited.
func (gParser *goliteParser) ExitConditional(ctx *ConditionalContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, expKey := GetTokenInfo((ctx.Expression()))
	exp := gParser.nodes[expKey].(*ast.Expression)

	blocks := ctx.AllBlock()

	var ifB, elseB ast.Block
	withElse := false
	_, _, ifBKey := GetTokenInfo(blocks[0])
	ifB = *gParser.nodes[ifBKey].(*ast.Block)

	if len(blocks) == 2 {
		_, _, elseBKey := GetTokenInfo(blocks[1])
		elseB = *gParser.nodes[elseBKey].(*ast.Block)
		withElse = true
	}

	gParser.nodes[key] = ast.NewConditional(*exp, ifB, withElse, elseB, token.NewToken(line, col))

}

// ExitLoop is called when production loop is exited.
func (gParser *goliteParser) ExitLoop(ctx *LoopContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, expKey := GetTokenInfo((ctx.Expression()))
	exp := gParser.nodes[expKey].(*ast.Expression)

	_, _, blockKey := GetTokenInfo(ctx.Block())
	block := gParser.nodes[blockKey].(*ast.Block)

	gParser.nodes[key] = ast.NewLoop(*exp, *block, token.NewToken(line, col))

}

// ExitReturn is called when production return is exited.
func (gParser *goliteParser) ExitReturn(ctx *ReturnContext) {
	line, col, key := GetTokenInfo(ctx)
	var exp ast.Expression

	if expCtx := ctx.Expression(); expCtx != nil {
		_, _, expKey := GetTokenInfo(expCtx)
		exp = *gParser.nodes[expKey].(*ast.Expression)
		gParser.nodes[key] = ast.NewReturn(exp, true, token.NewToken(line, col))
	} else {
		gParser.nodes[key] = ast.NewReturn(exp, false, token.NewToken(line, col))
	}
}

// ExitInvocation is called when production invocation is exited.
func (gParser *goliteParser) ExitInvocation(ctx *InvocationContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, argsKey := GetTokenInfo(ctx.Arguments())
	args := gParser.nodes[argsKey].(*ast.Arguments)

	gParser.nodes[key] = ast.NewInvocation(ctx.ID().GetText(), *args, token.NewToken(line, col))
}

// ExitArguments is called when production arguments is exited.
func (gParser *goliteParser) ExitArguments(ctx *ArgumentsContext) {
	line, col, key := GetTokenInfo(ctx)

	var exprs []ast.Expression

	for _, exprCtx := range ctx.AllExpression() {
		_, _, exprKey := GetTokenInfo(exprCtx)
		astExpr := gParser.nodes[exprKey].(*ast.Expression)
		exprs = append(exprs, *astExpr)
	}
	gParser.nodes[key] = ast.NewArguments(exprs, token.NewToken(line, col))

}

// ExitLValue is called when production lValue is exited.
func (gParser *goliteParser) ExitLValue(ctx *LValueContext) {
	line, col, key := GetTokenInfo(ctx)
	var variables []string
	for _, variable := range ctx.AllID() {
		variables = append(variables, variable.GetText())
	}
	gParser.nodes[key] = ast.NewLValue(variables, token.NewToken(line, col))
}

// ExitExpression is called when production expression is exited.
func (gParser *goliteParser) ExitExpression(ctx *ExpressionContext) {
	line, col, key := GetTokenInfo(ctx)

	var boolTerms []ast.BoolTerm

	for _, boolTerm := range ctx.AllBoolTerm() {
		_, _, btKey := GetTokenInfo(boolTerm)
		astBoolTerm := gParser.nodes[btKey].(*ast.BoolTerm)
		boolTerms = append(boolTerms, *astBoolTerm)
	}
	gParser.nodes[key] = ast.NewExpression(boolTerms, token.NewToken(line, col))
}

// ExitBoolTerm is called when production boolTerm is exited.
func (gParser *goliteParser) ExitBoolTerm(ctx *BoolTermContext) {
	line, col, key := GetTokenInfo(ctx)

	var equalTerms []ast.EqualTerm

	for _, equalTerm := range ctx.AllEqualTerm() {
		_, _, eTKey := GetTokenInfo(equalTerm)
		astEqualTerm := gParser.nodes[eTKey].(*ast.EqualTerm)
		equalTerms = append(equalTerms, *astEqualTerm)
	}
	gParser.nodes[key] = ast.NewBoolTerm(equalTerms, token.NewToken(line, col))
}

// ExitEqualTerm is called when production equalTerm is exited.
func (gParser *goliteParser) ExitEqualTerm(ctx *EqualTermContext) {
	line, col, key := GetTokenInfo(ctx)

	var relationTerms []ast.RelationTerm
	var operators []string

	_, _, rlTmKey := GetTokenInfo(ctx.RelationTerm())
	rlTm := gParser.nodes[rlTmKey].(*ast.RelationTerm)
	relationTerms = append(relationTerms, *rlTm)
	for _, opRlTm := range ctx.AllOpRelationTerm() {
		op := opRlTm.GetOp().GetText()
		_, _, rlTmKey := GetTokenInfo(opRlTm.GetRt())
		rlTm := gParser.nodes[rlTmKey].(*ast.RelationTerm)

		relationTerms = append(relationTerms, *rlTm)
		operators = append(operators, op)
	}
	gParser.nodes[key] = ast.NewEqualTerm(operators, relationTerms, token.NewToken(line, col))

}

// ExitRelationTerm is called when production relationTerm is exited.
func (gParser *goliteParser) ExitRelationTerm(ctx *RelationTermContext) {
	line, col, key := GetTokenInfo(ctx)

	var simpleTerms []ast.SimpleTerm
	var operators []string

	_, _, smplTmKey := GetTokenInfo(ctx.SimpleTerm())
	smplTm := gParser.nodes[smplTmKey].(*ast.SimpleTerm)
	simpleTerms = append(simpleTerms, *smplTm)

	for _, opSmplTm := range ctx.AllOpSimpleTerm() {
		op := opSmplTm.GetOp().GetText()
		_, _, smplTmKey := GetTokenInfo(opSmplTm.GetSt())
		smplTm := gParser.nodes[smplTmKey].(*ast.SimpleTerm)

		simpleTerms = append(simpleTerms, *smplTm)
		operators = append(operators, op)
	}
	gParser.nodes[key] = ast.NewRelationTerm(operators, simpleTerms, token.NewToken(line, col))
}

// ExitSimpleTerm is called when production simpleTerm is exited.
func (gParser *goliteParser) ExitSimpleTerm(ctx *SimpleTermContext) {
	line, col, key := GetTokenInfo(ctx)

	var terms []ast.Term
	var operators []string

	_, _, tmKey := GetTokenInfo(ctx.Term())
	tm := gParser.nodes[tmKey].(*ast.Term)
	terms = append(terms, *tm)

	for _, opTm := range ctx.AllOpTerm() {
		op := opTm.GetOp().GetText()
		_, _, tmKey := GetTokenInfo(opTm.GetT())
		tm := gParser.nodes[tmKey].(*ast.Term)

		terms = append(terms, *tm)
		operators = append(operators, op)
	}
	gParser.nodes[key] = ast.NewSimpleTerm(terms, operators, token.NewToken(line, col))
}

// ExitTerm is called when production term is exited.
func (gParser *goliteParser) ExitTerm(ctx *TermContext) {
	line, col, key := GetTokenInfo(ctx)

	var unaryTerms []ast.UnaryTerm
	var operators []string

	_, _, uTmKey := GetTokenInfo(ctx.UnaryTerm())
	uTm := gParser.nodes[uTmKey].(*ast.UnaryTerm)
	unaryTerms = append(unaryTerms, *uTm)

	for _, opUTm := range ctx.AllOpUnaryTerm() {
		op := opUTm.GetOp().GetText()
		_, _, uTmKey := GetTokenInfo(opUTm.GetUt())
		uTm := gParser.nodes[uTmKey].(*ast.UnaryTerm)

		unaryTerms = append(unaryTerms, *uTm)
		operators = append(operators, op)
	}
	gParser.nodes[key] = ast.NewTerm(unaryTerms, operators, token.NewToken(line, col))
}

// ExitUnaryTerm is called when production unaryTerm is exited.
func (gParser *goliteParser) ExitUnaryTerm(ctx *UnaryTermContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, slctTmKey := GetTokenInfo(ctx.SelectorTerm())
	slctTm := gParser.nodes[slctTmKey].(*ast.SelectorTerm)
	var operator string
	if ctx.EXCLAMATION() != nil {
		operator = ctx.EXCLAMATION().GetText()
	} else if ctx.MINUS() != nil {
		operator = ctx.MINUS().GetText()
	} else {
		operator = ""
	}

	gParser.nodes[key] = ast.NewUnaryTerm(*slctTm, operator, token.NewToken(line, col))
}

// ExitSelectorTerm is called when production selectorTerm is exited.
func (gParser *goliteParser) ExitSelectorTerm(ctx *SelectorTermContext) {
	line, col, key := GetTokenInfo(ctx)

	_, _, factorKey := GetTokenInfo(ctx.Factor())
	factor := gParser.nodes[factorKey].(ast.Factor)

	var variables []string

	for _, id := range ctx.AllID() {
		variables = append(variables, id.GetText())
	}
	gParser.nodes[key] = ast.NewSelectorTerm(factor, variables, token.NewToken(line, col))
}

// ExitFactor is called when production factor is exited.
func (gParser *goliteParser) ExitFactor(ctx *FactorContext) {
	line, col, key := GetTokenInfo(ctx)

	if ctx.Expression() != nil { //LEFTBRAC expression RIGHTBRAC
		_, _, expKey := GetTokenInfo(ctx.Expression())
		exp := gParser.nodes[expKey].(*ast.Expression)
		gParser.nodes[key] = ast.NewFactorExpression(*exp, token.NewToken(line, col))
		return
	} else if ctx.NEW() != nil { //NEW ID
		gParser.nodes[key] = ast.NewNewID(ctx.ID().GetText(), token.NewToken(line, col))
		return
	} else if ctx.ID() != nil { //ID arguments?
		if ctx.Arguments() != nil {

			_, _, argsKey := GetTokenInfo(ctx.Arguments())
			args := gParser.nodes[argsKey].(*ast.Arguments)
			gParser.nodes[key] = ast.NewFactorID(ctx.ID().GetText(), true, *args, token.NewToken(line, col))
		} else {
			var args ast.Arguments
			gParser.nodes[key] = ast.NewFactorID(ctx.ID().GetText(), false, args, token.NewToken(line, col))
		}
	} else if ctx.NUMBER() != nil { // NUMBER
		intValue, _ := strconv.Atoi(ctx.NUMBER().GetText())
		gParser.nodes[key] = ast.NewIntLit(int64(intValue), token.NewToken(line, col))
		return
	} else if ctx.TRUE() != nil { // TRUE
		gParser.nodes[key] = ast.NewBoolLit(true, token.NewToken(line, col))
		return
	} else if ctx.FALSE() != nil { // FALSE
		gParser.nodes[key] = ast.NewBoolLit(false, token.NewToken(line, col))
		return
	} else if ctx.NIL() != nil {
		gParser.nodes[key] = ast.NewNilLit(token.NewToken(line, col))
		return
	}

}
