package compiler

import (
	"OLC2/core/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
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

	vectorObject := NewVector(response.GetValue().(*Temporal), metadata)

	value.SetData(MatrixTemporal, vectorObject)

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

	initVector := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstruction(fmt.Sprintf("%v = H;", initVector), "Inicio del vector")

	for _, value := range ctx.AllExpr() {
		val := c.Visit(value).(*ValueResponse)
		c.TAC.AppendInstructions([]string{
			fmt.Sprintf("heap[(int)H] = %v;", val.GetValue()),
			"H = H + 1;",
		}, "")
	}

	return &ValueResponse{
		Value:       initVector,
		Type:        IntTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) CreateVectorValues(ctx *parser.VectorListValueContext) (*ValueResponse, []int) {

	ctx_ := ctx.VectorValues().(*parser.VectorValuesContext)

	initVector := c.TAC.NewTemporal(IntTemporal)
	counter := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = H;", initVector), // Inicio del vector
			"H = H + 1;",                       // Aumentamos un espacio para dejar la posicion vacia donde va la propiedad .count
		},
		"Inicio del vector",
	)

	for _, value := range ctx_.AllExpr() {
		val := c.Visit(value).(*ValueResponse)
		c.TAC.AppendInstructions([]string{
			fmt.Sprintf("heap[(int)H] = %v;", val.GetValue()),
			"H = H + 1;",
		}, "")
	}

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = H - %v;", counter, initVector), // Obtenemos la cantidad de elementos del vector
			fmt.Sprintf("%v = %v - 1;", counter, counter),    // Restamos uno porque el contador empieza en 1
			fmt.Sprintf("heap[(int)%v] = %v;", initVector, counter),
		},
		"Cantidad de elementos del vector",
	)

	return &ValueResponse{
			Value:       initVector,
			Type:        IntTemporal,
			ContextType: TemporalType,
		}, []int{
			0,
			len(ctx_.AllExpr()),
		}
}

func (c *Compiler) VisitVectorAccess(ctx *parser.VectorAccessContext) interface{} {
	ids := c.Visit(ctx.IdChain()).([]antlr.TerminalNode)
	id := ids[0].GetText()

	index := c.Visit(ctx.Expr()).(*ValueResponse)

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	vectorObject := value.GetValue().(*Matrix)
	baseTemporal := c.TAC.NewTemporal(IntTemporal)

	vectorPosition := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = stack[(int)%v];", baseTemporal, value.GetAddress()),
			fmt.Sprintf("%v = %v + 1;", baseTemporal, baseTemporal),
			fmt.Sprintf("%v = %v - %v;", vectorPosition, index.GetValue(), vectorObject.GetInit()),
			fmt.Sprintf("%v = %v + %v;", vectorPosition, vectorPosition, baseTemporal),
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
