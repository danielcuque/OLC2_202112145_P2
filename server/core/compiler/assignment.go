package compiler

import (
	"OLC2/core/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	ids := c.Visit(ctx.IdChain()).([]antlr.TerminalNode)
	id := ids[0].GetText()

	response := c.Visit(ctx.Expr()).(*ValueResponse)

	// 1. Get the value of the expression
	value := c.Env.GetValue(id)

	// Value have stack address
	c.TAC.NewTemporal(response.GetType())
	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("t%d = %s;", c.TAC.TemporalQuantity(), response.GetValue()),
			fmt.Sprintf("stack[(int)%d] = t%d;", value.GetAddress(), c.TAC.TemporalQuantity()),
		},
		fmt.Sprintf("Asignación de la variable '%s'", id),
	)

	return nil
}

// func (c *Compiler) handleVariableAssignment(ctx antlr.ParserRuleContext, variable *Variable, value V.IValue, op string) interface{} {

// 	if variable.GetType() != value.GetType() {
// 		if variable.GetType() == V.FloatType && value.GetType() == V.IntType {
// 			value = V.NewFloatValue(float64(value.GetValue().(int)))
// 		}
// 	}

// 	var newValue *ValueResponse
// 	switch op {
// 	case "=":
// 		newValue = value.Copy()
// 	case "+=":
// 		newValue = c.arithmeticOp(variable.Value, value, "+", ctx.GetStart()).(*ValueResponse)
// 	case "-=":
// 		newValue = c.arithmeticOp(variable.Value, value, "-", ctx.GetStart()).(*ValueResponse)
// 	default:
// 	}

// 	variable.SetValue(newValue)
// 	return nil
// }
