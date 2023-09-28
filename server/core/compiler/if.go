package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	// Change this code to generate TAC for if statements

	executeElse := true

	for _, ifStmt := range ctx.AllIfTail() {

		executeElse = !c.Visit(ifStmt).(bool)

		if !executeElse {
			break
		}
	}

	if executeElse && ctx.ElseStatement() != nil {
		c.Visit(ctx.ElseStatement())
	}

	return nil
}

func (c *Compiler) VisitIfTail(ctx *parser.IfTailContext) interface{} {
	condition := c.Visit(ctx.Expr()).(*ValueResponse).GetValue()

	if condition == "1" {

		c.Env.PushEnv(IfEnv)

		c.Visit(ctx.Block())

		c.Env.PopEnv()

		return true
	}

	return false
}

func (c *Compiler) VisitElseStatement(ctx *parser.ElseStatementContext) interface{} {

	c.Env.PushEnv(ElseEnv)
	c.Visit(ctx.Block())
	c.Env.PopEnv()

	return nil
}
