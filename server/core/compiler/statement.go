package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitStatement(ctx *parser.StatementContext) interface{} {

	if ctx.ControlFlowStatement() != nil {
		return c.Visit(ctx.ControlFlowStatement())
	}

	if ctx.ForStatement() != nil {
		return c.Visit(ctx.ForStatement())
	}

	if ctx.FunctionCall() != nil {
		return c.Visit(ctx.FunctionCall())
	}

	if ctx.FunctionDeclarationStatement() != nil {
		return c.Visit(ctx.FunctionDeclarationStatement())
	}

	if ctx.GuardStatement() != nil {
		return c.Visit(ctx.GuardStatement())
	}

	if ctx.IfStatement() != nil {
		return c.Visit(ctx.IfStatement())
	}

	if ctx.MatrixAssignment() != nil {
		return c.Visit(ctx.MatrixAssignment())
	}

	if ctx.MatrixDeclaration() != nil {
		return c.Visit(ctx.MatrixDeclaration())
	}

	if ctx.StructDeclaration() != nil {
		return c.Visit(ctx.StructDeclaration())
	}

	if ctx.SwitchStatement() != nil {
		return c.Visit(ctx.SwitchStatement())
	}

	if ctx.VariableAssignment() != nil {
		return c.Visit(ctx.VariableAssignment())
	}

	if ctx.VariableDeclaration() != nil {
		return c.Visit(ctx.VariableDeclaration())
	}

	if ctx.VectorAssignment() != nil {
		return c.Visit(ctx.VectorAssignment())
	}

	if ctx.VectorDeclaration() != nil {
		return c.Visit(ctx.VectorDeclaration())
	}

	if ctx.WhileStatement() != nil {
		return c.Visit(ctx.WhileStatement())
	}

	return nil
}
