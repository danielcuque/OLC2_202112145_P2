package compiler

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
	Type        string
	Value       interface{}
	ContextType ContextType
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

func (v *ValueResponse) GetType() string {
	return v.Type
}
