grammar Swift;

// Rules for the Swift language
options {
	tokenVocab = SwiftLexer;
}

program: block EOF;

block: (statement)*;

statement: declaration;

declaration: (Kw_VAR | Kw_LET) ID (
		(COLON data_type)
		| (Op_ASSIGN expr)
	);

data_type: (Kw_INT | Kw_DOUBLE | Kw_DOUBLE | Kw_STRING | Kw_BOOL);

expr:
	left = expr op = (Op_LT | Op_GT | Op_LE | Op_GE) right = expr # Op
	// Less than, greater than, less than or equal, greater than or equal
	| left = expr op = (Op_MUL | Op_DIV) right = expr		# Op // Multiplication and division
	| left = expr op = (Op_PLUS | Op_MINUS) right = expr	# Op // Addition and subtraction
	| left = expr Op_MOD right = expr						# Op // Modulo
	| digit = INT											# Digit
	| ID													# Id
	| LPAREN expr RPAREN									# Paren;
