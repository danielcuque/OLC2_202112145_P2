package interfaces

import (
	"OLC2/chore/parser"
)

func (v *Visitor) VisitSimpleIfStatement(ctx *parser.SimpleIfStatementContext) interface{} {
	expr, ok := v.Visit(ctx.Expr()).(IValue).GetValue().(bool)

	if !ok {
		v.NewError("Se esperaba una expresión booleana para la sentencia If", ctx.GetStart())
		return nil
	}

	v.Scope.PushScope(IfScope)

	if expr {
		return v.Visit(ctx.Block())
	}

	v.Scope.PopScope()

	return nil
}

func (v *Visitor) VisitIfElseStatement(ctx *parser.IfElseStatementContext) interface{} {
	return nil
}

func (v *Visitor) VisitIfElseIfStatement(ctx *parser.IfElseIfStatementContext) interface{} {
	return nil
}
