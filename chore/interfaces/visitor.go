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
	Error  []error
}

func NewVisitor() *Visitor {
	return &Visitor{
		Memory: make(map[string]Value),
		Error:  make([]error, 0),
	}
}

func NewEvaluator(input string) map[string]Value {
	lexer := parser.NewSwiftLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSwiftParser(stream)

	p.BuildParseTrees = true
	tree := p.Program()
	eval := NewVisitor()
	eval.Visit(tree)

	return eval.Memory
}

func (v *Visitor) NewError(msg string) {
	v.Error = append(v.Error, fmt.Errorf(msg))
}
