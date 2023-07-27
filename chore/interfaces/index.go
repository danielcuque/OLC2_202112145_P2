package interfaces

import (
	"OLC2/chore/parser"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) Visit(tree antlr.ParseTree) Value {

	visitFuncs := map[reflect.Type]func(ctx antlr.ParseTree) Value{
		reflect.TypeOf((*parser.ProgramContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitProg(ctx.(*parser.ProgramContext))
		},
		reflect.TypeOf((*parser.BlockContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitBlock(ctx.(*parser.BlockContext))
		},
		reflect.TypeOf((*parser.StatementContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitStmt(ctx.(*parser.StatementContext))
		},
		reflect.TypeOf((*parser.AssignmentContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitAssignstmt(ctx.(*parser.AssignmentContext))
		},
		reflect.TypeOf((*parser.DeclarationContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitDecl(ctx.(*parser.DeclarationContext))
		},
		reflect.TypeOf((*parser.IfstmtContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIfstmt(ctx.(*parser.IfstmtContext))
		},
		reflect.TypeOf((*parser.WhilestmtContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitWhilestmt(ctx.(*parser.WhilestmtContext))
		},
		reflect.TypeOf((*parser.OpExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitOpExpr(ctx.(*parser.OpExprContext))
		},
		reflect.TypeOf((*parser.IntExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIntExpr(ctx.(*parser.IntExprContext))
		},
		reflect.TypeOf((*parser.IdExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIdExpr(ctx.(*parser.IdExprContext))
		},
		reflect.TypeOf((*parser.StrExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitStrExpr(ctx.(*parser.StrExprContext))
		},
		reflect.TypeOf((*parser.BoolExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitBoolExpr(ctx.(*parser.BoolExprContext))
		},
	}

	ctxType := reflect.TypeOf(tree)

	if fn, ok := visitFuncs[ctxType]; ok {
		return fn(tree)
	}

	// If the type is not found in the map we panic
	panic("Unknown context " + ctxType.String())
}

func (v *Visitor) VisitProg(ctx *parser.ProgramContext) Value {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value {
	for i := 0; ctx.Statement(i) != nil; i++ {
		v.Visit(ctx.Statement(i))
	}
	return Value{ParseValue: true}
}
