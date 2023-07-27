package main

import (
	I "OLC2/chore/interfaces"
	"OLC2/chore/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Read file to test
	input, _ := antlr.NewFileStream("./examples/test.swift")
	lexer := parser.NewSwiftLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSwiftParser(stream)

	p.BuildParseTrees = true
	tree := p.Program()
	eval := I.NewVisitor()
	eval.Visit(tree)

	fmt.Println(eval.Memory)

}
