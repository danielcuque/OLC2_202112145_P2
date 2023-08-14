package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.VariableAssignment() != nil {
		return v.Visit(ctx.VariableAssignment())
	}
	if ctx.VariableDeclaration() != nil {
		return v.Visit(ctx.VariableDeclaration())
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

	return nil
}
