package interfaces

import (
	"fmt"

	"OLC2/core/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseSwiftVisitor
	Errors []*VisitorError
	Logs   []string
	Env    *EnvTree
	Stack  *Stack
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors    []*VisitorError
	TypeError TypeError
}

func (e *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e2 antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("Error(%d:%d): %s ", line, column, msg)
	e.Errors = append(e.Errors, NewVisitorError(line, column, errorMsg, e.TypeError))
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		Errors: make([]*VisitorError, 0),
	}
}

func NewVisitor() *Visitor {
	return &Visitor{
		Errors: make([]*VisitorError, 0),
		Logs:   make([]string, 0),
		Env:    NewEnvTree(),
		Stack:  NewStack(),
	}
}

func NewEvaluator(input string) *Visitor {
	errorListener := NewErrorListener()
	errorListener.TypeError = Lexical

	lexer := parser.NewSwiftLexer(antlr.NewInputStream(input))

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewSwift(stream)

	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)
	errorListener.TypeError = Syntactic

	parser.BuildParseTrees = true

	tree := parser.Program()

	if len(errorListener.Errors) > 0 {
		errorVisitor := NewVisitor()
		errorVisitor.Errors = errorListener.Errors
		return errorVisitor
	}

	visitor := NewVisitor()
	visitor.Visit(tree)

	return visitor
}

func (v *Visitor) NewError(msg string, ctx antlr.Token) {
	errorMsg := fmt.Sprintf("Error(%d:%d): %s ", ctx.GetLine(), ctx.GetColumn(), msg)
	v.Errors = append(v.Errors, NewVisitorError(ctx.GetLine(), ctx.GetColumn(), errorMsg, Semantic))
}

func (v *Visitor) NewLog(msg string) {
	v.Logs = append(v.Logs, msg)
}

func (v *Visitor) GetLogs() string {
	var logs string
	for _, log := range v.Logs {
		logs += log + "\n"
	}
	return logs
}
