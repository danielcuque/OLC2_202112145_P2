package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
)

func (v *Visitor) VisitGuardStatement(ctx *parser.GuardStatementContext) interface{} {
	allStatements := ctx.Block().AllStatement()
	allStatementsLen := len(allStatements)
	lastStatement := allStatements[allStatementsLen-1]

	if lastStatement.ControlFlowStatement() == nil {
		v.NewError("La sentencia guard debe contener una instrucción de control de flujo", ctx.GetStart())
		return nil
	}

	expr, ok := v.Visit(ctx.Expr()).(V.IValue).GetValue().(bool)

	if !ok {
		v.NewError("Se esperaba una expresión booleana", ctx.GetStart())
		return nil
	}

	if expr {
		v.Visit(ctx.Block())
		return nil
	}

	return nil
}
