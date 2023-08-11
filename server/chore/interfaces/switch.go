package interfaces

import (
	"OLC2/chore/parser"
)

var switchContextExpr IValue

func (v *Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	expr, ok := v.Visit(ctx.Expr()).(IValue)

	if !ok {
		v.NewError("La expresion del switch no es de tipo IValue", ctx.GetStart())
		return nil
	}

	executeDefault := true

	switchContextExpr = expr

	for _, switchCase := range ctx.AllSwitchCase() {
		executeDefault = !v.Visit(switchCase).(bool)

		if !executeDefault {
			break
		}
	}

	if executeDefault && ctx.SwitchDefault() != nil {
		v.Visit(ctx.SwitchDefault())
	}

	return nil
}

func (v *Visitor) VisitSwitchCase(ctx *parser.SwitchCaseContext) interface{} {

	expr := v.Visit(ctx.Expr()).(IValue)

	if expr.GetType() != switchContextExpr.GetType() {
		v.NewError("El tipo de dato de la expresion del case no coincide con el tipo de dato de la expresion del switch", ctx.GetStart())
	}

	if expr.GetValue() == switchContextExpr.GetValue() {
		v.Scope.PushScope(SwitchScope)
		v.Visit(ctx.Block())
		v.Scope.PopScope()
		return true
	}

	return false
}

func (v *Visitor) VisitSwitchDefault(ctx *parser.SwitchDefaultContext) interface{} {
	v.Scope.PushScope(SwitchScope)
	v.Visit(ctx.Block())
	v.Scope.PopScope()
	return nil
}
