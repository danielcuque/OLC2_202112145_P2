grammar Swift;

// Rules for the Swift language
options {
	tokenVocab = SwiftLexer;
}

program: expr EOF;

expr:
	left = expr op = (Op_MUL | Op_DIV) right = expr			# Op
	| left = expr op = (Op_PLUS | Op_MINUS) right = expr	# Op
	| digit = INT											# Digit
	| LPAREN expr RPAREN									# Paren;
