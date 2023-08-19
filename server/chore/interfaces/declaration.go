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
		v.NewError(InvalidExpression, ctx.GetStart())
		return nil
	}

	_, ok := v.Env.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return false
	}

	// Get line, column and scope
	line, column, scope := v.GetVariableAttr(ctx.GetStart())

	newVariable := NewVariable(id, isConstant == "let", value, value.GetType(), line, column, scope)

	v.Env.AddVariable(id, newVariable)

	return true
}

func (v *Visitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	isConstant := ctx.GetVarType().GetText() == "let"
	id := ctx.ID().GetText()
	value, okVal := v.Visit(ctx.Expr()).(V.IValue)
	valueType := v.Visit(ctx.VariableType()).(string)

	if !okVal {
		v.NewError(InvalidExpression, ctx.GetStart())
		return nil
	}

	_, ok := v.Env.GetVariable(id).(Variable)

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
	line, column, scope := v.GetVariableAttr(ctx.GetStart())

	newVariable := NewVariable(id, isConstant, value, valueType, line, column, scope)

	v.Env.AddVariable(id, newVariable)

	return true
}

func (v *Visitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	// Declaration without value

	isConstant := ctx.GetVarType().GetText() == "let"
	id := ctx.ID().GetText()
	valueType := v.Visit(ctx.VariableType()).(string)

	_, ok := v.Env.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return false
	}

	if isConstant {
		v.NewError(fmt.Sprintf("La variable %s debe ser inicializada", id), ctx.GetStart())
		return false
	}

	// Get line, column and scope
	line, column, scope := v.GetVariableAttr(ctx.GetStart())
	newVariable := NewVariable(id, isConstant, V.NewNilValue(nil), valueType, line, column, scope)

	v.Env.AddVariable(id, newVariable)

	return true

}

func (v *Visitor) GetVariableAttr(lc antlr.Token) (int, int, *EnvNode) {
	line := lc.GetLine()
	column := lc.GetColumn()
	scope := v.Env.GetCurrentScope()
	return line, column, scope
}

// Vectors

func (v *Visitor) VisitVectorDeclaration(ctx *parser.VectorDeclarationContext) interface{} {

	dataList, okVal := v.Visit(ctx.VectorDefinition()).([]V.IValue) // [1,2,3]

	if !okVal {
		v.NewError(InvalidVectorValue, ctx.GetStart())
		return nil
	}

	id := ctx.ID().GetText()

	_, ok := v.Env.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return nil
	}

	line, column, scope := v.GetVariableAttr(ctx.GetStart())

	isConstant := ctx.GetVarType().GetText() == "const" // let | var

	valueType := v.Visit(ctx.VariableType()).(string) // Id: int | float | string

	// Verify that dataList is not empty

	if len(dataList) != 0 && dataList[0].GetType() != valueType {
		v.NewError(fmt.Sprintf("El tipo de dato del vector debe ser %s", valueType), ctx.GetStart())
		return nil
	}

	// New Vector Variable
	newVector := NewVector(valueType, dataList)

	// Create a new generic object
	newObj := NewObjectV(V.VectorType, newVector)

	// Add native properties
	count := NewVariable("count", true, V.NewIntValue(len(dataList)), V.IntType, line, column, scope)
	isEmpty := NewVariable("isEmpty", true, V.NewBooleanValue(len(dataList) == 0), V.BooleanType, line, column, scope)

	newObj.AddProp("count", count)
	newObj.AddProp("isEmpty", isEmpty)

	v.Env.AddVariable(id,
		NewVariable(id, isConstant, newObj, valueType, line, column, scope),
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
	value, ok := v.Visit(ctx.Expr()).(*ObjectV).GetValue().(V.IValue)

	if !ok {
		v.NewError(InvalidVectorValue, ctx.GetStart())
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
			v.NewError(InvalidVectorValue, ctx.GetStart())
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
