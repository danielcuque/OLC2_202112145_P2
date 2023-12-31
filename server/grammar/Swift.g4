parser grammar Swift;

options {
	tokenVocab = SwiftLexer;
}

idChain: ID (DOT ID)* # IDChain;

program: block EOF;

block: (statement)*;

// Statements
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
	| structDeclaration;

// Variable types
variableType:
	Kw_INT
	| Kw_FLOAT
	| Kw_BOOL
	| Kw_STRING
	| Kw_CHAR
	| ID;

// Variable declaration

variableDeclaration:
	varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_ASSIGN expr 	# TypeValueDeclaration
	| varType = (Kw_LET | Kw_VAR) ID Op_ASSIGN expr 	# ValueDeclaration
	| varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_TERNARY?  # TypeDeclaration;

// Function declaration
functionDeclarationStatement:
	Kw_MUTATING? Kw_FUNC ID LPAREN functionParameters? RPAREN functionReturnType? LBRACE block
		RBRACE;

functionParameters:
	functionParameter (COMMA functionParameter)*;

// Change to accept vectors 
functionParameter:
	ID? ID COLON Kw_INOUT? (
		variableType
		| matrixType
	);

functionParameterCompound:
	LBRACKET functionParameterCompound RBRACKET	# FunctionParameterCompoundNested
	| LBRACKET variableType RBRACKET		# FunctionParameterCompoundSingle;

functionReturnType: Op_ARROW variableType;

// Function call 
functionCall:
	(idChain | variableType) LPAREN functionCallArguments? RPAREN;

functionCallArguments:
	expr (COMMA expr)*						# Arguments
	| ID COLON expr (COMMA ID COLON expr)*	# NamedArguments;

// Variable assignment
variableAssignment:
	idChain op = (Op_ASSIGN | Op_PLUS_ASSIGN | Op_MINUS_ASSIGN) expr;

variableAssignmentObject:
	objectChain op = (
		Op_ASSIGN
		| Op_PLUS_ASSIGN
		| Op_MINUS_ASSIGN
	) expr;

// If statement 
ifStatement: ifTail (Kw_ELSE ifTail)* elseStatement?;

ifTail: Kw_IF expr LBRACE block RBRACE;

elseStatement: Kw_ELSE LBRACE block RBRACE;

// While statement
whileStatement: Kw_WHILE expr LBRACE block RBRACE;

// Switch statement
switchStatement:
	Kw_SWITCH expr LBRACE switchCase* switchDefault? RBRACE;

switchCase: Kw_CASE expr COLON block Kw_BREAK?;

switchDefault: Kw_DEFAULT COLON block Kw_BREAK?;

// For statement, we will have two types of for statements, one for arrays and one for ranges
forStatement: Kw_FOR ID Kw_IN expr LBRACE block RBRACE;

// Guard statement
guardStatement: Kw_GUARD expr Kw_ELSE LBRACE block RBRACE;

// Vector declarations

vectorDeclaration:
	varType = (Kw_LET | Kw_VAR) ID COLON LBRACKET variableType RBRACKET Op_ASSIGN vectorDefinition # VectorTypeValue
	| varType = (Kw_LET | Kw_VAR) ID Op_ASSIGN LBRACKET ID RBRACKET LPAREN RPAREN # VectorStructValue;

vectorDefinition:
	LBRACKET vectorValues? RBRACKET	# VectorListValue
	| expr							# VectorSingleValue;

vectorValues: expr (COMMA expr)*;

objectChain: vectorAccess (DOT ID)+;

vectorAccess: idChain LBRACKET expr RBRACKET;

vectorAssignment:
	vectorAccess op = (
		Op_ASSIGN
		| Op_PLUS_ASSIGN
		| Op_MINUS_ASSIGN
	) expr;

// Matrix declarations

matrixDeclaration:
	varType = (Kw_LET | Kw_VAR) idChain COLON matrixType Op_ASSIGN (
		matrixRepeatingDefinition
		| matrixDefinition
	);

matrixType:
	LBRACKET matrixType RBRACKET		# MatrixTypeNested
	| LBRACKET variableType RBRACKET	# MatrixTypeSingle;

matrixDefinition: LBRACKET matrixValues? RBRACKET | expr;

matrixValues: matrixDefinition (COMMA matrixDefinition)*;

matrixRepeatingDefinition:
	matrixType LPAREN ID COLON matrixRepeatingDefinition COMMA ID COLON expr RPAREN #
		MatrixRepeatingDefinitionNested
	| matrixType LPAREN ID COLON expr COMMA ID COLON expr RPAREN # MatrixRepeatingDefinitionSingle;

// Matrix access can be matrix1[0][1]...[n]
matrixAccess:
	idChain LBRACKET expr RBRACKET (LBRACKET expr RBRACKET)*;

matrixAssignment:
	matrixAccess op = (
		Op_ASSIGN
		| Op_PLUS_ASSIGN
		| Op_MINUS_ASSIGN
	) expr;

// Structs

structDeclaration: Kw_STRUCT ID LBRACE structBody RBRACE;

structBody: structProperty*;

structProperty:
	variableDeclaration
	| functionDeclarationStatement;

// Expressions
expr:
	objectChain														# ObjectChainExpr
	| functionCall													# FunctionCallExpr
	| vectorAccess													# VectorAccessExpr
	| matrixAccess													# MatrixAccessExpr
	| Op_MINUS expr													# UnaryExpr
	| Op_NOT right = expr											# NotExpr
	| left = expr op = (Op_MUL | Op_DIV) right = expr				# ArithmeticExpr
	| left = expr op = (Op_PLUS | Op_MINUS) right = expr			# ArithmeticExpr
	| left = expr op = Op_MOD right = expr							# ArithmeticExpr
	| left = expr op = (Op_GE | Op_GT) right = expr					# ComparisonExpr
	| left = expr op = (Op_LE | Op_LT) right = expr					# ComparisonExpr
	| left = expr op = (Op_EQ | Op_NEQ) right = expr				# ComparisonExpr
	| condition = expr Op_TERNARY cTrue = expr COLON cFalse = expr	# TernaryExpr
	| left = expr op = Op_AND right = expr							# LogicalExpr
	| left = expr op = Op_OR right = expr							# LogicalExpr
	| left = expr Kw_RANGE right = expr								# RangeExpr
	| LPAREN expr RPAREN											# ParExpr
	| INT															# IntExpr
	| Kw_AMPER? idChain												# IdExpr
	| FLOAT															# FloatExpr
	| STRING														# StrExpr
	| NIL															# NilExpr
	| CHAR															# CharExpr
	| BOOL															# BoolExpr;

// ControlFlow expressions
controlFlowStatement:
	Kw_BREAK			# ControlBreak
	| Kw_CONTINUE		# ControlContinue
	| Kw_RETURN expr?	# ControlReturn;
