package test

import (
	"testing"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

type testCases struct {
	input          string
	expectedTokens []int
}

func TestSwiftLexerKeywordDataTypes(t *testing.T) {
	testCases := []testCases{
		{
			input: "Int Double String Bool",
			expectedTokens: []int{
				parser.SwiftLexerKw_INT,
				parser.SwiftLexerKw_DOUBLE,
				parser.SwiftLexerKw_STRING,
				parser.SwiftLexerKw_BOOL,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerLiteralDataTypes(t *testing.T) {
	testCases := []testCases{
		{
			input: `
			42               // Integer literal
			3.14159          // Floating-point literal
			"Hello, world!"  // String literal
			true             // Boolean literal
			`,
			expectedTokens: []int{
				parser.SwiftLexerINT,
				parser.SwiftLexerDOUBLE,
				parser.SwiftLexerSTRING,
				parser.SwiftLexerBOOL,
			},
		},
	}

	TraverseCases(t, testCases)

}

func TraverseCases(t *testing.T, testCases []testCases) {

	for _, tc := range testCases {
		input := antlr.NewInputStream(tc.input)
		lexer := parser.NewSwiftLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		tokenNames := lexer.GetSymbolicNames()

		for _, expected := range tc.expectedTokens {
			token := stream.LT(1)
			if token.GetTokenType() == antlr.TokenEOF {
				t.Errorf("Esperado token %s pero se alcanzó el fin de archivo", tokenNames[expected])
				break
			}
			if token.GetTokenType() != expected {
				t.Errorf("Esperado token %s pero se obtuvo %s", tokenNames[expected], tokenNames[token.GetTokenType()])
			}
			stream.Consume()
		}
	}
}
