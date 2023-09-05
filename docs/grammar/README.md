# Gramática para el lenguaje T-Swift

A continuación se detalla la gramática para el lenguaje de programación T-Swift.

# No terminales

```	
program: block EOF

block: (statement)*

statement:
	variableAssignment SEMICOLON?
	| variableDeclaration SEMICOLON?
	| variableAssignmentObject SEMICOLON?
	| ifStatement
	| whileStatement
	| switchStatement
	| forStatement
	| guardStatement
	| controlFlowStatement
	| functionDeclarationStatement
	| functionCall SEMICOLON?
	| vectorDeclaration SEMICOLON?
	| vectorAssignment SEMICOLON?
	| matrixDeclaration SEMICOLON?
	| matrixAssignment SEMICOLON?
	| structDeclaration

variableType:
	Kw_INT
	| Kw_FLOAT
	| Kw_BOOL
	| Kw_STRING
	| Kw_CHAR
	| ID


variableDeclaration:
	varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_ASSIGN expr
	| varType = (Kw_LET | Kw_VAR) ID Op_ASSIGN expr 	
	| varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_TERNARY?  


functionDeclarationStatement:
	Kw_MUTATING? Kw_FUNC ID LPAREN functionParameters? RPAREN functionReturnType? LBRACE block
		RBRACE

functionParameters:
	functionParameter (COMMA functionParameter)*

functionParameter:
	ID? ID COLON Kw_INOUT? (
		variableType
		| matrixType
	)

functionParameterCompound:
	LBRACKET functionParameterCompound RBRACKET	
	| LBRACKET variableType RBRACKET		

functionReturnType: Op_ARROW variableType


functionCall:
	(idChain | variableType) LPAREN functionCallArguments? RPAREN

functionCallArguments:
	expr (COMMA expr)*						
	| ID COLON expr (COMMA ID COLON expr)*	

variableAssignment:
	idChain op = (Op_ASSIGN | Op_PLUS_ASSIGN | Op_MINUS_ASSIGN) expr

variableAssignmentObject:
	objectChain op = (
		Op_ASSIGN
		| Op_PLUS_ASSIGN
		| Op_MINUS_ASSIGN
	) expr

ifStatement: ifTail (Kw_ELSE ifTail)* elseStatement?

ifTail: Kw_IF expr LBRACE block RBRACE

elseStatement: Kw_ELSE LBRACE block RBRACE


whileStatement: Kw_WHILE expr LBRACE block RBRACE


switchStatement:
	Kw_SWITCH expr LBRACE switchCase* switchDefault? RBRACE

switchCase: Kw_CASE expr COLON block Kw_BREAK?

switchDefault: Kw_DEFAULT COLON block Kw_BREAK?

forStatement: Kw_FOR ID Kw_IN expr LBRACE block RBRACE

guardStatement: Kw_GUARD expr Kw_ELSE LBRACE block RBRACE

vectorDeclaration:
	varType = (Kw_LET | Kw_VAR) ID COLON LBRACKET variableType RBRACKET Op_ASSIGN vectorDefinition 
	| varType = (Kw_LET | Kw_VAR) ID Op_ASSIGN LBRACKET ID RBRACKET LPAREN RPAREN 

vectorDefinition:
	LBRACKET vectorValues? RBRACKET	
	| expr							

vectorValues: expr (COMMA expr)*

objectChain: vectorAccess (DOT ID)+

vectorAccess: idChain LBRACKET expr RBRACKET

vectorAssignment:
	vectorAccess op = (
		Op_ASSIGN
		| Op_PLUS_ASSIGN
		| Op_MINUS_ASSIGN
	) expr

matrixDeclaration:
	varType = (Kw_LET | Kw_VAR) idChain COLON matrixType Op_ASSIGN (
		matrixRepeatingDefinition
		| matrixDefinition
	)

matrixType:
	LBRACKET matrixType RBRACKET		# MatrixTypeNested
	| LBRACKET variableType RBRACKET	# MatrixTypeSingle

matrixDefinition: LBRACKET matrixValues? RBRACKET | expr

matrixValues: matrixDefinition (COMMA matrixDefinition)*

matrixRepeatingDefinition:
	matrixType LPAREN ID COLON matrixRepeatingDefinition COMMA ID COLON expr RPAREN #
	| matrixType LPAREN ID COLON expr COMMA ID COLON expr RPAREN 


matrixAccess:
	idChain LBRACKET expr RBRACKET (LBRACKET expr RBRACKET)*

matrixAssignment:
	matrixAccess op = (
		Op_ASSIGN
		| Op_PLUS_ASSIGN
		| Op_MINUS_ASSIGN
	) expr

structDeclaration: Kw_STRUCT ID LBRACE structBody RBRACE

structBody: structProperty*

structProperty:
	variableDeclaration
	| functionDeclarationStatement


expr:
	objectChain														
	| functionCall													
	| vectorAccess													
	| matrixAccess													
	| Op_MINUS expr													
	| Op_NOT right = expr											
	| left = expr op = (Op_MUL | Op_DIV) right = expr				
	| left = expr op = (Op_PLUS | Op_MINUS) right = expr			
	| left = expr op = Op_MOD right = expr							
	| left = expr op = (Op_GE | Op_GT) right = expr					
	| left = expr op = (Op_LE | Op_LT) right = expr					
	| left = expr op = (Op_EQ | Op_NEQ) right = expr				
	| condition = expr Op_TERNARY cTrue = expr COLON cFalse = expr	
	| left = expr op = Op_AND right = expr							
	| left = expr op = Op_OR right = expr							
	| left = expr Kw_RANGE right = expr								
	| LPAREN expr RPAREN											
	| INT															
	| Kw_AMPER? idChain												
	| FLOAT															
	| STRING														
	| NIL															
	| CHAR															
	| BOOL															


controlFlowStatement:
	Kw_BREAK			
	| Kw_CONTINUE		
	| Kw_RETURN expr?	

idChain: ID (DOT ID)* 

```

# Terminales

```
WHITESPACE: [ \t\n\r\f]+ 
COMMENT: '//' ~[\r\n]* 
BLOCK_COMMENT: '/*' .*? '*/' 

Kw_LET: 'let'
Kw_VAR: 'var'
Kw_FUNC: 'func'
Kw_STRUCT: 'struct'
Kw_MUTATING: 'mutating'
Kw_AMPER: '&'

Kw_IF: 'if'
Kw_ELSE: 'else'
Kw_SWITCH: 'switch'
Kw_CASE: 'case'
Kw_DEFAULT: 'default'
Kw_FOR: 'for'
Kw_WHILE: 'while'
Kw_BREAK: 'break'
Kw_CONTINUE: 'continue'
Kw_RETURN: 'return'
Kw_DO: 'do'
Kw_REPEAT: 'repeat'
Kw_IN: 'in'
Kw_GUARD: 'guard'

Kw_INT: 'Int'
Kw_FLOAT: 'Float'
Kw_BOOL: 'Bool'
Kw_STRING: 'String'
Kw_CHAR: 'Character'

Kw_RANGE: '...'
Kw_INOUT: 'inout'

INT: [0-9]+
FLOAT: [0-9]* '.' [0-9]+
BOOL: 'true' | 'false'
STRING: '"' ( '\\' [nrt"\\] | ~[\n\r"])* '"'
CHAR: '\'' ( '\\' [nrt"\\] | ~[\n\r']) '\''
NIL: 'nil'
ID: [a-zA-Z_][a-zA-Z0-9_]*

Op_ARROW: '->'

Op_EQ: '=='
Op_NEQ: '!='
Op_LT: '<'
Op_GT: '>'
Op_LE: '<='
Op_GE: '>='

Op_ASSIGN: '='
Op_MUL_ASSIGN: '*='
Op_DIV_ASSIGN: '/='
Op_PLUS_ASSIGN: '+='
Op_MINUS_ASSIGN: '-='
Op_MOD_ASSIGN: '%='

Op_MUL: '*'
Op_DIV: '/'
Op_PLUS: '+'
Op_MINUS: '-'
Op_MOD: '%'

Op_AND: '&&'
Op_OR: '||'
Op_NOT: '!'
Op_TERNARY: '?'

LPAREN: '('
RPAREN: ')'
LBRACE: '{'
RBRACE: '}'
LBRACKET: '['
RBRACKET: ']'
BACKSLASH: '\\'

COMMA: ','
SEMICOLON: ';'
COLON: ':'
DOT: '.'
```