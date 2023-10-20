package compiler

import (
	"fmt"
)

type Value struct {
	Value        interface{}
	StackAddress *StackIndex
	IsRelative   bool
	Type         TemporalCast
}

func NewSimpleValue(index int) *Value {
	return &Value{
		StackAddress: &StackIndex{
			Index: index,
		},
		IsRelative: false,
	}
}

func (v *Value) GetValue() interface{} {
	return v.Value
}

func (v *Value) GetAddress() string {
	return v.StackAddress.String()
}

func (v *Value) GetType() TemporalCast {
	return v.Type
}

func (v *Value) SetData(Type TemporalCast, value interface{}) {
	v.Type = Type
	v.Value = value
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
