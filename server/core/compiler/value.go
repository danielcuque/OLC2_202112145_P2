package compiler

import "fmt"

type Value struct {
	Value        interface{}
	StackAddress int
}

func NewValue(value interface{}, stackAddress int) *Value {
	return &Value{
		Value:        value,
		StackAddress: stackAddress,
	}
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

func (v *ValueResponse) GetValue() interface{} {
	if v.ContextValue == TemporalType {
		return fmt.Sprintf("t%d", v.Value.(*Temporal).Index)
	}
	return v.Value
}

func (v *ValueResponse) GetType() string {
	return v.Type
}
