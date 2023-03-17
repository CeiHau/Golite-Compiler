lexer grammar GoliteLexer; 

NUMBER: [1-9][0-9]* |[0-9];

PLUS: '+';

MINUS: '-';

ASSIGN: '=';

ASTERISK: '*';

FSLASH: '/';

PRINT: 'print';

LET: 'let';

TYPE: 'type';

STRUCT: 'struct';

INT_TYPE: 'int';
BOOL_TYPE: 'bool';
VAR: 'var';
FUNC: 'func';

IF: 'if';
ELSE: 'else';
FOR: 'for';
RETURN: 'return';
SCAN: 'scan';
DELETE: 'delete';
PRINTF: 'printf';
NEW: 'new';
TRUE:'true';
FALSE:'false';
NIL:'nil';

LEFTCURL: '{';
RIGHTCURL: '}';
LEFTBRAC: '(';
RIGHTBRAC: ')';

DOT: '.';
COMMA: ',';
EXCLAMATION: '!';

OR: '||';
AND: '&&';
EQUALS: '==';
NOTEQUALS: '!=';
MORETHAN: '>';
LESSTHAN: '<';
GEQ: '>=';
LEQ: '<=';

ID:[a-zA-Z][a-zA-Z0-9]*;

SEMICOLON: ';';

COMMENT: '//' ('\\' ["\\] | ~["\\\r\n])* -> skip ;

STRING: '"' .*? '"' ;

WS: [ \r\n\t]+ -> skip ;


