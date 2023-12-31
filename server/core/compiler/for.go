package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	c.Env.Next()
	if _, ok := ctx.Expr().(*parser.RangeExprContext); ok {
		c.ForRange(ctx)
	}
	c.Env.PopEnv()
	return nil
}

func (c *Compiler) ForRange(ctx *parser.ForStatementContext) {

	id := ctx.ID().GetText()

	value := c.Env.GetValue(id)

	if value == nil {
		return
	}

	rangeExpr := c.Visit(ctx.Expr()).(map[string]interface{})

	iteratorTemp := c.TAC.NewTemporal(IntTemporal)
	loopLabel := c.TAC.NewLabel("ForLoop")
	endLabel := c.TAC.NewLabel("ForEnd")

	value.SetData(IntTemporal, iteratorTemp)

	// Assign the iterator to the first value of the range
	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf(
				"stack[(int)%s] = %s;",
				c.TAC.GetValueAddress(value),
				rangeExpr["left"],
			),

			loopLabel.Declare(),

			fmt.Sprintf(
				"%s = stack[(int)%s];",
				iteratorTemp,
				c.TAC.GetValueAddress(value),
			),

			fmt.Sprintf(
				"if (%s < %s) goto %s;",
				rangeExpr["right"],
				iteratorTemp,
				endLabel,
			),
		},
		"For range",
	)

	// Visit the block
	c.Visit(ctx.Block())

	// Increase the iterator
	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf(
				"%s = %s + 1;",
				iteratorTemp,
				iteratorTemp,
			),

			fmt.Sprintf(
				"stack[(int)%s] = %s;",
				c.TAC.GetValueAddress(value),
				iteratorTemp,
			),

			fmt.Sprintf("goto %s;", loopLabel),

			endLabel.Declare(),
		},
		"",
	)

}
