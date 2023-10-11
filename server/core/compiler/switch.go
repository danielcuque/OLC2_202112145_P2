package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	expr := c.Visit(ctx.Expr()).(*ValueResponse)

	exitLabel := c.NewLabelFlow("", []LabelFlowType{BreakLabel})
	currentLabel := c.TAC.NewLabel("")

	c.TAC.AppendInstruction(currentLabel.Declare(), "Sentencia Switch")

	for _, switchCase := range ctx.AllSwitchCase() {

		swCase := switchCase.(*parser.SwitchCaseContext)

		nextLabel := c.TAC.NewLabel("")

		c.Env.PushEnv(SwitchEnv)

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf(
					"if(%v != %v) goto %s;",
					expr.GetValue(),
					c.Visit(swCase.Expr()).(*ValueResponse).GetValue(),
					nextLabel,
				),
			},
			"",
		)

		c.Visit(swCase.Block())

		c.TAC.AppendInstruction(fmt.Sprintf("goto %s;", exitLabel), "")

		currentLabel = nextLabel
		c.TAC.AppendInstruction(currentLabel.Declare(), "")

		c.Env.PopEnv()
	}

	if ctx.SwitchDefault() != nil {
		c.Env.PushEnv(SwitchEnv)
		c.Visit(ctx.SwitchDefault().Block())
		c.Env.PopEnv()
	}

	c.TAC.AppendInstruction(exitLabel.Declare(), "")

	return nil
}
