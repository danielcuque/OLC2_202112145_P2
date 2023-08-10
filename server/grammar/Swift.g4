grammar Swift;

options {
	tokenVocab = SwiftLexer;
}

program: block EOF;

block: (statement)*;

// Statements
statement:
	variableAssignment
	| variableDeclaration
	| ifStatement;
//  | whileStatement; | forStatement;

// Variable types
variableType:
	Kw_INT
	| Kw_FLOAT
	| Kw_BOOL
	| Kw_STRING
	| Kw_CHAR
	| Kw_NIL;

// Variable declaration

variableDeclaration:
	varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_ASSIGN expr (
		SEMICOLON
	)?																# TypeValueDeclaration
	| varType = (Kw_LET | Kw_VAR) ID Op_ASSIGN expr (SEMICOLON)?	# ValueDeclaration
	| varType = (Kw_LET | Kw_VAR) ID COLON variableType Op_TERNARY (
		SEMICOLON
	)? # TypeDeclaration;

// Variable assignment
variableAssignment:
	ID op = (Op_ASSIGN | Op_PLUS_ASSIGN | Op_MINUS_ASSIGN) expr;

// If statement 
ifStatement:
	Kw_IF expr LBRACE block RBRACE									# SimpleIfStatement
	| Kw_IF expr LBRACE block RBRACE Kw_ELSE LBRACE block RBRACE	# IfElseStatement
	| Kw_IF expr LBRACE block RBRACE Kw_ELSE ifStatement			# IfElseIfStatement;

// LBRACE block RBRACE ifStatementTail;

// ifStatementTail: elseIfTail elseStatement | elseStatement | elseIfTail;

// elseStatement: Kw_ELSE LBRACE block RBRACE;

// elseIfTail: elseIfTail elseIf | elseIf;

// elseIf: Kw_ELSE Kw_IF LPAREN expr RPAREN LBRACE block RBRACE;

// whileStatement: Kw_WHILE LPAREN expr RPAREN LBRACE block RBRACE;

// forStatement: Kw_FOR expr Kw_IN expr LBRACE block RBRACE;

// Expressions
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
	| left = expr Kw_RANGE right = expr								# RangeExpr
	| LPAREN expr RPAREN											# ParExpr
	| INT															# IntExpr
	| ID															# IdExpr
	| FLOAT															# FloatExpr
	| STRING														# StrExpr
	| CHAR															# CharExpr
	| BOOL															# BoolExpr;
