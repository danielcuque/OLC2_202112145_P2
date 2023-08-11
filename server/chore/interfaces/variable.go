package interfaces

// Create struct to handle Variables

type Variable struct {
	Name    string
	IsConst bool
	Value   IValue
	Type    string
	Line    int
	Column  int
	Scope   *ScopeNode
}

func NewVariable(name string, isConst bool, value IValue, typeVar string, line, column int, scope *ScopeNode) *Variable {
	return &Variable{
		Name:    name,
		IsConst: isConst,
		Value:   value,
		Type:    typeVar,
		Line:    line,
		Column:  column,
		Scope:   scope,
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

func (v *Variable) GetLine() int {
	return v.Line
}

func (v *Variable) GetColumn() int {
	return v.Column
}

func (v *Variable) GetScopeName() string {
	return v.Scope.GetType()
}
