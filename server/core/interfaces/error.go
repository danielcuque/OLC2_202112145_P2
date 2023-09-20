package interfaces

type TypeError string

// Lexical | Syntactic | Semantic

const (
	Lexical   TypeError = "Léxico"
	Syntactic TypeError = "Sintáctico"
	Semantic  TypeError = "Semántico"
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

func NewVisitorError(line int, column int, msg string, typeError TypeError) *VisitorError {
	return &VisitorError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Type:   typeError,
	}
}

// Standar error message

const (
	InvalidExpression         = "No se pudo obtener el valor de la expresión"
	InvalidParameter          = "No se pudo obtener el valor del parámetro"
	InvalidMatrixValue        = "No se pudo obtener el valor de la matriz"
	FunctionNotFound          = "No se encontró la función o struct en este ambiente"
	FunctionAlreadyExists     = "Ya existe una función con este nombre en este ambiente"
	InvalidParameterType      = "El tipo del parámetro no es válido"
	InvalidNumberOfParameters = "El número de parámetros no es válido"
	InvalidReturnTypeFunction = "El tipo de retorno de la función no es válido"
	InvalidArgumentName       = "El nombre del argumento no es existe"
	InvalidFunctionBody       = "El cuerpo de la función no es válido"
	InvalidArgument           = "El argumento no es válido"
	InvalidVectorValue        = "No se pudo obtener el valor del vector"
	ObjectNotFound            = "No se encontró el objeto en este ambiente"
	MethodNotFound            = "No se encontró el método para este objeto"
)
