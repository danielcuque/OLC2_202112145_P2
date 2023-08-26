package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Variables

func (v *Visitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	isConstant := ctx.GetVarType().GetText() == "let"
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

	newVariable := NewVariable(v, id, isConstant, value, value.GetType(), ctx.GetStart())

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

	newVariable := NewVariable(v, id, isConstant, value, valueType, ctx.GetStart())

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

	newVariable := NewVariable(v, id, isConstant, V.NewNilValue(nil), valueType, ctx.GetStart())

	v.Env.AddVariable(id, newVariable)

	return true

}

// Vectors

func (v *Visitor) VisitVectorDeclaration(ctx *parser.VectorDeclarationContext) interface{} {

	id := ctx.ID().GetText()

	_, ok := v.Env.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return nil
	}

	dataList, okVal := v.Visit(ctx.VectorDefinition()).([]V.IValue) // [1,2,3]

	if !okVal {
		v.NewError(InvalidVectorValue, ctx.GetStart())
		return nil
	}

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
	count := NewVariable(v, "count", true, V.NewIntValue(len(dataList)), V.IntType, ctx.GetStart())
	isEmpty := NewVariable(v, "isEmpty", true, V.NewBooleanValue(len(dataList) == 0), V.BooleanType, ctx.GetStart())

	newObj.AddProp("count", count)
	newObj.AddProp("isEmpty", isEmpty)

	v.Env.AddVariable(id,
		NewVariable(v, id, isConstant, newObj, valueType, ctx.GetStart()),
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

func (v *Visitor) VisitMatrixDeclaration(ctx *parser.MatrixDeclarationContext) interface{} {
	id := v.Visit(ctx.IdChain()).([]antlr.TerminalNode)[0].GetText()

	_, ok := v.Env.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La variable %s ya existe", id), ctx.GetStart())
		return nil
	}

	// Get matrixType

	matrixType := v.Visit(ctx.MatrixType()).(string)
	body, ok := v.GetMatrixBody(ctx).(*MatrixNode)

	if !ok {
		return nil
	}

	matrixTypeDimensions := v.GetMatrixTypeDimensions(ctx.MatrixType().GetText())
	bodyDimension := v.GetMatrixDimensions(body, 0)

	if matrixTypeDimensions != bodyDimension {
		v.NewError(fmt.Sprintf("La cantidad de dimensiones de la matriz %s no coincide con la definición", id), ctx.GetStart())
		return nil
	}

	// Create a new generic object

	newObj := NewObjectV(V.MatrixType, body)

	// This is a matrix, so, matrix does not have native properties

	v.Env.AddVariable(id,
		NewVariable(v, id, true, newObj, matrixType, ctx.GetStart()),
	)

	return nil
}

func (v *Visitor) VisitMatrixTypeNested(ctx *parser.MatrixTypeNestedContext) interface{} {
	return v.Visit(ctx.MatrixType())
}

func (v *Visitor) VisitMatrixTypeSingle(ctx *parser.MatrixTypeSingleContext) interface{} {
	return v.Visit(ctx.VariableType())
}

// Returns Ivalues but MatrixNode structs
func (v *Visitor) VisitMatrixDefinition(ctx *parser.MatrixDefinitionContext) interface{} {
	if ctx.MatrixValues() != nil {
		return v.Visit(ctx.MatrixValues())
	}

	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitMatrixValues(ctx *parser.MatrixValuesContext) interface{} {

	// There are two cases, when matrix values are other matrix values or when they are expressions

	// Get the first matrix definition

	// Check if is matrixNode or expression, if is expr then, the matrix value ends
	expr, ok := v.Visit(ctx.MatrixDefinition(0)).(V.IValue)

	if !ok {
		return expr // If is not ok, then is a matrixNode
	}

	baseNode := NewMatrixNode(expr.GetType(), []V.IValue{expr})

	// Get the rest of the matrix definitions
	for _, matrixDef := range ctx.AllMatrixDefinition()[1:] {
		matrixDef, ok := v.Visit(matrixDef).(V.IValue)

		if !ok {
			return expr
		}

		if matrixDef.GetType() != baseNode.DataType {
			v.NewError("Todos los valores de la matriz deben ser del mismo tipo", ctx.GetStart())
			return nil
		}

		baseNode.Body = append(baseNode.Body, matrixDef)
	}

	return baseNode
}

func (v *Visitor) VisitMatrixRepeatingDefinitionNested(ctx *parser.MatrixRepeatingDefinitionNestedContext) interface{} {
	matrixNode, ok := v.Visit(ctx.MatrixRepeatingDefinition()).(*MatrixNode)

	if !ok {
		return nil
	}

	matrixType := v.Visit(ctx.MatrixType()).(string)

	if matrixNode.DataType != matrixType {
		v.NewError(fmt.Sprintf("Se esperaba un valor de tipo %s", matrixType), ctx.GetStart())
		return nil
	}

	repeatingTimes := v.Visit(ctx.Expr()).(V.IValue)

	// Check if the repeating times is an integer
	if repeatingTimes.GetType() != V.IntType {
		v.NewError("El número de repeticiones debe ser un entero", ctx.GetStart())
		return nil
	}

	if repeatingTimes.GetValue().(int) < 0 {
		v.NewError("El número de repeticiones debe ser mayor o igual a 0", ctx.GetStart())
		return nil
	}

	// Create a new matrix node
	body := make([]V.IValue, repeatingTimes.GetValue().(int))

	for i := 0; i < repeatingTimes.GetValue().(int); i++ {
		body[i] = matrixNode
	}

	return NewMatrixNode(matrixType, body)
}

func (v *Visitor) VisitMatrixRepeatingDefinitionSingle(ctx *parser.MatrixRepeatingDefinitionSingleContext) interface{} {
	// Get the value
	matrixType := v.Visit(ctx.MatrixType()).(string)

	repeatingValue := v.Visit(ctx.Expr(0)).(V.IValue)
	repeatingTimes := v.Visit(ctx.Expr(1)).(V.IValue)

	// Check if the repeating times is an integer
	if repeatingTimes.GetType() != V.IntType {
		v.NewError("El número de repeticiones debe ser un entero", ctx.GetStart())
		return nil
	}

	if repeatingTimes.GetValue().(int) < 0 {
		v.NewError("El número de repeticiones debe ser mayor o igual a 0", ctx.GetStart())
		return nil
	}

	// Check if the repeating value is the same type as the matrix type
	if repeatingValue.GetType() != matrixType {
		v.NewError(fmt.Sprintf("Se esperaba un valor de tipo %s, se obtuvo %s", matrixType, repeatingValue.GetType()), ctx.GetStart())
		return nil
	}

	// Create a new matrix node
	body := make([]V.IValue, repeatingTimes.GetValue().(int))

	for i := 0; i < repeatingTimes.GetValue().(int); i++ {
		body[i] = repeatingValue
	}

	return NewMatrixNode(matrixType, body)
}

// This function should return an array of n-dimensional arrays
func (v *Visitor) GetMatrixBody(ctx *parser.MatrixDeclarationContext) interface{} {
	// The body can be defined explicitly or implicitly
	// Explicitly: [[1,2,3],[4,5,6]]

	var body interface{}

	if ctx.MatrixDefinition() != nil {
		// Convert node to array
		body = v.Visit(ctx.MatrixDefinition())
	} else {
		body = v.Visit(ctx.MatrixRepeatingDefinition())
	}

	// Check if the body is a matrixNode
	matrixNode, ok := body.(*MatrixNode)

	if !ok {
		v.NewError(InvalidMatrixValue, ctx.GetStart())
		return nil
	}

	return matrixNode
}

func (v *Visitor) GetMatrixDimensions(node interface{}, counter int) int {
	// Recursive
	nd, ok := node.(*MatrixNode)

	if !ok {
		return counter
	}

	if len(nd.Body) == 0 {
		return counter
	}

	counter++

	return v.GetMatrixDimensions(nd.Body[0], counter)
}

func (v *Visitor) GetMatrixTypeDimensions(ctx string) int {
	var counter int
	for _, dim := range ctx {
		if dim == '[' {
			counter++
		}
	}

	return counter
}

func (v *Visitor) VisitMatrixAccess(ctx *parser.MatrixAccessContext) interface{} {
	ids := v.Visit(ctx.IdChain()).([]antlr.TerminalNode)
	id := ids[0].GetText()

	props := v.GetProps(ids)

	object, okV := v.LookUpObject(id, props, ctx.GetStart())

	if !okV {
		return nil
	}

	// Check if the object is a matrix
	if object.GetType() != V.MatrixType {
		v.NewError(fmt.Sprintf("El objeto %s no es una matriz", id), ctx.GetStart())
		return nil
	}

	// Get the all the indexes

	indexes := make([]V.IValue, 0)

	for _, index := range ctx.AllExpr() {
		index, ok := v.Visit(index).(V.IValue)

		if !ok || index.GetType() != V.IntType {
			v.NewError("El indice debe ser un entero", ctx.GetStart())
			return nil
		}

		if index.GetValue().(int) < 0 {
			v.NewError("El indice debe ser mayor o igual a 0", ctx.GetStart())
			return nil
		}

		indexes = append(indexes, index)
	}

	// Check if the indexes are in range

	value := v.CheckMatrixIndexes(object.GetValue().(*MatrixNode), indexes, ctx.GetStart())

	if value == nil {
		return nil
	}

	// Return matrix, indexes and value, then, handle when is called as expression, all this as dictionary

	dict := map[string]interface{}{
		"matrix":  object,
		"indexes": indexes,
		"value":   value,
	}

	return dict
}

func (v *Visitor) CheckMatrixIndexes(matrix interface{}, indexes []V.IValue, start antlr.Token) V.IValue {
	// Recursive function
	// Check if the matrix is a matrixNode or a value
	nd, ok := matrix.(*MatrixNode)

	if !ok {
		if len(indexes) == 0 {
			// Devolver el IValue si es un nodo hoja y no quedan más índices
			return matrix.(V.IValue)
		}
		v.NewError("No se puede acceder a un valor no matriz con índices", start)
		return nil
	}

	// Check if the indexes are in range
	if len(indexes) == 0 {
		return nd
	}

	if len(nd.Body) <= indexes[0].GetValue().(int) {
		v.NewError("El índice está fuera de rango", start)
		return nil
	}

	return v.CheckMatrixIndexes(nd.Body[indexes[0].GetValue().(int)], indexes[1:], start)
}

func (v *Visitor) VisitMatrixAssignment(ctx *parser.MatrixAssignmentContext) interface{} {
	// Get the matrix, then, get the indexes and the value
	values, ok := v.Visit(ctx.MatrixAccess()).(map[string]interface{})

	if !ok {
		v.NewError(InvalidMatrixValue, ctx.GetStart())
		return nil
	}

	matrix := values["matrix"].(*ObjectV)

	expr := v.Visit(ctx.Expr()).(V.IValue)

	// Check is not necessary, because the indexes are checked in the matrix access

	// Check if the value type is the same as the matrix type
	if expr.GetType() != matrix.GetValue().(*MatrixNode).DataType {
		v.NewError(fmt.Sprintf("El tipo de dato de la matriz debe ser %s", matrix.GetValue().(*MatrixNode).DataType), ctx.GetStart())
		return nil
	}

	indexes := values["indexes"].([]V.IValue)

	// Replace the value
	v.ReplaceMatrixValue(matrix.GetValue().(*MatrixNode), indexes, expr)

	return nil
}

func (v *Visitor) ReplaceMatrixValue(matrix *MatrixNode, indexes []V.IValue, value V.IValue) {
	// Check if childrens are matrixNodes or values

	if len(indexes) == 1 {
		matrix.Body[indexes[0].GetValue().(int)] = value
		return
	}

	// Recursive
	v.ReplaceMatrixValue(matrix.Body[indexes[0].GetValue().(int)].(*MatrixNode), indexes[1:], value)
}

// Structs

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {
	if v.Env.GetCurrentScope().ScopeType != RootEnv {
		v.NewError("Las estructuras solo pueden ser declaradas en el scope global", ctx.GetStart())
		return nil
	}

	id := ctx.ID().GetText()

	_, ok := v.Env.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("La estructura %s ya existe", id), ctx.GetStart())
		return nil
	}

	return nil
}

func (v *Visitor) VisitStructBody(ctx *parser.StructBodyContext) interface{} {
	return nil
}

func (v *Visitor) VisitStructProperty(ctx *parser.StructPropertyContext) interface{} {
	return nil
}
