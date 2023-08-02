package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) Value {
	if ctx.VariableAssignment() != nil {
		return v.Visit(ctx.VariableAssignment())
	}
	if ctx.IfStatement() != nil {
		return v.Visit(ctx.IfStatement())
	}
	if ctx.WhileStatement() != nil {
		return v.Visit(ctx.WhileStatement())
	}
	if ctx.VariableDeclaration() != nil {
		return v.Visit(ctx.VariableDeclaration())
	}
	if ctx.ForStatement() != nil {
		return v.Visit(ctx.ForStatement())
	}
	return Value{ParseValue: true}
}
