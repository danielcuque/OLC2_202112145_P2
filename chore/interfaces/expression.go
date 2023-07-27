package interfaces

import (
	"strconv"
	"strings"

	"OLC2/chore/parser"
)

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
	case ">":
		if l > r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "<=":
		if l <= r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case ">=":
		if l >= r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "==":
		if l == r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "!=":
		if l != r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "&&":
		if l != 0 && r != 0 {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "||":
		if l != 0 || r != 0 {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "!":
		if r == 0 {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	}
	return Value{ParseValue: false}
}
