package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitStmt(ctx *parser.StatementContext) Value {
	if ctx.VariableAssignment() != nil {
		return v.Visit(ctx.VariableAssignment())
	}
	if ctx.IfStatement() != nil {
		return v.Visit(ctx.IfStatement())
	}
	if ctx.WhiteStatement() != nil {
		return v.Visit(ctx.WhiteStatement())
	}
	if ctx.VariableDeclaration() != nil {
		return v.Visit(ctx.VariableDeclaration())
	}
	if ctx.ForStatement() != nil {
		return v.Visit(ctx.ForStatement())
	}
	return Value{ParseValue: true}
}
