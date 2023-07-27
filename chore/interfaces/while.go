package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).ParseValue.(bool)
	for ok && value {
		v.Visit(ctx.Block())
		value, ok = v.Visit(ctx.Expr()).ParseValue.(bool)
	}
	return Value{ParseValue: true}
}
