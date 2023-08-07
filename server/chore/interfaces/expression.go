package interfaces

import (
	"strconv"
	"strings"

	"OLC2/chore/parser"
	U "OLC2/chore/utils"
)

var intT = U.GetType(int64(1))
var floatT = U.GetType(float64(1))
var stringT = U.GetType("")

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{ParseValue: i}
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	f, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Value{ParseValue: f}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	return Value{ParseValue: value}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText()
	value, ok := v.Memory[id]
	if ok {
		return value
	} else {
		return Value{ParseValue: nil}
	}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{ParseValue: value}
}

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	value := v.Visit(ctx.Expr()).(Value).ParseValue
	valueT := U.GetType(value)
	if valueT == intT {
		return Value{ParseValue: -value.(int64)}
	} else if valueT == floatT {
		return Value{ParseValue: -value.(float64)}
	}
	v.NewError("Error: No se puede aplicar el operador unario - a " + U.GetType(value))
	return Value{ParseValue: nil}
}

func (v *Visitor) VisitArithmeticExpr(ctx *parser.ArithmeticExprContext) interface{} {
	l := v.Visit(ctx.GetLeft()).(Value).ParseValue
	r := v.Visit(ctx.GetRight()).(Value).ParseValue
	op := ctx.GetOp().GetText()
	return v.arithmeticOp(l, r, op)
}

func (v *Visitor) arithmeticOp(l, r interface{}, op string) Value {
	if l == nil || r == nil {
		return Value{ParseValue: nil}
	}

	leftT := U.GetType(l)
	rightT := U.GetType(r)

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
		v.NewError("Error: No se puede sumar " + leftT + " con " + rightT)
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
		v.NewError("Error: No se puede restar " + leftT + " con " + rightT)
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
		v.NewError("Error: No se puede multiplicar " + leftT + " con " + rightT)
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
		v.NewError("Error: No se puede dividir " + leftT + " con " + rightT)
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
		v.NewError("Error: No se puede modular " + leftT + " con " + rightT)
	}

	return Value{ParseValue: nil}
}

func (v *Visitor) VisitComparasionExpr(ctx *parser.ComparasionExprContext) interface{} {

	// Parse left and right
	l := v.Visit(ctx.GetLeft()).(Value).ParseValue
	r := v.Visit(ctx.GetRight()).(Value).ParseValue
	op := ctx.GetOp().GetText()

	// Get types
	leftT := U.GetType(l)
	rightT := U.GetType(r)

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

	v.NewError("Error: No se puede comparar " + leftT + " con " + rightT)
	return Value{ParseValue: false}
}

// Logical operators

func (v *Visitor) VisitTernaryExpr(ctx *parser.TernaryExprContext) interface{} {
	cond := v.Visit(ctx.GetCondition()).(Value).ParseValue.(bool)
	if cond {
		return v.Visit(ctx.GetCTrue())
	} else {
		return v.Visit(ctx.GetCFalse())
	}

}

func (v *Visitor) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	l := v.Visit(ctx.GetLeft()).(Value).ParseValue.(bool)
	r := v.Visit(ctx.GetRight()).(Value).ParseValue.(bool)
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
func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	value := v.Visit(ctx.GetRight()).(Value).ParseValue.(bool)
	return Value{ParseValue: !value}
}

// Range operator
func (v *Visitor) VisitRangeExpr(ctx *parser.RangeExprContext) interface{} {
	// Catch if the left and right are not integers
	l, okL := v.Visit(ctx.GetLeft()).(Value).ParseValue.(int64)
	r, okR := v.Visit(ctx.GetRight()).(Value).ParseValue.(int64)

	if !okL || !okR {
		v.NewError("Left and right values must be integers")
		return Value{ParseValue: nil}
	}

	if l > r {
		v.NewError("Left value is greater than right value")
		return Value{ParseValue: nil}
	}

	newVal := make([]int64, r-l+1)

	for i := l; i <= r; i++ {
		newVal[i-l] = i
	}

	return Value{ParseValue: newVal}
}

func (v *Visitor) VisitVariableType(ctx *parser.VariableTypeContext) interface{} {
	return ""
}
