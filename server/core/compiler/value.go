package compiler

type Value struct {
	Type  string
	Value interface{}
}

type ContextValue string

const (
	LiteralType  ContextValue = "literal"
	TemporalType ContextValue = "temporal"
	PointerType  ContextValue = "pointer"
)

type ValueResponse struct {
	Type         string
	Value        interface{}
	ContextValue ContextValue
}
