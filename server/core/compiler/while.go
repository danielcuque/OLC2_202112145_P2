package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {

	c.Env.PushEnv(WhileEnv)

	recurrence := c.NewLabelFlow("", []LabelFlowType{ContinueLabel})
	end := c.NewLabelFlow("", []LabelFlowType{BreakLabel})

	c.TAC.AppendInstruction(recurrence.Declare(), "")

	condition := c.Visit(ctx.Expr()).(*ValueResponse)

	c.TAC.AppendInstructions(
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

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("goto %s;", recurrence),
			end.Declare(),
		},
		"",
	)

	c.Env.PopEnv()

	return nil
}
