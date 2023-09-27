package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	return c.DeclareValue(id, response)
}

func (c *Compiler) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	return c.DeclareValue(id, response)
}

func (c *Compiler) DeclareValue(id string, response *ValueResponse) *Value {

	if response.GetContextValue() == LiteralType {

		c.TAC.AppendCode(
			// tn = literal
			fmt.Sprintf("t%d = %s", c.TAC.TemporalQuantity(), response.GetValue()),
			"",
		)

		response.SetValue(fmt.Sprintf("t%d", c.TAC.TemporalQuantity()))
	}

	c.TAC.AppendCode(
		fmt.Sprintf("stack[(int)P] = %s", response.GetValue()),
		fmt.Sprintf("Declaración de la variable '%s'", id),
	)

	c.TAC.AppendCode(
		fmt.Sprintf("P = P + 1"),
		"",
	)

	newValue := NewValue(response.GetValue(), c.StackPointer.GetPointer(), TemporalCast(response.GetType()))

	c.StackPointer.AddPointer()

	c.Env.AddValue(id, newValue)

	return newValue
}
