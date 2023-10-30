package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	id, props := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))

	expr := c.Visit(ctx.Expr()).(*ValueResponse)

	// 1. Get the value of the expression
	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	// Value have stack address
	newTemporal := c.TAC.NewTemporal(expr.GetType())
	// Can be three types of assignment, =, +=, -=

	op := ctx.GetOp().GetText()

	var responseValue string

	switch op {
	case "=":
		responseValue = fmt.Sprintf("%s = %s;", newTemporal, expr.GetValue())
	case "+=":
		responseValue = fmt.Sprintf("%s = %v + %s;", newTemporal, newTemporal, expr.GetValue())
	case "-=":
		responseValue = fmt.Sprintf("%s = %v - %v;", newTemporal, newTemporal, expr.GetValue())
	}

	if len(props) > 1 {
		baseTemporal := c.TAC.NewTemporal(value.GetType())

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("%s = stack[(int)%s];", baseTemporal, c.TAC.GetValueAddress(value)),
			},
			fmt.Sprintf("Acceso a la variable '%s'", id),
		)

		temporalPointer := c.GetProps(value, props[1:], baseTemporal, true)

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("heap[(int)%s] = %s;", temporalPointer, expr.GetValue()),
			},
			"",
		)

	} else {
		c.TAC.AppendInstructions(
			[]string{
				responseValue,
				fmt.Sprintf("stack[(int)%s] = %s;", c.TAC.GetValueAddress(value), newTemporal),
			},
			fmt.Sprintf("Asignación de la variable '%s'", id),
		)
	}

	return nil
}
