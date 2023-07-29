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
		reflect.TypeOf((*parser.VariableAssignmentContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitVariableAssign(ctx.(*parser.VariableAssignmentContext))
		},
		reflect.TypeOf((*parser.VariableDeclarationContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitVariableDeclaration(ctx.(*parser.VariableDeclarationContext))
		},
		reflect.TypeOf((*parser.IfStatementContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIfstmt(ctx.(*parser.IfStatementContext))
		},
		reflect.TypeOf((*parser.WhiteStatementContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitWhilestmt(ctx.(*parser.WhiteStatementContext))
		},
		reflect.TypeOf((*parser.ForStatementContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitForStatement(ctx.(*parser.ForStatementContext))
		},

		// Expressions
		reflect.TypeOf((*parser.ArithmeticExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitArithmeticExp(ctx.(*parser.ArithmeticExprContext))
		},
		reflect.TypeOf((*parser.LogicalExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitLogicalExp(ctx.(*parser.LogicalExprContext))
		},
		reflect.TypeOf((*parser.ComparasionExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitComparasionExp(ctx.(*parser.ComparasionExprContext))
		},
		reflect.TypeOf((*parser.ParExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitParExpr(ctx.(*parser.ParExprContext))
		},

		reflect.TypeOf((*parser.NotExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitNotExp(ctx.(*parser.NotExprContext))
		},
		reflect.TypeOf((*parser.UnaryExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitUnaryMinusExp(ctx.(*parser.UnaryExprContext))
		},
		reflect.TypeOf((*parser.TernaryExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitTernaryExpr(ctx.(*parser.TernaryExprContext))
		},

		// Types
		reflect.TypeOf((*parser.IntExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIntExpr(ctx.(*parser.IntExprContext))
		},
		reflect.TypeOf((*parser.DoubleExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitDoubleExpr(ctx.(*parser.DoubleExprContext))
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

	// fmt.Println(ctxType)

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
