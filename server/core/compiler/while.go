package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {

	c.Env.PushEnv(WhileEnv)

	recurrence := c.NewLabelFlow("", []LabelFlowType{ContinueLabel})
	end := c.NewLabelFlow("", []LabelFlowType{BreakLabel})

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

	c.Visit(ctx.Block())

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("goto %s;", recurrence),
		},
		"",
	)

	c.TAC.AppendCode(
		[]string{
			end.Declare(),
		},
		"",
	)

	c.Env.PopEnv()

	return nil
}
