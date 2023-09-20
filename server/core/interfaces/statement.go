package interfaces

import "OLC2/core/parser"

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.VariableAssignment() != nil {
		return v.Visit(ctx.VariableAssignment())
	}
	if ctx.VariableDeclaration() != nil {
		return v.Visit(ctx.VariableDeclaration())
	}

	if ctx.VariableAssignmentObject() != nil {
		return v.Visit(ctx.VariableAssignmentObject())
	}

	if ctx.IfStatement() != nil {
		return v.Visit(ctx.IfStatement())
	}

	if ctx.WhileStatement() != nil {
		return v.Visit(ctx.WhileStatement())
	}

	if ctx.SwitchStatement() != nil {
		return v.Visit(ctx.SwitchStatement())
	}

	if ctx.ForStatement() != nil {
		return v.Visit(ctx.ForStatement())
	}

	if ctx.ControlFlowStatement() != nil {
		return v.Visit(ctx.ControlFlowStatement())
	}

	if ctx.GuardStatement() != nil {
		return v.Visit(ctx.GuardStatement())
	}

	if ctx.FunctionCall() != nil {
		return v.Visit(ctx.FunctionCall())
	}

	if ctx.FunctionDeclarationStatement() != nil {
		return v.Visit(ctx.FunctionDeclarationStatement())
	}

	if ctx.VectorDeclaration() != nil {
		return v.Visit(ctx.VectorDeclaration())
	}

	if ctx.VectorAssignment() != nil {
		return v.Visit(ctx.VectorAssignment())
	}

	if ctx.MatrixDeclaration() != nil {
		return v.Visit(ctx.MatrixDeclaration())
	}

	if ctx.MatrixAssignment() != nil {
		return v.Visit(ctx.MatrixAssignment())
	}

	if ctx.StructDeclaration() != nil {
		return v.Visit(ctx.StructDeclaration())
	}

	return nil
}
