package interfaces

type TypeError string

// Lexical | Syntactic | Semantic

const (
	Lexical   TypeError = "Lexical"
	Syntactic TypeError = "Syntactic"
	Semantic  TypeError = "Semantic"
)

type VisitorError struct {
	Line   int
	Column int
	Msg    string
	Type   TypeError
}

func (e *VisitorError) Error() string {
	return e.Msg
}

func (e *VisitorError) GetLine() int {
	return e.Line
}

func (e *VisitorError) GetColumn() int {
	return e.Column
}

func NewVisitorError(line int, column int, msg string) *VisitorError {
	return &VisitorError{
		Line:   line,
		Column: column,
		Msg:    msg,
	}
}
