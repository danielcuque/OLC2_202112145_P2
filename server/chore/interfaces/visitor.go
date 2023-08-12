package interfaces

import (
	"fmt"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseSwiftVisitor
	Errors []*VisitorError
	Logs   []string
	Scope  *ScopeTree
	Stack  *Stack
}

func NewVisitor() *Visitor {
	return &Visitor{
		Errors: make([]*VisitorError, 0),
		Logs:   make([]string, 0),
		Scope:  NewScopeTree(),
		Stack:  NewStack(),
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

func (v *Visitor) NewError(msg string, ctx antlr.Token) {
	errorMsg := fmt.Sprintf("Error(%d:%d): %s ", ctx.GetLine(), ctx.GetColumn(), msg)
	v.Errors = append(v.Errors, NewVisitorError(ctx.GetLine(), ctx.GetColumn(), errorMsg))
}
