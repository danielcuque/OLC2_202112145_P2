package interfaces

import (
	"fmt"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseSwiftVisitor
	Scope  *ScopeTree
	Errors []error
}

func NewVisitor() *Visitor {
	return &Visitor{
		Scope:  NewScopeTree(),
		Errors: make([]error, 0),
	}
}

func NewEvaluator(input string) *Visitor {
	lexer := parser.NewSwiftLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewSwiftParser(stream)

	parser.BuildParseTrees = true

	tree := parser.Program()

	visitor := NewVisitor()
	visitor.Visit(tree)

	return visitor
}

func (v *Visitor) NewError(msg string) {
	v.Errors = append(v.Errors, fmt.Errorf(msg))
}
