package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	value := c.Visit(ctx.Expr()).(*ValueResponse)

	if value == nil {
		return nil
	}

	c.TAC.AppendCode(
		fmt.Sprintf("stack[(int)P] = %s", value.GetValue()),
		fmt.Sprintf("Declaración de %s", id),
	)

	c.TAC.AppendCode(
		fmt.Sprintf("P = P + 1"),
		"",
	)

	return nil
}
