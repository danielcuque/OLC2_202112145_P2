package interfaces

import (
	"strconv"
	"strings"

	"OLC2/chore/parser"
)

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{ParseValue: i}
}

func (v *Visitor) VisitDoubleExpr(ctx *parser.DoubleExprContext) Value {
	f, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Value{ParseValue: f}
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

func (v *Visitor) VisitArithmeticExp(ctx *parser.ArithmeticExprContext) Value {
	l := v.Visit(ctx.GetLeft()).ParseValue.(int64)
	r := v.Visit(ctx.GetRight()).ParseValue.(int64)
	op := ctx.GetOp().GetText()

	operators := map[string]func(int64, int64) int64{
		"+": func(a, b int64) int64 { return a + b },
		"-": func(a, b int64) int64 { return a - b },
		"*": func(a, b int64) int64 { return a * b },
		"/": func(a, b int64) int64 { return a / b },
		"%": func(a, b int64) int64 { return a % b },
	}

	if f, ok := operators[op]; ok {
		return Value{ParseValue: f(l, r)}
	}

	return Value{ParseValue: false}
}

func (v *Visitor) VisitComparasionExp(ctx *parser.ComparasionExprContext) Value {
	// Parse left and right
	l := v.Visit(ctx.GetLeft()).ParseValue.(int64)
	r := v.Visit(ctx.GetRight()).ParseValue.(int64)
	op := ctx.GetOp().GetText()

	operators := map[string]func(int64, int64) bool{
		"<":  func(a, b int64) bool { return a < b },
		">":  func(a, b int64) bool { return a > b },
		"<=": func(a, b int64) bool { return a <= b },
		">=": func(a, b int64) bool { return a >= b },
		"==": func(a, b int64) bool { return a == b },
		"!=": func(a, b int64) bool { return a != b },
	}

	if f, ok := operators[op]; ok {
		return Value{ParseValue: f(l, r)}
	}

	return Value{ParseValue: false}
}

func (v *Visitor) VisitLogicalExp(ctx *parser.LogicalExprContext) Value {
	l := v.Visit(ctx.GetLeft()).ParseValue.(bool)
	r := v.Visit(ctx.GetRight()).ParseValue.(bool)
	op := ctx.GetOp().GetText()

	operators := map[string]func(bool, bool) bool{
		"&&": func(a, b bool) bool { return a && b },
		"||": func(a, b bool) bool { return a || b },
	}

	if f, ok := operators[op]; ok {
		return Value{ParseValue: f(l, r)}
	}

	return Value{ParseValue: false}
}

// Not operator
func (v *Visitor) VisitNotExp(ctx *parser.NotExprContext) Value {
	value := v.Visit(ctx.GetRight()).ParseValue.(bool)
	return Value{ParseValue: !value}
}
