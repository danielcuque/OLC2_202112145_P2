package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	// First save variable in tac

	id := ctx.ID().GetText()

	c.TAC.AppendNewTemporal(fmt.Sprintf("Declaracion de variable %s", id), NewTemporal())

	// Second, save value in tac
	// value := c.Visit(ctx.Expr())

	return nil
}
