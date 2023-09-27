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
	return v.Value
}

func (v *ValueResponse) SetValue(value interface{}) {
	v.Value = value
}

func (v *ValueResponse) GetContextValue() ContextValue {
	return v.ContextValue
}

func (v *ValueResponse) GetType() string {
	return v.Type
}
