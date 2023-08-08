package interfaces

import (
	"strconv"
	"strings"

	"OLC2/chore/parser"
	U "OLC2/chore/utils"
)

var intT = INT_STR
var floatT = FLOAT_STR
var stringT = STRING_STR

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	i, _ := strconv.Atoi(ctx.GetText())
	return NewIntValue(i)
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	f, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return NewFloatValue(f)
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	s := strings.Trim(ctx.GetText(), "\"")
	return NewStringValue(s)
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	b, _ := strconv.ParseBool(ctx.GetText())
	return NewBooleanValue(b)
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText()
	value, ok := v.Memory[id]
	if ok {
		return value
	} else {
		v.NewError("Error: Variable " + id)
		return nil
	}
}

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	op := v.Visit(ctx.Expr()).(IValue)
	value := op.GetValue()
	opType := op.GetType()

	if opType == intT {
		return NewIntValue(-value.(int))
	} else if opType == floatT {
		return NewFloatValue(-value.(float64))
	}
	v.NewError("Error: No se puede aplicar el operador unario - a " + opType)
	return nil
}

func (v *Visitor) VisitArithmeticExpr(ctx *parser.ArithmeticExprContext) interface{} {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	return v.arithmeticOp(l, r, op)
}

func (v *Visitor) arithmeticOp(l, r interface{}, op string) interface{} {
	if l == nil || r == nil {
		return false
	}

	leftT := l.(IValue).GetType()
	rightT := r.(IValue).GetType()

	l = l.(IValue).GetValue()
	r = r.(IValue).GetValue()

	switch op {
	case "+":

		if leftT == intT && rightT == intT {
			return NewIntValue(l.(int) + r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewFloatValue(l.(float64) + r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewFloatValue(l.(float64) + float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewFloatValue(float64(l.(int)) + r.(float64))
		}
		if leftT == stringT && rightT == stringT {
			return NewStringValue(l.(string) + r.(string))
		}
		v.NewError("Error: No se puede sumar " + leftT + " con " + rightT)
	case "-":
		if leftT == intT && rightT == intT {
			return NewIntValue(l.(int) - r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewFloatValue(l.(float64) - r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewFloatValue(l.(float64) - float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewFloatValue(float64(l.(int)) - r.(float64))
		}
		v.NewError("Error: No se puede restar " + leftT + " con " + rightT)
	case "*":
		if leftT == intT && rightT == intT {
			return NewIntValue(l.(int) * r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewFloatValue(l.(float64) * r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewFloatValue(l.(float64) * float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewFloatValue(float64(l.(int)) * r.(float64))
		}
		v.NewError("Error: No se puede multiplicar " + leftT + " con " + rightT)
	case "/":
		if leftT == intT && rightT == intT {
			return NewIntValue(l.(int) / r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewFloatValue(l.(float64) / r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewFloatValue(l.(float64) / float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewFloatValue(float64(l.(int)) / r.(float64))
		}
		v.NewError("Error: No se puede dividir " + leftT + " con " + rightT)
	case "%":
		if leftT == intT && rightT == intT {
			return NewIntValue(l.(int) % r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewIntValue(int(l.(float64)) % int(r.(float64)))
		}
		if leftT == floatT && rightT == intT {
			return NewIntValue(int(l.(float64)) % r.(int))
		}
		if leftT == intT && rightT == floatT {
			return NewIntValue(l.(int) % int(r.(float64)))
		}
		v.NewError("Error: No se puede modular " + leftT + " con " + rightT)
	}

	return nil
}

func (v *Visitor) VisitComparasionExpr(ctx *parser.ComparasionExprContext) interface{} {

	// Parse left and right
	l := v.Visit(ctx.GetLeft()).(IValue).GetValue()
	r := v.Visit(ctx.GetRight()).(IValue).GetValue()
	op := ctx.GetOp().GetText()

	// Get types
	leftT := U.GetType(l)
	rightT := U.GetType(r)

	// Compare
	switch op {
	case "==":
		if leftT == intT && rightT == intT {
			return NewBooleanValue(l.(int) == r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewBooleanValue(l.(float64) == r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewBooleanValue(l.(float64) == float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewBooleanValue(float64(l.(int)) == r.(float64))
		}
		if leftT == stringT && rightT == stringT {
			return NewBooleanValue(l.(string) == r.(string))
		}
	case "!=":
		if leftT == intT && rightT == intT {
			return NewBooleanValue(l.(int) != r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewBooleanValue(l.(float64) != r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewBooleanValue(l.(float64) != float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewBooleanValue(float64(l.(int)) != r.(float64))
		}
		if leftT == stringT && rightT == stringT {
			return NewBooleanValue(l.(string) != r.(string))
		}
	case ">":
		if leftT == intT && rightT == intT {
			return NewBooleanValue(l.(int) > r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewBooleanValue(l.(float64) > r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewBooleanValue(l.(float64) > float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewBooleanValue(float64(l.(int)) > r.(float64))
		}
	case "<":
		if leftT == intT && rightT == intT {
			return NewBooleanValue(l.(int) < r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewBooleanValue(l.(float64) < r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewBooleanValue(l.(float64) < float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewBooleanValue(float64(l.(int)) < r.(float64))
		}
	case ">=":
		if leftT == intT && rightT == intT {
			return NewBooleanValue(l.(int) >= r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewBooleanValue(l.(float64) >= r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewBooleanValue(l.(float64) >= float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewBooleanValue(float64(l.(int)) >= r.(float64))
		}
	case "<=":
		if leftT == intT && rightT == intT {
			return NewBooleanValue(l.(int) <= r.(int))
		}
		if leftT == floatT && rightT == floatT {
			return NewBooleanValue(l.(float64) <= r.(float64))
		}
		if leftT == floatT && rightT == intT {
			return NewBooleanValue(l.(float64) <= float64(r.(int)))
		}
		if leftT == intT && rightT == floatT {
			return NewBooleanValue(float64(l.(int)) <= r.(float64))
		}
	}

	v.NewError("Error: No se puede comparar " + leftT + " con " + rightT)
	return nil
}

// Logical operators

func (v *Visitor) VisitTernaryExpr(ctx *parser.TernaryExprContext) interface{} {
	cond := v.Visit(ctx.GetCondition()).(IValue).GetValue().(bool)
	if cond {
		return v.Visit(ctx.GetCTrue())
	} else {
		return v.Visit(ctx.GetCFalse())
	}

}

func (v *Visitor) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	l := v.Visit(ctx.GetLeft()).(IValue).GetValue().(bool)
	r := v.Visit(ctx.GetRight()).(IValue).GetValue().(bool)
	op := ctx.GetOp().GetText()

	operators := map[string]func(bool, bool) bool{
		"&&": func(a, b bool) bool { return a && b },
		"||": func(a, b bool) bool { return a || b },
	}

	if f, ok := operators[op]; ok {
		return NewBooleanValue(f(l, r))
	}

	return nil
}

// Not operator
func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	value := v.Visit(ctx.GetRight()).(IValue).GetValue().(bool)
	return NewBooleanValue(!value)
}

// Range operator
func (v *Visitor) VisitRangeExpr(ctx *parser.RangeExprContext) interface{} {
	// Catch if the left and right are not integers
	l, okL := v.Visit(ctx.GetLeft()).(IValue).GetValue().(int)
	r, okR := v.Visit(ctx.GetRight()).(IValue).GetValue().(int)

	if !okL || !okR {
		v.NewError("Left and right values must be integers")
		return nil
	}

	if l > r {
		v.NewError("Left value is greater than right value")
		return nil
	}

	newVal := make([]IValue, r-l+1)

	for i := l; i <= r; i++ {
		newVal[i-l] = NewIntValue(i)
	}

	return NewArrayValue(newVal)
}

func (v *Visitor) VisitVariableType(ctx *parser.VariableTypeContext) interface{} {
	return ""
}
