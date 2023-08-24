grammar Swift;

options {
	tokenVocab = SwiftLexer;
}

idChain: ID (DOT ID)* # IDChain;

program: block EOF;

block: (statement)*;

// Statements
statement:
	variableAssignment
	| variableDeclaration
	| ifStatement
	| whileStatement
	| switchStatement
	| forStatement
	| guardStatement
	| controlFlowStatement
	| functionDeclarationStatement
	| functionCall
	| vectorDeclaration
	| vectorAssignment
	| matrixDeclaration
	| structDeclaration;

// Variable types
variableType: Kw_INT | Kw_FLOAT | Kw_BOOL | Kw_STRING | Kw_CHAR;

// Variable declaration

variableDeclaration:
	varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_ASSIGN expr (
		SEMICOLON
	)?																# TypeValueDeclaration
	| varType = (Kw_LET | Kw_VAR) ID Op_ASSIGN expr (SEMICOLON)?	# ValueDeclaration
	| varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_TERNARY (
		SEMICOLON
	)? # TypeDeclaration;

// Function declaration
functionDeclarationStatement:
	Kw_FUNC ID LPAREN functionParameters? RPAREN functionReturnType? LBRACE block RBRACE;

functionParameters:
	functionParameter (COMMA functionParameter)*;

functionParameter: ID? ID COLON Kw_INOUT? variableType;

functionReturnType: Op_ARROW variableType;

// Function call 
functionCall:
	(idChain | variableType) LPAREN functionCallArguments? RPAREN;

functionCallArguments:
	expr (COMMA expr)*						# Arguments
	| ID COLON expr (COMMA ID COLON expr)*	# NamedArguments;

// Variable assignment
variableAssignment:
	ID op = (Op_ASSIGN | Op_PLUS_ASSIGN | Op_MINUS_ASSIGN) expr;

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
	varType = (Kw_LET | Kw_VAR) ID COLON LBRACKET variableType RBRACKET Op_ASSIGN vectorDefinition;

vectorDefinition:
	LBRACKET vectorValues? RBRACKET	# VectorListValue
	| expr							# VectorSingleValue;

vectorValues: expr (COMMA expr)*;

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

// Structs

structDeclaration: Kw_STRUCT ID LBRACE structBody RBRACE;

structBody: structProperty*;

structProperty:
	variableDeclaration
	| Kw_MUTATING? functionDeclarationStatement;

// Expressions
expr:
	functionCall													# FunctionCallExpr
	| vectorAccess													# VectorAccessExpr
	| Op_MINUS expr													# UnaryExpr
	| Op_NOT right = expr											# NotExpr
	| left = expr op = (Op_MUL | Op_DIV) right = expr				# ArithmeticExpr
	| left = expr op = (Op_PLUS | Op_MINUS) right = expr			# ArithmeticExpr
	| left = expr op = Op_MOD right = expr							# ArithmeticExpr
	| left = expr op = (Op_GE | Op_GT) right = expr					# ComparasionExpr
	| left = expr op = (Op_LE | Op_LT) right = expr					# ComparasionExpr
	| left = expr op = (Op_EQ | Op_NEQ) right = expr				# ComparasionExpr
	| condition = expr Op_TERNARY cTrue = expr COLON cFalse = expr	# TernaryExpr
	| left = expr op = Op_AND right = expr							# LogicalExpr
	| left = expr op = Op_OR right = expr							# LogicalExpr
	| left = expr Kw_RANGE right = expr								# RangeExpr
	| LPAREN expr RPAREN											# ParExpr
	| INT															# IntExpr
	| idChain														# IdExpr
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
