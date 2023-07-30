package interfaces

import (
	"OLC2/chore/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Value struct {
	ParseValue interface{}
}

type Visitor struct {
	parser.SwiftVisitor
	Memory map[string]Value
	Errors []error
}

func NewVisitor() *Visitor {
	return &Visitor{
		Memory: make(map[string]Value),
		Errors: make([]error, 0),
	}
}

func NewEvaluator(input string) *Visitor {
	lexer := parser.NewSwiftLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSwiftParser(stream)

	p.BuildParseTrees = true
	tree := p.Program()
	eval := NewVisitor()
	eval.Visit(tree)

	return eval
}

func (v *Visitor) NewError(msg string) {
	v.Errors = append(v.Errors, fmt.Errorf(msg))
}
