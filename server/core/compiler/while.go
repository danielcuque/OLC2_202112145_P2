package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {

	response := c.Visit(ctx.Expr()).(*ValueResponse)

	fmt.Println(response)

	recurrence := c.TAC.NewLabel("")
	c.TAC.AppendCode(
		[]string{
			recurrence.Declare(),
		},
		"",
	)

	return nil
}
