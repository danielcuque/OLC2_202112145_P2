package test

import (
	"testing"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

func TestSwiftLexer(t *testing.T) {
	input := antlr.NewInputStream("2 + 3") // Cambia la entrada según lo que desees probar
	lexer := parser.NewSwiftLexer(input)   // Reemplaza "NewSwiftLexer" con el nombre de la clase del lexer generado
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Agrega aquí los valores numéricos de los tokens esperados en el orden en que se espera que aparezcan en la entrada
	expectedTokens := []int{
		parser.SwiftLexerINT,
		parser.SwiftLexerOp_PLUS,
		parser.SwiftLexerINT,
	}

	// Verifica que los tokens generados por el lexer coincidan con los tokens esperados
	for _, expected := range expectedTokens {
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
