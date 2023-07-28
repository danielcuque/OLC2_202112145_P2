grammar Swift;

// Skip comments

WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

// Keywords used in declarations
Kw_LET: 'let';
Kw_VAR: 'var';
Kw_FUNC: 'func';
Kw_STRUCT: 'struct';

// Keywords used in control flow
Kw_IF: 'if';
Kw_ELSE: 'else';
Kw_SWITCH: 'switch';
Kw_CASE: 'case';
Kw_DEFAULT: 'default';
Kw_FOR: 'for';
Kw_WHILE: 'while';
Kw_BREAK: 'break';
Kw_CONTINUE: 'continue';
Kw_RETURN: 'return';
Kw_DO: 'do';
Kw_REPEAT: 'repeat';
Kw_IN: 'in';

// Types
Kw_INT: 'Int';
Kw_DOUBLE: 'Double';
Kw_BOOL: 'Bool';
Kw_STRING: 'String';
Kw_NIL: 'nil';

// Literals
INT: [0-9]+;
DOUBLE: [0-9]* '.' [0-9]+;
BOOL: 'true' | 'false';
STRING: '"' (~["\\\r\n] | '\\' [\\\r\n])* '"';
CHAR: '\'' (~['\\\r\n] | '\\' [\\\r\n])* '\'';

// Identifiers
ID: [a-zA-Z_][a-zA-Z0-9_]*;

// Operators
Op_ARROW: '->';

// Comparasion operators
Op_EQ: '==';
Op_NEQ: '!=';
Op_LT: '<';
Op_GT: '>';
Op_LE: '<=';
Op_GE: '>=';

// Assignment operators
Op_ASSIGN: '=';
Op_MUL_ASSIGN: '*=';
Op_DIV_ASSIGN: '/=';
Op_PLUS_ASSIGN: '+=';
Op_MINUS_ASSIGN: '-=';
Op_MOD_ASSIGN: '%=';

// Arithmetic operators
Op_MUL: '*';
Op_DIV: '/';
Op_PLUS: '+';
Op_MINUS: '-';
Op_MOD: '%';

// Logical operators
Op_AND: '&&';
Op_OR: '||';
Op_NOT: '!';
Op_TERNARY: '?';

// Delimiters
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';
BACKSLASH: '\\';

// Punctuation
COMMA: ',';
SEMICOLON: ';';
COLON: ':';
DOT: '.';

data_type: Kw_INT | Kw_DOUBLE | Kw_BOOL | Kw_STRING | Kw_NIL;

program: block EOF;

block: (statement)*;

statement: assignment | declaration | ifstmt | whilestmt;

assignment: ID Op_ASSIGN expr;

declaration: (Kw_VAR | Kw_LET) ID (COLON data_type)?
	| (Kw_VAR | Kw_LET) ID Op_ASSIGN expr;

ifstmt: Kw_IF LPAREN expr RPAREN LBRACE block RBRACE;

whilestmt: Kw_WHILE LPAREN expr RPAREN LBRACE block RBRACE;

expr:
	Op_MINUS expr													# UnaryExpr
	| left = expr op = (Op_MUL | Op_DIV) right = expr				# ArithmeticExpr
	| left = expr op = (Op_PLUS | Op_MINUS) right = expr			# ArithmeticExpr
	| left = expr op = Op_MOD right = expr							# ArithmeticExpr
	| left = expr op = (Op_GE | Op_GT) right = expr					# ComparasionExpr
	| left = expr op = (Op_LE | Op_LT) right = expr					# ComparasionExpr
	| left = expr op = (Op_EQ | Op_NEQ) right = expr				# ComparasionExpr
	| Op_NOT right = expr											# NotExpr
	| condition = expr Op_TERNARY cTrue = expr COLON cFalse = expr	# TernaryExpr
	| left = expr op = Op_AND right = expr							# LogicalExpr
	| left = expr op = Op_OR right = expr							# LogicalExpr
	| LPAREN expr RPAREN											# ParExpr
	| INT															# IntExpr
	| ID															# IdExpr
	| DOUBLE														# DoubleExpr
	| STRING														# StrExpr
	| BOOL															# BoolExpr;
