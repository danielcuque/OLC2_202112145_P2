package interpreter

import (
	"fmt"

	C "OLC2/core/compiler"
	E "OLC2/core/error"
	"OLC2/core/parser"
	"OLC2/core/static"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseSwiftVisitor
	Errors []*E.VisitorError
	Logs   []string
	Env    *EnvTree
	Stack  *Stack
}

func NewVisitor() *Visitor {
	return &Visitor{
		Errors: make([]*E.VisitorError, 0),
		Logs:   make([]string, 0),
		Env:    NewEnvTree(),
		Stack:  NewStack(),
	}
}

func NewEvaluator(input string) (*C.Compiler, *Visitor) {
	errorListener := E.NewErrorListener()
	errorListener.TypeError = E.Lexical

	lexer := parser.NewSwiftLexer(antlr.NewInputStream(input))

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewSwift(stream)

	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)
	errorListener.TypeError = E.Syntactic

	parser.BuildParseTrees = true

	tree := parser.Program()

	if len(errorListener.Errors) > 0 {
		semanticChecker := NewVisitor()
		semanticChecker.Errors = errorListener.Errors
		return C.NewCompiler(0), semanticChecker
	}

	checker := NewVisitor()
	checker.Visit(tree)

	staticVisitor := static.NewStaticVisitor()
	staticVisitor.Visit(tree)

	compiler := C.NewCompiler(staticVisitor.Counter)
	// compiler.Env = staticVisitor.Env
	compiler.Visit(tree)

	return compiler, checker
}

func (v *Visitor) NewError(msg string, ctx antlr.Token) {
	errorMsg := fmt.Sprintf("Error(%d:%d): %s ", ctx.GetLine(), ctx.GetColumn(), msg)
	v.Errors = append(v.Errors, E.NewVisitorError(ctx.GetLine(), ctx.GetColumn(), errorMsg, E.Semantic))
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

func (v *Visitor) HasErrors() bool {
	return len(v.Errors) > 0
}

func (v *Visitor) GetErrors() []*E.VisitorError {
	return v.Errors
}
