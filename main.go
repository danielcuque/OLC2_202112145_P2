package main

import (
	"fmt"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseSwiftVisitor
}

func main() {
	expression := "(1+2)*3+(3+2)"

	input := antlr.NewInputStream(expression)
	lexer := parser.NewSwiftLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSwiftParser(stream)

	p.BuildParseTrees = true
	tree := p.Expr()
	var visitor = Visitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, "=", result)

}
