package compiler

import (
	"fmt"
)

type Value struct {
	Value        interface{}
	StackAddress int
	Type         TemporalCast
}

func NewValue(value interface{}, stackAddress int, Type TemporalCast) *Value {
	return &Value{
		Value:        value,
		StackAddress: stackAddress,
		Type:         Type,
	}
}

func (v *Value) GetValue() interface{} {
	return v.Value
}

func (v *Value) GetAddress() int {
	return v.StackAddress
}

func (v *Value) GetType() TemporalCast {
	return v.Type
}

type ContextType string

const (
	LiteralType  ContextType = "literal"
	TemporalType ContextType = "temporal"
	PointerType  ContextType = "pointer"
	LabelType    ContextType = "label"
)

func DefaultNil() string {
	return "9999999"
}

type ValueResponse struct {
	Type        TemporalCast
	Value       interface{}
	ContextType ContextType
}

func (v *ValueResponse) ToPrint() string {
	if v.ContextType == TemporalType {
		temporal := v.Value.(*Temporal)
		return fmt.Sprintf("printf(\"%%%s\", %s);", temporal.Type, temporal.Cast())
	}

	if v.ContextType == LiteralType {
		return fmt.Sprintf("printf(\"%%%s\", %s);", v.Type, v.Value)
	}

	return ""
}

func (v *ValueResponse) Cast() string {
	if v.ContextType == TemporalType {
		temporal := v.Value.(*Temporal)
		return temporal.Cast()
	}

	if v.ContextType == LiteralType {
		switch v.Type {
		case FloatTemporal:
			return fmt.Sprintf("(float) %s", v.Value)
		case IntTemporal:
			return fmt.Sprintf("(int) %s", v.Value)
		case CharTemporal:
			return fmt.Sprintf("(char) %s", v.Value)
		}
	}

	return ""
}

func (v *ValueResponse) GetValue() interface{} {
	return v.Value
}

func (v *ValueResponse) SetValue(value interface{}) {
	v.Value = value
}

func (v *ValueResponse) GetContextValue() ContextType {
	return v.ContextType
}

func (v *ValueResponse) GetType() TemporalCast {
	return v.Type
}
