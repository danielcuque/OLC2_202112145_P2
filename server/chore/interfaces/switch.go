package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
)

var switchContextExpr V.IValue

func (v *Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	expr, ok := v.Visit(ctx.Expr()).(V.IValue)

	if !ok {
		v.NewError(InvalidExpressionError, ctx.GetStart())
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

	expr := v.Visit(ctx.Expr()).(V.IValue)

	if expr.GetType() != switchContextExpr.GetType() {
		v.NewError("El tipo de dato de la expresion del case no coincide con el tipo de dato de la expresion del switch", ctx.GetStart())
	}

	if expr.GetValue() == switchContextExpr.GetValue() {
		v.Env.PushEnv(SwitchEnv)
		v.Visit(ctx.Block())
		v.Env.PopEnv()
		return true
	}

	return false
}

func (v *Visitor) VisitSwitchDefault(ctx *parser.SwitchDefaultContext) interface{} {
	v.Env.PushEnv(SwitchEnv)
	v.Visit(ctx.Block())
	v.Env.PopEnv()
	return nil
}
