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
		v.Visit(ctx.Block())
	}

	v.Scope.PopScope()

	return nil
}

func (v *Visitor) VisitIfElseStatement(ctx *parser.IfElseStatementContext) interface{} {
	expr, ok := v.Visit(ctx.Expr()).(IValue).GetValue().(bool)

	if !ok {
		v.NewError("Se esperaba una expresión booleana para la sentencia If", ctx.GetStart())
		return false
	}

	v.Scope.PushScope(IfScope)

	if expr {
		v.Visit(ctx.Block(0))
	} else {
		v.Visit(ctx.Block(1))
	}

	v.Scope.PopScope()

	return true
}

func (v *Visitor) VisitIfElseIfStatement(ctx *parser.IfElseIfStatementContext) interface{} {
	// Gramar looks like -> Kw_IF expr LBRACE block RBRACE Kw_ELSE ifStatement			# IfElseIfStatement;
	// So, we need to check if the else statement is another if statement

	expr, ok := v.Visit(ctx.Expr()).(IValue).GetValue().(bool)

	if !ok {
		v.NewError("Se esperaba una expresión booleana para la sentencia If", ctx.GetStart())
		return nil
	}

	v.Scope.PushScope(IfScope)

	if expr {
		v.Visit(ctx.Block())
	} else {
		v.Visit(ctx.IfStatement())
	}

	v.Scope.PopScope()

	return nil
}
