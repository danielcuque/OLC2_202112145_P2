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

// Standar error message

const (
	InvalidExpressionError            = "No se pudo obtener el valor de la expresión"
	InvalidParameterError             = "No se pudo obtener el valor del parámetro"
	FunctionNotFoundError             = "No se encontró la función en este ambiente"
	FunctionAlreadyExistsError        = "Ya existe una función con este nombre en este ambiente"
	InvalidReturnTypeDeclarationError = "El tipo de retorno no es válido"
	InvalidParameterTypeError         = "El tipo del parámetro no es válido"
	InvalidNumberOfParametersError    = "El número de parámetros no es válido"
	InvalidReturnTypeFunctionError    = "El tipo de retorno de la función no es válido"
	InvalidArgumentNameError          = "El nombre del argumento no es existe"
	InvalidFunctionBodyError          = "El cuerpo de la función no es válido"
)
