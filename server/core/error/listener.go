package error

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

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
