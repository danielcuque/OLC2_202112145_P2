package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitControlBreak(ctx *parser.ControlBreakContext) interface{} {
	c.DeclareControlFlow(BreakLabel)
	return nil
}

func (c *Compiler) VisitControlContinue(ctx *parser.ControlContinueContext) interface{} {
	c.DeclareControlFlow(ContinueLabel)
	return nil
}

func (c *Compiler) VisitControlReturn(ctx *parser.ControlReturnContext) interface{} {

	if ctx.Expr() != nil {
		value := c.Visit(ctx.Expr()).(*ValueResponse)

		procedure := c.TAC.GetCurrentProcedure()

		if procedure == nil {
			return nil
		}

		c.TAC.AppendInstruction(
			fmt.Sprintf(
				"%s = %s;",
				procedure.ReturnTemporal,
				value.GetValue(),
			),
			"Valor de retorno",
		)
	}

	c.DeclareControlFlow(ReturnLabel)
	return nil
}

func (c *Compiler) DeclareControlFlow(LabelFlowType LabelFlowType) {
	label := c.Env.GetLabel(LabelFlowType)

	if label == nil {
		return
	}

	c.TAC.AppendInstruction(fmt.Sprintf("goto %s;", label), fmt.Sprintf("Control %s", LabelFlowType))
}
