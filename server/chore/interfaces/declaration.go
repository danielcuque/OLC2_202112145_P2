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

func (v *Visitor) VisitVectorAccess(ctx *parser.VectorAccessContext) interface{} {

	ids := v.Visit(ctx.IdChain()).([]antlr.TerminalNode)
	id := ids[0].GetText()

	props := v.GetProps(ids)

	object, okV := v.LookUpObject(id, props, ctx.GetStart())

	if !okV {
		return nil
	}

	// Check if the object is a vector
	if object.GetType() != V.VectorType {
		v.NewError(fmt.Sprintf("El objeto %s no es un vector", id), ctx.GetStart())
		return nil
	}

	// Get the index
	index := v.Visit(ctx.Expr()).(V.IValue)

	// Check if the index is an integer
	if index.GetType() != V.IntType {
		v.NewError("El indice debe ser un entero", ctx.GetStart())
		return nil
	}

	// Get the value
	value := object.GetValue().(*VectorV).Get(index.GetValue().(int))

	if value == nil {
		v.NewError("El índice esta fuera de rango", ctx.GetStart())
		return nil
	}
	// Return vector, index and value, then, handle when is called as expression, all this as dictionary

	dict := map[string]interface{}{
		"vector": object,
		"index":  index.GetValue().(int),
		"value":  value,
	}

	return dict
}

func (v *Visitor) VisitVectorAssignment(ctx *parser.VectorAssignmentContext) interface{} {
	// Get the vector, then, get the index and the value
	values, ok := v.Visit(ctx.VectorAccess()).(map[string]interface{})

	if !ok {
		v.NewError(InvalidVectorValue, ctx.GetStart())
		return nil
	}

	vector := values["vector"].(*ObjectV)

	expr := v.Visit(ctx.Expr()).(V.IValue)

	// Check is not necessary, because the index is checked in the vector access

	// Check if the value type is the same as the vector type
	if expr.GetType() != vector.GetValue().(*VectorV).GetType() {
		v.NewError(fmt.Sprintf("El tipo de dato del vector debe ser %s", vector.GetValue().(*VectorV).GetType()), ctx.GetStart())
		return nil
	}
	index := values["index"].(int)
	vectorValue := vector.GetValue().(*VectorV).Get(index)
	op := ctx.GetOp().GetText()

	if op != "=" {
		vector.GetValue().(*VectorV).Set(index, v.arithmeticOp(vectorValue, expr, string(op[0]), ctx.GetStart()).(V.IValue))
	} else {
		vector.GetValue().(*VectorV).Set(index, expr)
	}

	return nil
}

// Matrix

/*
Consideraciones:
- La declaración del tamaño puede ser explícita o en base a se definición.
- Si la declaración es explícita pero su definición no es acorde a esta declaración se
debe marcar como un error. Por lo tanto se deben verificar que la cantidad de
dimensiones sea acorde a la definida.
- La asignación y lectura valores se realizará con la notación [ ]
- Los índices de declaración comienzan a partir de 1
- Los índices de acceso comienzan a partir de 0
- Las matrices no van a cambiar su tamaño durante la ejecución.
- Si se hace un acceso con índices en fuera de rango se devuelve nil y se debe
notificar como un error.
- Si se declara una matriz con índices negativos o 0, será considerado un error
- El atributo count solo recibirá número enteros en forma de literales, no podrán ser
asignadas ni variables ni elementos de otras estructuras a este atributo.
*/
func (v *Visitor) VisitMatrixListDeclaration(ctx *parser.MatrixListDeclarationContext) interface{} {
	fmt.Println("VisitMatrixListDeclaration")
	return nil
}

func (v *Visitor) VisitMatrixRepeatingDeclaration(ctx *parser.MatrixRepeatingDeclarationContext) interface{} {
	fmt.Println("VisitMatrixRepeatingDeclaration")
	return nil
}

func (v *Visitor) VisitMatrixTypeNested(ctx *parser.MatrixTypeNestedContext) interface{} {
	return nil
}

func (v *Visitor) VisitMatrixTypeSingle(ctx *parser.MatrixTypeSingleContext) interface{} {
	return nil
}

func (v *Visitor) VisitMatrixDefinition(ctx *parser.MatrixDefinitionContext) interface{} {
	return nil
}

func (v *Visitor) VisitMatrixValues(ctx *parser.MatrixValuesContext) interface{} {
	return nil
}

func (v *Visitor) VisitMatrixRepeatingDefinitionNested(ctx *parser.MatrixRepeatingDefinitionNestedContext) interface{} {
	return nil
}

func (v *Visitor) VisitMatrixRepeatingDefinitionSingle(ctx *parser.MatrixRepeatingDefinitionSingleContext) interface{} {
	return nil
}

// Structs

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {
	fmt.Println("VisitStructDeclaration")
	return nil
}

func (v *Visitor) VisitStructBody(ctx *parser.StructBodyContext) interface{} {
	return nil
}

func (v *Visitor) VisitStructProperty(ctx *parser.StructPropertyContext) interface{} {
	return nil
}
