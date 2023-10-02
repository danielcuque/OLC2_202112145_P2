package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitStatement(ctx *parser.StatementContext) interface{} {

	if ctx.FunctionCall() != nil {
		return c.Visit(ctx.FunctionCall())
	}

	if ctx.IfStatement() != nil {
		return c.Visit(ctx.IfStatement())
	}

	if ctx.VariableDeclaration() != nil {
		return c.Visit(ctx.VariableDeclaration())
	}

	if ctx.VariableAssignment() != nil {
		return c.Visit(ctx.VariableAssignment())
	}

	return nil
}
