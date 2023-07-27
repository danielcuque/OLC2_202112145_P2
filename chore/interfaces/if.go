package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).ParseValue.(bool)
	if ok && value {
		return v.Visit(ctx.Block())
	}
	return Value{ParseValue: false}
}
