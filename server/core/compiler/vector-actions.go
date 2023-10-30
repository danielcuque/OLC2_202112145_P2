package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitVectorTypeValue(ctx *parser.VectorTypeValueContext) interface{} {
	id := ctx.ID().GetText()

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	valueType := c.Visit(ctx.VariableType()).(TemporalCast)

	values := c.GetVectorValues(ctx.VectorDefinition().(*parser.VectorListValueContext))

	initVector := c.InitNewMatrix() // Declaramos un nuevo vector

	c.DeclareVectorValues(values, initVector[1]) // Declaramos su contenido

	c.DeclareMatrixProps(initVector[0], initVector[1], initVector[2]) // Definimos las propiedades del vector

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[(int)%v] = %v;", value.GetAddress(), initVector[0]), // Inicio del vector
		"Direccion de vector",
	)

	newVectorObject := NewVector(valueType)

	value.SetData(MatrixTemporal, newVectorObject)

	return nil
}

func (c *Compiler) VisitVectorStructValue(ctx *parser.VectorStructValueContext) interface{} {
	return nil
}

func (c *Compiler) VisitVectorListValue(ctx *parser.VectorListValueContext) interface{} {
	return c.Visit(ctx.VectorValues())
}

func (c *Compiler) VisitVectorSingleValue(ctx *parser.VectorSingleValueContext) interface{} {
	return c.Visit(ctx.Expr())
}

func (c *Compiler) VisitVectorValues(ctx *parser.VectorValuesContext) interface{} {
	return ctx.AllExpr()
}

func (c *Compiler) GetVectorValues(ctx *parser.VectorListValueContext) []*ValueResponse {

	vectorValues := make([]parser.IExprContext, 0)
	valueResponses := make([]*ValueResponse, 0)

	if ctx.VectorValues() != nil {
		vectorValues = c.Visit(ctx.VectorValues()).([]parser.IExprContext)
	}

	for _, value := range vectorValues {
		val := c.Visit(value).(*ValueResponse)

		if value == nil {
			continue
		}
		valueResponses = append(valueResponses, val)

	}

	return valueResponses
}

func (c *Compiler) DeclareVectorValues(values []*ValueResponse, counter *Temporal) {

	for _, value := range values {
		c.AppendVectorValue(value.GetValue())
	}

	c.TAC.AppendInstruction(
		fmt.Sprintf("%v = %v;", counter, len(values)),
		"Contador de vector",
	)
}

func (c *Compiler) AppendVectorValue(value interface{}) {
	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("heap[(int)H] = %v;", value),
			c.HeapPointer.IncreasePointer(),
		},
		"",
	)
}

func (c *Compiler) InitNewMatrix() []*Temporal {

	initVector := c.TAC.NewTemporal(IntTemporal)
	counter := c.TAC.NewTemporal(IntTemporal)
	isEmptyAddress := c.TAC.NewTemporal(BooleanTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = H;", initVector),     // Inicio del vector
			c.HeapPointer.IncreasePointer(),        // Aumentamos un espacio para dejar la posicion vacia donde va la propiedad .count
			fmt.Sprintf("%v = H;", isEmptyAddress), // Inicio del contador
			c.HeapPointer.IncreasePointer(),        // Aumentamos un espacio para dejar la posicion vacia donde va la propiedad .isEmpty
		},
		"Inicio del vector",
	)

	return []*Temporal{
		initVector,
		counter,
		isEmptyAddress,
	}
}

func (c *Compiler) DeclareNewMatrix(matrix *InitMatrix) {
	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = H;", matrix.InitVector),     // Inicio del vector
			c.HeapPointer.IncreasePointer(),               // Aumentamos un espacio para dejar la posicion vacia donde va la propiedad .count
			fmt.Sprintf("%v = H;", matrix.IsEmptyAddress), // Inicio del contador
			c.HeapPointer.IncreasePointer(),               // Aumentamos un espacio para dejar la posicion vacia donde va la propiedad .isEmpty
		},
		"Inicio del vector",
	)
}

func (c *Compiler) DeclareMatrixProps(initVector, counter, isEmptyAddress *Temporal) {
	isEmptyLabel := c.TAC.NewLabel("isEmpty")
	end := c.TAC.NewLabel("end")

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("heap[(int)%v] = %v;", initVector, counter),
			fmt.Sprintf("if (%v == 0) goto %v;", counter, isEmptyLabel),
			fmt.Sprintf("heap[(int)%v] = 0;", isEmptyAddress),
			fmt.Sprintf("goto %v;", end),

			isEmptyLabel.Declare(),

			fmt.Sprintf("heap[(int)%v] = 1;", isEmptyAddress),

			end.Declare(),
		},
		"Propiedades del vector",
	)
}

func (c *Compiler) VisitVectorAccess(ctx *parser.VectorAccessContext) interface{} {
	id, _ := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))

	index := c.Visit(ctx.Expr()).(*ValueResponse)

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	genericObject := value.GetValue().(*Object)
	vectorObject := genericObject.GetValue().(*Matrix)

	return c.VectorAccess(c.TAC.GetValueAddress(value), index, true, vectorObject.GetType())
}

func (c *Compiler) VectorAccess(initAddress string, index *ValueResponse, isStackValue bool, vectorType TemporalCast) *ValueResponse {
	vectorAddress := c.TAC.NewTemporal(IntTemporal)
	vectorAccessPosition := c.TAC.NewTemporal(IntTemporal)
	vectorCountValue := c.TAC.NewTemporal(IntTemporal)

	outOfBoundsLabel := c.TAC.NewLabel("outOfBounds")
	end := c.TAC.NewLabel("end")

	initialInstructions := ""

	if isStackValue {
		initialInstructions += fmt.Sprintf("%v = stack[(int)%s];\n", vectorAddress, initAddress)
		initialInstructions += fmt.Sprintf("%v = heap[(int)%v];\n", vectorCountValue, vectorAddress)
	} else {
		initialInstructions += fmt.Sprintf("%v = heap[(int)%s];\n", vectorAddress, initAddress)
		initialInstructions += fmt.Sprintf("%v = %v\n;", vectorCountValue, vectorAddress)
	}

	initialInstructions += fmt.Sprintf("if (%v > %v) goto %v;\n", index.GetValue(), vectorCountValue, outOfBoundsLabel)
	initialInstructions += fmt.Sprintf("if (%v < 0) goto %v;\n", index.GetValue(), outOfBoundsLabel)

	if isStackValue {
		initialInstructions += fmt.Sprintf("%v = %v + 2;\n", vectorAddress, vectorAddress) // Obtenemos la posicion inicial del vector
	} else {
		initialInstructions += fmt.Sprintf("%v = %v + 2;", vectorAddress, initAddress) // Obtenemos la posicion inicial del vector
	}

	c.TAC.AppendInstructions(
		[]string{
			initialInstructions,
			// Obtenemos la posicion del valor del vector
			fmt.Sprintf("%v = %v - %v;", vectorAccessPosition, index.GetValue(), 0),
			fmt.Sprintf("%v = %v + %v;", vectorAccessPosition, vectorAccessPosition, vectorAddress),

			fmt.Sprintf("goto %v;", end),

			outOfBoundsLabel.Declare(),

			"printf(\"%c\", 79);",  //O
			"printf(\"%c\", 117);", //u
			"printf(\"%c\", 116);", //t
			"printf(\"%c\", 32);",  //
			"printf(\"%c\", 111);", //o
			"printf(\"%c\", 102);", //f
			"printf(\"%c\", 32);",  //
			"printf(\"%c\", 66);",  //B
			"printf(\"%c\", 111);", //o
			"printf(\"%c\", 117);", //u
			"printf(\"%c\", 110);", //n
			"printf(\"%c\", 100);", //d
			"printf(\"%c\", 115);", //s
			"printf(\"%c\", 10);",  //

			end.Declare(),
		},
		"Posicion de vector",
	)

	return &ValueResponse{
		Value:       vectorAccessPosition,
		Type:        vectorType,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitVectorAssignment(ctx *parser.VectorAssignmentContext) interface{} {
	response := c.Visit(ctx.VectorAccess()).(*ValueResponse)
	vectorPosition := response.GetValue()
	value := c.Visit(ctx.Expr()).(*ValueResponse)

	c.TAC.AppendInstruction(
		fmt.Sprintf("heap[(int)%v] = %v;", vectorPosition, value.GetValue()),
		"Asignacion a vector",
	)

	return nil
}
