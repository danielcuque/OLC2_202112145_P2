package interfaces

import (
	"OLC2/chore/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProg(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StatementContext:
		return v.VisitStmt(val)
	case *parser.AssignmentContext:
		return v.VisitAssignstmt(val)
	case *parser.IfstmtContext:
		return v.VisitIfstmt(val)
	case *parser.WhilestmtContext:
		return v.VisitWhilestmt(val)
	case *parser.OpExprContext:
		return v.VisitOpExpr(val)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	default:
		panic("Unknown context " + val.GetText())
	}

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

func (v *Visitor) VisitStmt(ctx *parser.StatementContext) Value {
	if ctx.Assignment() != nil {
		return v.Visit(ctx.Assignment())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Whilestmt() != nil {
		return v.Visit(ctx.Whilestmt())
	}
	return Value{ParseValue: true}
}

func (v *Visitor) VisitAssignstmt(ctx *parser.AssignmentContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.Memory[id] = value
	return Value{ParseValue: true}
}

func (v *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).ParseValue.(bool)
	if ok && value {
		return v.Visit(ctx.Block())
	}
	return Value{ParseValue: false}
}

func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).ParseValue.(bool)
	for ok && value {
		v.Visit(ctx.Block())
		value, ok = v.Visit(ctx.Expr()).ParseValue.(bool)
	}
	return Value{ParseValue: true}
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{ParseValue: i}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	value := strings.Trim(ctx.GetText(), "\"")
	return Value{ParseValue: value}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()
	value, ok := v.Memory[id]
	if ok {
		return value
	} else {
		panic("no such variable: " + id)
	}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{ParseValue: value}
}

func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext) Value {
	l := v.Visit(ctx.GetLeft()).ParseValue.(int64)
	r := v.Visit(ctx.GetRight()).ParseValue.(int64)
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		return Value{ParseValue: l + r}
	case "-":
		return Value{ParseValue: l - r}
	case "*":
		return Value{ParseValue: l * r}
	case "/":
		return Value{ParseValue: l / r}
	case "%":
		return Value{ParseValue: l % r}
	case "<":
		if l < r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	}
	return Value{ParseValue: false}
}
