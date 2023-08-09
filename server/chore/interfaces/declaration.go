package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	varType := ctx.GetVarType().GetText()
	id := ctx.ID().GetText()
	value, okVal := v.Visit(ctx.Expr()).(IValue)

	if !okVal {
		return nil
	}

	_, ok := v.Scope.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("Error: La variable %s ya existe", id))
		return false
	}

	newVariable := NewVariable(id, varType == "let", value, value.GetType())

	v.Scope.AddVariable(id, newVariable)

	return true
}

func (v *Visitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	varType := ctx.GetVarType().GetText()
	id := ctx.ID().GetText()
	value, okVal := v.Visit(ctx.Expr()).(IValue)
	valueType := v.Visit(ctx.VariableType()).(string)

	if !okVal {
		return nil
	}

	_, ok := v.Scope.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("Error: La variable %s ya existe", id))
		return false
	}

	// Check if the explicit type is the same as the value type, except if explicit type is Float and value type is Int
	if valueType != value.GetType() {
		// Check if the explicit type is Float and the value type is Int
		// Change the value type to Float
		if valueType == floatT && value.GetType() == intT {
			value = NewFloatValue(float64(value.GetValue().(int)))
		} else {
			v.NewError(fmt.Sprintf("Error: El tipo de la variable %s no coincide con el valor asignado, se esperaba %s y se obtuvo %s", id, valueType, value.GetType()))
			return false
		}

	}

	newVariable := NewVariable(id, varType == "let", value, valueType)

	v.Scope.AddVariable(id, newVariable)

	return true
}

func (v *Visitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	return nil
}
