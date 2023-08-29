package interfaces

import (
	V "OLC2/chore/values"

	"github.com/antlr4-go/antlr/v4"
)

// Create struct to handle Variables

type Variable struct {
	Name    string
	IsConst bool
	Value   V.IValue
	Type    string
	Line    int
	Column  int
	Env     *EnvNode
}

func NewVariable(v *Visitor, name string, isConst bool, value V.IValue, typeVar string, lc antlr.Token) *Variable {
	line := lc.GetLine()
	column := lc.GetColumn()

	return &Variable{
		Name:    name,
		IsConst: isConst,
		Value:   value,
		Type:    typeVar,
		Line:    line,
		Column:  column,
		Env:     v.Env.GetCurrentScope(),
	}
}

func (v *Variable) GetValue() interface{} {
	return v.Value.GetValue()
}

func (v *Variable) String() string {
	return v.Value.String()
}

func (v *Variable) SetValue(value V.IValue) {
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

func (v *Variable) Copy() V.IValue {
	return v.Value.Copy()
}

func (v *Variable) GetScopeName() string {
	return v.Env.GetType()
}
