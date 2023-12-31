lexer grammar SwiftLexer;

// Skip comments

WHITESPACE: [ \t\n\r\f]+ -> skip;
COMMENT: '//' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

// Keywords used in declarations
Kw_LET: 'let';
Kw_VAR: 'var';
Kw_FUNC: 'func';
Kw_STRUCT: 'struct';
Kw_MUTATING: 'mutating';
Kw_AMPER: '&';

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
Kw_GUARD: 'guard';

// Types
Kw_INT: 'Int';
Kw_FLOAT: 'Float';
Kw_BOOL: 'Bool';
Kw_STRING: 'String';
Kw_CHAR: 'Character';

// Range expressions
Kw_RANGE: '...';
Kw_INOUT: 'inout';

// Literals
INT: [0-9]+;
FLOAT: [0-9]* '.' [0-9]+;
BOOL: 'true' | 'false';
// String recognize scape sequences
STRING: '"' ( '\\' [nrt"\\] | ~[\n\r"])* '"';
CHAR: '\'' ( '\\' [nrt"\\] | ~[\n\r']) '\'';
NIL: 'nil';
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