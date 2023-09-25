package interfaces

import (
	E "OLC2/core/error"
	"OLC2/core/parser"
	V "OLC2/core/values"
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
	condition, ok := v.Visit(ctx.Expr()).(V.IValue)

	if !ok {
		v.NewError(E.InvalidExpression, ctx.GetStart())
		return false
	}

	if condition.GetType() != V.BooleanType {
		v.NewError("Se esperaba una expresión booleana", ctx.GetStart())
		return false
	}

	if condition.(*V.BooleanV).Value {

		v.Env.PushEnv(IfEnv)

		v.Visit(ctx.Block())

		v.Env.PopEnv()

		return true
	}

	return false
}

func (v *Visitor) VisitElseStatement(ctx *parser.ElseStatementContext) interface{} {

	v.Env.PushEnv(ElseEnv)
	v.Visit(ctx.Block())
	v.Env.PopEnv()

	return nil
}
