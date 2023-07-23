package lexer

import (
	"testing"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

func TestSwiftLexerOperations(t *testing.T) {

	testCases := []struct {
		input          string
		expectedTokens []int
	}{
		{
			input:          "2 + 3",
			expectedTokens: []int{parser.SwiftLexerINT, parser.SwiftLexerOp_PLUS, parser.SwiftLexerINT},
		},
		{
			input:          "(2 + 3) * 4",
			expectedTokens: []int{parser.SwiftLexerLPAREN, parser.SwiftLexerINT, parser.SwiftLexerOp_PLUS, parser.SwiftLexerINT, parser.SwiftLexerRPAREN, parser.SwiftLexerOp_MUL, parser.SwiftLexerINT},
		},
	}

	for _, tc := range testCases {

		input := antlr.NewInputStream(tc.input)
		lexer := parser.NewSwiftLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		for _, expected := range tc.expectedTokens {
			token := stream.LT(1)
			if token.GetTokenType() == antlr.TokenEOF {
				t.Errorf("Esperado token %d pero se alcanzó el fin de archivo", expected)
				break
			}
			if token.GetTokenType() != expected {
				t.Errorf("Esperado token %d pero se obtuvo %d", expected, token.GetTokenType())
			}
			stream.Consume()
		}
	}
}
