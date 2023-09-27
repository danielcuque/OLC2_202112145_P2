package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	response := c.Visit(ctx.Expr()).(*ValueResponse)

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
