package interfaces

import (
	"OLC2/chore/parser"
)

func (v *Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	condition, ok := v.Visit(ctx.Expr()).(IValue).GetValue().(bool)

	if !ok {
		v.NewError("La condición debe ser de tipo bool", ctx.GetStart())
		return nil
	}

	v.Scope.PushScope(WhileScope)

	for condition {
		v.Scope.ResetScope()
		v.Visit(ctx.Block())
		condition = v.Visit(ctx.Expr()).(IValue).GetValue().(bool)
	}

	v.Scope.PopScope()

	return nil
}
