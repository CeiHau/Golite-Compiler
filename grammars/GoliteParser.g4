parser grammar GoliteParser; 

options {
    tokenVocab = GoliteLexer;
}

/*
Program = Types Declarations Functions 'eof'                                                       ;
Types = {TypeDeclaration}                                                                          ;
TypeDeclaration = 'type' 'id' 'struct' '{' Fields '}' ';'                                          ;
Fields = Decl ';' {Decl ';'}                                                                       ;
Decl = 'id' Type                                                                                   ;
Type = 'int' | 'bool' | '*' 'id'                                                                   ;
Declarations = {Declaration}                                                                       ;
Declaration = 'var' Ids Type ';'                                                                   ;
Ids = 'id' {',' 'id'}                                                                              ;
Functions = {Function}                                                                             ;
Function = 'func' 'id' Parameters [ReturnType] '{' Declarations Statements '}'                     ;
Parameters = '(' [ Decl {',' Decl}] ')'                                                            ;
ReturnType = type                                                                                  ;
Statements = {Statement}                                                                           ;
Statement = Block | Assignment | Print | Delete | Conditional | Loop | Return | Read | Invocation  ;
Read = 'scan' LValue ';'                                                                       ;
Block = '{' Statements '}'                                                                         ;
Delete = 'delete' Expression ';'                                                                   ;
Assignment = LValue '=' (Expression) ';'                                                  ;
Print = 'printf' '(' 'string' { ',' Expression} ')'  ';'                                           ;
Conditional = 'if' '(' Expression ')' Block ['else' Block]                                         ;
Loop = 'for' '(' Expression ')' Block                                                              ;
Return = 'return' [Expression] ';'                                                                 ;
Invocation = 'id' Arguments ';'                                                                    ;
Arguments = '(' [Expression {',' Expression}] ')'                                                  ;
LValue = 'id' {'.' id}                                                                             ;
Expression = BoolTerm {'||' BoolTerm}                                                              ;
BoolTerm = EqualTerm {'&&' EqualTerm}                                                              ;
EqualTerm =  RelationTerm {('=='| '!=') RelationTerm}                                              ;
RelationTerm = SimpleTerm {('>'| '<' | '<=' | '>=') SimpleTerm}                                    ;
SimpleTerm = Term {('+'| '-') Term}                                                                ;
Term = UnaryTerm {('*'| '/') UnaryTerm}                                                            ;
UnaryTerm = '!' SelectorTerm | '-' SelectorTerm | SelectorTerm                                     ;
SelectorTerm = Factor {'.' 'id'}                                                                   ;
Factor = '(' Expression ')' | 'id' [Arguments] | 'number' | 'new' 'id' | 'true' | 'false' | 'nil'  ;
 */

program: types declarations functions EOF;
types: typeDeclaration* ;
typeDeclaration: TYPE ID STRUCT LEFTCURL fields RIGHTCURL SEMICOLON ;
fields: decl SEMICOLON (decl SEMICOLON)* ;
decl: ID ty=type ;
type: INT_TYPE | BOOL_TYPE | ASTERISK ID ;
declarations: declaration* ;
declaration: VAR ids ty=type SEMICOLON ;
ids: ID (COMMA ID)* ;
functions: function* ;
function: FUNC ID  parameters returnType? LEFTCURL declarations statements RIGHTCURL ;
parameters: LEFTBRAC (decl (COMMA decl)*)? RIGHTBRAC ;
returnType: ty=type ;
statements: statement* ;
statement: block | assignment | print | delete | conditional | loop | return | read | invocation ;
read: SCAN lValue SEMICOLON ;
block: LEFTCURL statements RIGHTCURL ;
delete: DELETE expression SEMICOLON ;
assignment: lValue ASSIGN expression SEMICOLON ;
print: PRINTF LEFTBRAC STRING (COMMA expression)* RIGHTBRAC SEMICOLON ;
conditional: IF LEFTBRAC expression RIGHTBRAC block (ELSE block) ? ;
loop: FOR LEFTBRAC expression RIGHTBRAC block ;
return: RETURN expression? SEMICOLON ;
invocation: ID arguments SEMICOLON ;
arguments: LEFTBRAC (expression (COMMA expression)*)? RIGHTBRAC ;
lValue: ID (DOT ID)* ;
expression: boolTerm (OR boolTerm)* ;
boolTerm: equalTerm (AND equalTerm)* ;

equalTerm: relationTerm opRelationTerm* ;
opRelationTerm: op=(EQUALS|NOTEQUALS) rt=relationTerm ;

relationTerm: simpleTerm opSimpleTerm* ;
opSimpleTerm: op=(MORETHAN|LESSTHAN|LEQ|GEQ) st=simpleTerm ;

simpleTerm: term opTerm* ;
opTerm: op=(PLUS|MINUS) t=term ;

term: unaryTerm opUnaryTerm* ;
opUnaryTerm: op=(ASTERISK|FSLASH) ut=unaryTerm ;

unaryTerm: (EXCLAMATION selectorTerm|MINUS selectorTerm|selectorTerm) ;
selectorTerm: factor (DOT ID)* ;
factor: (LEFTBRAC expression RIGHTBRAC | ID arguments? | NUMBER | NEW ID | TRUE | FALSE | NIL) ;





