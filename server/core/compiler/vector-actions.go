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

	response := c.Visit(ctx.VectorDefinition()).(*ValueResponse)

	if response == nil {
		return nil
	}

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[%v] = %v;", value.GetAddress(), response.GetValue()),
		"Direccion de vector",
	)

	value.SetData(MatrixTemporal, response.GetValue())

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

func (c *Compiler) VisitVectorAccess(ctx *parser.VectorAccessContext) interface{} {
	fmt.Println("VisitVectorAccess")
	c.Visit(ctx.Expr())
	return nil
}

func (c *Compiler) VisitVectorAssignment(ctx *parser.VectorAssignmentContext) interface{} {
	fmt.Println("VisitVectorAssignment")
	c.Visit(ctx.VectorAccess())
	c.Visit(ctx.Expr())
	return nil
}
