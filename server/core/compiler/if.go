package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {

	endLabel := c.TAC.NewLabel("if")

	for _, ifStmt := range ctx.AllIfTail() {
		c.Env.PushEnv(IfEnv)
		ifStatement := ifStmt.(*parser.IfTailContext)

		currentLabel := c.TAC.NewLabel("if")
		nextLabel := c.TAC.NewLabel("if")

		condition := c.Visit(ifStatement.Expr()).(*ValueResponse).GetValue()

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("if (%s == 1) goto %s;", condition, currentLabel),
				fmt.Sprintf("goto %s;", nextLabel),
				currentLabel.Declare(),
			},
			"Declaración de if",
		)

		c.Visit(ifStatement.Block())

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("goto %s;", endLabel),
				nextLabel.Declare(),
			},
			"",
		)
		c.Env.PopEnv()
	}

	if ctx.ElseStatement() != nil {
		c.ExecuteElse(ctx.ElseStatement().(*parser.ElseStatementContext), endLabel)
	}

	c.TAC.AppendInstruction(endLabel.Declare(), "")

	return nil
}

func (c *Compiler) ExecuteElse(ctx *parser.ElseStatementContext, endLabel *Label) {
	c.Env.PushEnv(ElseEnv)
	c.Visit(ctx.Block())
	c.Env.PopEnv()
}
