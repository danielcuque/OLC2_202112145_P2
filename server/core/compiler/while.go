package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {

	recurrence := c.TAC.NewLabel("")
	end := c.TAC.NewLabel("")

	c.TAC.AppendCode(
		[]string{
			recurrence.Declare(),
		},
		"",
	)

	condition := c.Visit(ctx.Expr()).(*ValueResponse)

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf(
				"if (%s == 0) goto %s;",
				condition.GetValue(),
				end,
			),
		},
		"",
	)

	c.Context.Push(end)

	c.Visit(ctx.Block())

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("goto %s;", recurrence),
		},
		"",
	)

	c.Context.Pop()

	c.TAC.AppendCode(
		[]string{
			end.Declare(),
		},
		"",
	)

	return nil
}
