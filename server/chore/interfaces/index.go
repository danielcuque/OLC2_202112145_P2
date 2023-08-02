package interfaces

import (
	"OLC2/chore/parser"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) Visit(tree antlr.ParseTree) Value {

	switch val := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProg(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StatementContext:
		return v.VisitStatement(val)
	case *parser.VariableDeclarationContext:
		return v.VisitVariableDeclaration(val)
	case *parser.VariableAssignmentContext:
		return v.VisitVariableAssignment(val)
	case *parser.IfStatementContext:
		return v.VisitIfStatement(val)
	case *parser.WhileStatementContext:
		return v.VisitWhileStatement(val)
	case *parser.ForStatementContext:
		return v.VisitForStatement(val)

	// Expressions
	case *parser.ArithmeticExprContext:
		return v.VisitArithmeticExp(val)
	case *parser.LogicalExprContext:
		return v.VisitLogicalExp(val)
	case *parser.ComparasionExprContext:
		return v.VisitComparasionExp(val)
	case *parser.ParExprContext:
		return v.VisitParExpr(val)
	case *parser.NotExprContext:
		return v.VisitNotExp(val)
	case *parser.UnaryExprContext:
		return v.VisitUnaryMinusExp(val)
	case *parser.TernaryExprContext:
		return v.VisitTernaryExpr(val)
	case *parser.RangeExprContext:
		return v.VisitRangeExpr(val)

	// Types
	case *parser.IntExprContext:
		return v.VisitIntExpr(val)
	case *parser.DoubleExprContext:
		return v.VisitDoubleExpr(val)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	}

	panic("Unknown type" + reflect.TypeOf(tree).String())
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
