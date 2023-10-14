package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitVectorTypeValue(ctx *parser.VectorTypeValueContext) interface{} {
	c.Visit(ctx.VectorDefinition())
	return nil
}

func (c *Compiler) VisitVectorStructValue(ctx *parser.VectorStructValueContext) interface{} {

	return nil
}

func (c *Compiler) VisitVectorListValue(ctx *parser.VectorListValueContext) interface{} {
	fmt.Println("VisitVectorListValue")

	c.Visit(ctx.VectorValues())
	return nil
}

func (c *Compiler) VisitVectorSingleValue(ctx *parser.VectorSingleValueContext) interface{} {
	fmt.Println("VisitVectorSingleValue")
	c.Visit(ctx.Expr())
	return nil
}

func (c *Compiler) VisitVectorValues(ctx *parser.VectorValuesContext) interface{} {

	fmt.Println("VisitVectorValues")
	for _, value := range ctx.AllExpr() {
		c.Visit(value)
	}

	return nil
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
