package main

import (
	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {

	is := antlr.NewInputStream("2+3")

	lexer := parser.NewSwiftLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	tokenNames := lexer.GetSymbolicNames()

	// Print the tokens readed and the type of each one
	for {
		t := stream.LT(1)
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		tokenName := tokenNames[t.GetTokenType()]
		println(t.GetText(), " -> ", tokenName)
		stream.Consume()
	}

}
