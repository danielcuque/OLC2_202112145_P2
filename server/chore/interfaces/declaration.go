package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Variables

func (v *Visitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	isConstant := ctx.GetVarType().GetText()
	id := ctx.ID().GetText()
	value, okVal := v.Visit(ctx.Expr()).(V.IValue)

	if !okVal {
		v.NewError(InvalidExpressionError, ctx.GetStart())
		return nil
	}

	_, ok := v.Scope.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return false
	}

	// Get line, column and scope
	line, column, scope := GetVariableAttr(v, ctx.GetStart())

	newVariable := NewVariable(id, isConstant == "let", value, value.GetType(), line, column, scope)

	v.Scope.AddVariable(id, newVariable)

	return true
}

func (v *Visitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	isConstant := ctx.GetVarType().GetText() == "let"
	id := ctx.ID().GetText()
	value, okVal := v.Visit(ctx.Expr()).(V.IValue)
	valueType := v.Visit(ctx.VariableType()).(string)

	if !okVal {
		v.NewError(InvalidExpressionError, ctx.GetStart())
		return nil
	}

	_, ok := v.Scope.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return false
	}

	// Check if the explicit type is the same as the value type, except if explicit type is Float and value type is Int
	if valueType != value.GetType() {
		// Check if the explicit type is Float and the value type is Int
		// Change the value type to Float
		if valueType == V.FloatType && value.GetType() == V.IntType {
			value = V.NewFloatValue(float64(value.GetValue().(int)))
		} else {
			v.NewError(fmt.Sprintf("El tipo de la variable %s no coincide con el valor asignado, se esperaba %s y se obtuvo %s", id, valueType, value.GetType()), ctx.GetStart())
			return false
		}

	}

	// Get line, column and scope
	line, column, scope := GetVariableAttr(v, ctx.GetStart())

	newVariable := NewVariable(id, isConstant, value, valueType, line, column, scope)

	v.Scope.AddVariable(id, newVariable)

	return true
}

func (v *Visitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	// Declaration without value

	isConstant := ctx.GetVarType().GetText() == "let"
	id := ctx.ID().GetText()
	valueType := v.Visit(ctx.VariableType()).(string)

	_, ok := v.Scope.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return false
	}

	if isConstant {
		v.NewError(fmt.Sprintf("La variable %s debe ser inicializada", id), ctx.GetStart())
		return false
	}

	// Get line, column and scope
	line, column, scope := GetVariableAttr(v, ctx.GetStart())
	newVariable := NewVariable(id, isConstant, V.NewNilValue(nil), valueType, line, column, scope)

	v.Scope.AddVariable(id, newVariable)

	return true

}

func GetVariableAttr(v *Visitor, lc antlr.Token) (int, int, *ScopeNode) {
	line := lc.GetLine()
	column := lc.GetColumn()
	scope := v.Scope.GetCurrentScope()
	return line, column, scope
}

// Vectors

func (v *Visitor) VisitVectorDeclaration(ctx *parser.VectorDeclarationContext) interface{} {

	dataList, okVal := v.Visit(ctx.VectorDefinition()).([]V.IValue) // [1,2,3]

	if !okVal {
		v.NewError(InvalidVectorValueError, ctx.GetStart())
		return nil
	}

	id := ctx.ID().GetText() // a

	_, ok := v.Scope.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return nil
	}

	line, column, scope := GetVariableAttr(v, ctx.GetStart())

	isConstant := ctx.GetVarType().GetText() == "const" // let | var

	valueType := v.Visit(ctx.VariableType()).(string) // Id: int | float | string

	// Verify that dataList is not empty

	if len(dataList) != 0 && dataList[0].GetType() != valueType {
		v.NewError(fmt.Sprintf("El tipo de dato del vector debe ser %s", valueType), ctx.GetStart())
		return nil
	}

	// New Vector Variable

	newVector := NewVector(valueType, dataList)

	v.Scope.AddVariable(id,
		NewVariable(id, isConstant, newVector, valueType, line, column, scope),
	)

	return nil
}

func (v *Visitor) VisitVectorListValue(ctx *parser.VectorListValueContext) interface{} {
	if ctx.VectorValues() == nil {
		return make([]V.IValue, 0)
	}

	return v.Visit(ctx.VectorValues())
}

func (v *Visitor) VisitVectorSingleValue(ctx *parser.VectorSingleValueContext) interface{} {
	value, ok := v.Visit(ctx.Expr()).(*VectorV)

	if !ok {
		v.NewError(InvalidVectorValueError, ctx.GetStart())
		return make([]V.IValue, 0)
	}

	return value.GetValue()
}

func (v *Visitor) VisitVectorValues(ctx *parser.VectorValuesContext) interface{} {
	values := make([]V.IValue, 0)

	// Verify that all values are the same type
	var valueType string

	for _, value := range ctx.AllExpr() {
		value, ok := v.Visit(value).(V.IValue)

		if !ok {
			v.NewError(InvalidVectorValueError, ctx.GetStart())
			return make([]V.IValue, 0)
		}

		valueType = value.GetType()

		// Verify that all values are the same type
		if len(values) != 0 && values[0].GetType() != valueType {
			v.NewError("La lista de expresiones deben ser todas del mismo tipo", ctx.GetStart())
			return make([]V.IValue, 0)
		}

		values = append(values, value)
	}

	return values
}
