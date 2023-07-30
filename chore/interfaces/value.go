package interfaces

type ValueType int

const (
	STRING ValueType = iota
	INT
	DOUBLE
	BOOL
)

type Value struct {
	Type       ValueType
	ParseValue interface{}
}

type IValue interface {
	GetValue(scope Scope) interface{}
	GetType(Scope Scope) ValueType
}

func (v Value) GetValue(scope Scope) interface{} {
	return v.ParseValue
}
