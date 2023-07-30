grammar Swift;

options {
	tokenVocab = SwiftLexer;
}

program: block EOF;

block: (statement)*;

statement:
	variableAssignment
	| variableDeclaration
	| ifStatement
	| whileStatement
	| forStatement;

variableType: Kw_INT | Kw_DOUBLE | Kw_BOOL | Kw_STRING | Kw_NIL;

variableCase: (Kw_LET | Kw_VAR);

variableDeclaration:
	variableCase ID Op_ASSIGN expr (SEMICOLON)?;

variableAssignment: ID Op_ASSIGN expr;

ifStatement: Kw_IF LPAREN expr RPAREN LBRACE block RBRACE;

whileStatement: Kw_WHILE LPAREN expr RPAREN LBRACE block RBRACE;

forStatement: Kw_FOR expr Kw_IN expr LBRACE block RBRACE;

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
	| DOUBLE														# DoubleExpr
	| STRING														# StrExpr
	| BOOL															# BoolExpr;
