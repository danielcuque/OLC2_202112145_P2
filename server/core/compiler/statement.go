package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.VariableDeclaration() != nil {
		return c.Visit(ctx.VariableDeclaration())
	}

	if ctx.VariableAssignment() != nil {
		return c.Visit(ctx.VariableAssignment())
	}

	if ctx.IfStatement() != nil {
		return c.Visit(ctx.IfStatement())
	}

	return nil
}
