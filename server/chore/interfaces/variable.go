package interfaces

// Create struct to handle Variables

type Variable struct {
	Name    string
	IsConst bool
	Value   IValue
	Type    string
}

func NewVariable(name string, isConst bool, value IValue, typeVar string) *Variable {
	return &Variable{
		Name:    name,
		IsConst: isConst,
		Value:   value,
		Type:    typeVar,
	}
}

func (v *Variable) GetValue() interface{} {
	return v.Value.GetValue()
}

func (v *Variable) SetValue(value IValue) {
	v.Value = value
}

func (v *Variable) IsConstant() bool {
	return v.IsConst
}

func (v *Variable) GetName() string {
	return v.Name
}

func (v *Variable) GetType() string {
	return v.Type
}
