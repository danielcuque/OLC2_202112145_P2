package interfaces

import (
	"reflect"
	"strconv"
	"strings"

	"OLC2/chore/parser"
)

var intT = reflect.TypeOf(int64(1)).String()
var floatT = reflect.TypeOf(float64(1)).String()
var stringT = reflect.TypeOf("").String()

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
		return Value{ParseValue: nil}
	}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{ParseValue: value}
}

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext) Value {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitUnaryMinusExp(ctx *parser.UnaryExprContext) Value {
	value := v.Visit(ctx.Expr()).ParseValue
	if reflect.TypeOf(value).String() == reflect.TypeOf(int64(1)).String() {
		return Value{ParseValue: -value.(int64)}
	}
	if reflect.TypeOf(value).String() == reflect.TypeOf(float64(1)).String() {
		return Value{ParseValue: -value.(float64)}
	}
	panic("Unknown type")
}

func (v *Visitor) VisitArithmeticExp(ctx *parser.ArithmeticExprContext) Value {
	l := v.Visit(ctx.GetLeft()).ParseValue
	r := v.Visit(ctx.GetRight()).ParseValue
	op := ctx.GetOp().GetText()
	return v.arithmeticOp(l, r, op)
}

func (v *Visitor) arithmeticOp(l, r interface{}, op string) Value {
	leftT := reflect.TypeOf(l).String()
	rightT := reflect.TypeOf(r).String()

	switch op {
	case "+":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) + r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) + r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) + float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) + r.(float64)}
		}
		if leftT == stringT && rightT == stringT {
			return Value{ParseValue: l.(string) + r.(string)}
		}
	case "-":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) - r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) - r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) - float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) - r.(float64)}
		}
	case "*":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) * r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) * r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) * float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) * r.(float64)}
		}
	case "/":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) / r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) / r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) / float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) / r.(float64)}
		}
	case "%":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) % r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: int64(l.(float64)) % int64(r.(float64))}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: int64(l.(float64)) % r.(int64)}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: l.(int64) % int64(r.(float64))}
		}
	}

	return Value{ParseValue: false}
}

func (v *Visitor) VisitComparasionExp(ctx *parser.ComparasionExprContext) Value {

	// Parse left and right
	l := v.Visit(ctx.GetLeft()).ParseValue
	r := v.Visit(ctx.GetRight()).ParseValue
	op := ctx.GetOp().GetText()

	// Get types
	leftT := reflect.TypeOf(l).String()
	rightT := reflect.TypeOf(r).String()

	// Compare
	switch op {
	case "==":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) == r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) == r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) == float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) == r.(float64)}
		}
		if leftT == stringT && rightT == stringT {
			return Value{ParseValue: l.(string) == r.(string)}
		}
	case "!=":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) != r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) != r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) != float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) != r.(float64)}
		}
		if leftT == stringT && rightT == stringT {
			return Value{ParseValue: l.(string) != r.(string)}
		}
	case ">":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) > r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) > r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) > float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) > r.(float64)}
		}
	case "<":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) < r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) < r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) < float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) < r.(float64)}
		}
	case ">=":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) >= r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) >= r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) >= float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) >= r.(float64)}
		}
	case "<=":
		if leftT == intT && rightT == intT {
			return Value{ParseValue: l.(int64) <= r.(int64)}
		}
		if leftT == floatT && rightT == floatT {
			return Value{ParseValue: l.(float64) <= r.(float64)}
		}
		if leftT == floatT && rightT == intT {
			return Value{ParseValue: l.(float64) <= float64(r.(int64))}
		}
		if leftT == intT && rightT == floatT {
			return Value{ParseValue: float64(l.(int64)) <= r.(float64)}
		}
	}

	return Value{ParseValue: false}
}

// Logical operators

func (v *Visitor) VisitTernaryExpr(ctx *parser.TernaryExprContext) Value {
	cond := v.Visit(ctx.GetCondition()).ParseValue.(bool)
	if cond {
		return v.Visit(ctx.GetCTrue())
	} else {
		return v.Visit(ctx.GetCFalse())
	}

}

func (v *Visitor) VisitLogicalExp(ctx *parser.LogicalExprContext) Value {
	l := v.Visit(ctx.GetLeft()).ParseValue.(bool)
	r := v.Visit(ctx.GetRight()).ParseValue.(bool)
	op := ctx.GetOp().GetText()

	// Terinary operator is a special case

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

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) Value {
	// Get the for loop variables
	return Value{ParseValue: false}
}
