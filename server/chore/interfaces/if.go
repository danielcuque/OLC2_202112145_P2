package interfaces

import (
	"OLC2/chore/parser"
)

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	executeElse := true

	// Visit each if tail
	for _, ifStmt := range ctx.AllIfTail() {

		executeElse = !v.Visit(ifStmt).(bool)

		if !executeElse {
			break
		}
	}

	if executeElse && ctx.ElseStatement() != nil {
		v.Visit(ctx.ElseStatement())
	}

	return nil
}

func (v *Visitor) VisitIfTail(ctx *parser.IfTailContext) interface{} {
	condition, ok := v.Visit(ctx.Expr()).(IValue)

	if !ok {
		v.NewError("La expresión es inválida", ctx.GetStart())
		return false
	}

	if condition.GetType() != BOOLEAN_STR {
		v.NewError("Se esperaba una expresión booleana", ctx.GetStart())
		return false
	}

	if condition.(*BooleanV).Value {

		v.Scope.PushScope(IfScope)

		v.Visit(ctx.Block())

		v.Scope.PopScope()

		return true
	}

	return false
}

func (v *Visitor) VisitElseStatement(ctx *parser.ElseStatementContext) interface{} {

	v.Scope.PushScope(ElseScope)
	v.Visit(ctx.Block())
	v.Scope.PopScope()

	return nil
}
