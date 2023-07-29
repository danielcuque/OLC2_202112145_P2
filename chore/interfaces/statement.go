package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitStmt(ctx *parser.StatementContext) Value {
	if ctx.Variable_assignment() != nil {
		return v.Visit(ctx.Variable_assignment())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Whilestmt() != nil {
		return v.Visit(ctx.Whilestmt())
	}
	return Value{ParseValue: true}
}
