package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitGuardStatement(ctx *parser.GuardStatementContext) interface{} {
	expr := c.Visit(ctx.Expr()).(*ValueResponse)

	continueLabel := c.TAC.NewLabel("guard_exit")

	c.TAC.AppendInstruction(
		fmt.Sprintf(
			"if (%s == 1) goto %s;",
			expr.GetValue(),
			continueLabel,
		),
		"Sentencia Guard",
	)

	c.Visit(ctx.Block())

	c.TAC.AppendInstruction(continueLabel.Declare(), "")

	return nil
}
