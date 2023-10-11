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
	c.DeclareControlFlow(BreakLabel)
	return nil
}

func (c *Compiler) DeclareControlFlow(LabelFlowType LabelFlowType) {
	label := c.Env.GetLabel(BreakLabel)

	if label == nil {
		return
	}

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("goto %s;", label),
		},
		fmt.Sprintf("Control %s", LabelFlowType),
	)
}
