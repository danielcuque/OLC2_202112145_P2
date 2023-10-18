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
		fmt.Sprintf("stack[%v] = %v;", value.GetAddress(), response.GetValue()),
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

	c.TAC.AppendInstruction(fmt.Sprintf("%v = H;", initVector), "Inicio del vector")

	for _, value := range ctx_.AllExpr() {
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

	vectorPosition := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v - %v;", vectorPosition, index.GetValue(), vectorObject.GetInit()),
			fmt.Sprintf("%v = %v + %v;", vectorPosition, vectorPosition, value.GetAddress()),
		},
		"Acceso a vector",
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
