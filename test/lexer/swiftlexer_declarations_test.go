package lexer

import (
	"testing"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

func TestSwiftLexerDeclarations(t *testing.T) {
	testCases := []struct {
		input          string
		expectedTokens []int
	}{
		{
			input:          "var a = 2",
			expectedTokens: []int{parser.SwiftLexerKw_VAR, parser.SwiftLexerID, parser.SwiftLexerOp_ASSIGN, parser.SwiftLexerINT},
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
