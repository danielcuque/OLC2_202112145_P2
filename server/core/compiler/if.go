package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {

	endLabel := c.TAC.NewLabel("if")

	c.Env.PushEnv(IfEnv)

	for _, ifStmt := range ctx.AllIfTail() {

		ifStatement := ifStmt.(*parser.IfTailContext)

		currentLabel := c.TAC.NewLabel("if")
		nextLabel := c.TAC.NewLabel("if")

		condition := c.Visit(ifStatement.Expr()).(*ValueResponse).GetValue()

		c.TAC.AppendCode(
			[]string{
				fmt.Sprintf("if (%s == 1) goto %s;", condition, currentLabel),
				fmt.Sprintf("goto %s;", nextLabel),
				currentLabel.Declare(),
			},
			"Declaración de if",
		)

		c.Visit(ifStatement.Block())

		c.TAC.AppendCode(
			[]string{
				fmt.Sprintf("goto %s;", endLabel),
				nextLabel.Declare(),
			},
			"",
		)

	}

	if ctx.ElseStatement() != nil {
		c.ExecuteElse(ctx.ElseStatement().(*parser.ElseStatementContext), endLabel)
	}

	c.TAC.AppendCode(
		[]string{
			endLabel.Declare(),
		},
		"",
	)

	c.Env.PopEnv()
	return nil
}

func (c *Compiler) ExecuteIf(ctx *parser.IfTailContext, currentLabel, endLabel *Label) *Label {

	condition := c.Visit(ctx.Expr()).(*ValueResponse).GetValue()

	nextLabel := c.TAC.NewLabel("if")

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("if (%s == 1) goto %s;", condition, currentLabel),
			fmt.Sprintf("goto %s;", nextLabel),
			currentLabel.Declare(),
		},
		"Declaración de if",
	)

	c.Visit(ctx.Block())

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("goto %s;", endLabel),
		},
		"Declaración de if",
	)

	return nextLabel
}

func (c *Compiler) ExecuteElse(ctx *parser.ElseStatementContext, endLabel *Label) {
	c.Visit(ctx.Block())
}
