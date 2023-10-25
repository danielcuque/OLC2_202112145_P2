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

	response, metadata := c.CreateVectorValues(ctx.VectorDefinition().(*parser.VectorListValueContext))

	if response == nil {
		return nil
	}

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[(int)%v] = %v;", value.GetAddress(), response.GetValue()),
		"Direccion de vector",
	)

	newVectorObject := NewVector(response.GetValue().(*Temporal), metadata)
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

func (c *Compiler) CreateVectorValues(ctx *parser.VectorListValueContext) (*ValueResponse, []int) {

	vectorValues := make([]parser.IExprContext, 0)

	if ctx.VectorValues() != nil {
		vectorValues = c.Visit(ctx.VectorValues()).([]parser.IExprContext)
	}

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

	for _, value := range vectorValues {
		val := c.Visit(value).(*ValueResponse)
		c.TAC.AppendInstructions([]string{
			fmt.Sprintf("heap[(int)H] = %v;", val.GetValue()),
			"H = H + 1;",
		}, "")
	}

	isEmptyLabel := c.TAC.NewLabel("isEmpty")
	end := c.TAC.NewLabel("end")

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = H - %v;", counter, isEmptyAddress), // Obtenemos la cantidad de elementos del vector
			fmt.Sprintf("%v = %v - 1;", counter, counter),        // Restamos uno porque el contador empieza en 1
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

	return &ValueResponse{
			Value:       initVector,
			Type:        IntTemporal,
			ContextType: TemporalType,
		}, []int{
			0,
			len(vectorValues),
		}
}

func (c *Compiler) VisitVectorAccess(ctx *parser.VectorAccessContext) interface{} {
	id, _ := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))

	index := c.Visit(ctx.Expr()).(*ValueResponse)

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	newVectorObject := value.GetValue().(*Object).GetValue().(*Matrix)
	baseTemporal := c.TAC.NewTemporal(IntTemporal)

	vectorPosition := c.TAC.NewTemporal(IntTemporal)
	vectorCount := c.TAC.NewTemporal(IntTemporal)

	outOfBoundsLabel := c.TAC.NewLabel("outOfBounds")
	end := c.TAC.NewLabel("end")

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = stack[(int)%v];", baseTemporal, c.TAC.GetValueAddress(value)), // Obtenemos la posicion inicial del vector
			fmt.Sprintf("%v = heap[(int)%v];", vectorCount, baseTemporal),                   // Obtenemos la cantidad de elementos del vector
			fmt.Sprintf("if (%v > %v) goto %v;", index.GetValue(), vectorCount, outOfBoundsLabel),
			fmt.Sprintf("if (%v < 0) goto %v;", index.GetValue(), outOfBoundsLabel),
			fmt.Sprintf("%v = %v + 2;", baseTemporal, baseTemporal), // Obtenemos la posicion inicial del vector
			fmt.Sprintf("%v = %v - %v;", vectorPosition, index.GetValue(), newVectorObject.GetInit()),
			fmt.Sprintf("%v = %v + %v;", vectorPosition, vectorPosition, baseTemporal),
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
		fmt.Sprintf("Posicion de vector %v", id),
	)

	return &ValueResponse{
		Value:       vectorPosition,
		Type:        IntTemporal,
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
