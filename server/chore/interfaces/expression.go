package interfaces

import (
	"strconv"
	"strings"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
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
	// Check if is posible char or string
	s := strings.Trim(ctx.GetText(), "\"")

	if len(s) == 1 {
		return NewCharValue([]rune(s)[0])
	}
	return NewStringValue(s)
}

func (v *Visitor) VisitCharExpr(ctx *parser.CharExprContext) interface{} {
	c := []rune(strings.Trim(ctx.GetText(), "'"))[0]
	return NewCharValue(c)
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	b, _ := strconv.ParseBool(ctx.GetText())
	return NewBooleanValue(b)
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {

	id := ctx.GetText()
	value := v.Scope.GetVariable(id)
	if value == nil {
		v.NewError("La variable "+id+" no existe", ctx.GetStart())
		return nil
	}
	return value.(*Variable).Value
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
	v.NewError("No se puede aplicar el operador unario - a "+opType, ctx.GetStart())
	return nil
}

func (v *Visitor) VisitArithmeticExpr(ctx *parser.ArithmeticExprContext) interface{} {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	return v.arithmeticOp(l, r, op, ctx.GetStart())
}

func (v *Visitor) arithmeticOp(l, r interface{}, op string, lc antlr.Token) interface{} {
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
		v.NewError("No se puede sumar "+leftT+" con "+rightT, lc)
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
		v.NewError("No se puede restar "+leftT+" con "+rightT, lc)
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
		v.NewError("No se puede multiplicar "+leftT+" con "+rightT, lc)
	case "/":
		if rightT == intT && r.(int) == 0 {
			v.NewError("No se puede dividir entre 0", lc)
			return nil
		}

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
		v.NewError("No se puede dividir "+leftT+" con "+rightT, lc)
	case "%":
		if rightT == intT && r.(int) == 0 {
			v.NewError("No se puede dividir entre 0", lc)
			return nil
		}
		if leftT == intT && rightT == intT {
			return NewIntValue(l.(int) % r.(int))
		}
		v.NewError("No se puede modular "+leftT+" con "+rightT, lc)
	}

	return nil
}

func (v *Visitor) VisitComparasionExpr(ctx *parser.ComparasionExprContext) interface{} {

	// Parse left and right
	l := v.Visit(ctx.GetLeft()).(IValue).GetValue()
	r := v.Visit(ctx.GetRight()).(IValue).GetValue()
	op := ctx.GetOp().GetText()

	// Get types
	leftT := v.Visit(ctx.GetLeft()).(IValue).GetType()
	rightT := v.Visit(ctx.GetRight()).(IValue).GetType()

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

	v.NewError("No se puede comparar "+leftT+" con "+rightT, ctx.GetStart())
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
		v.NewError("Los valores de rango deben ser enteros", ctx.GetStart())
		return nil
	}

	if l > r {
		v.NewError("El valor izquierdo del rango debe ser menor o igual al valor derecho", ctx.GetStart())
		return nil
	}

	newVal := make([]IValue, r-l+1)

	for i := l; i <= r; i++ {
		newVal[i-l] = NewIntValue(i)
	}

	return NewArrayValue(newVal)
}

func (v *Visitor) VisitVariableType(ctx *parser.VariableTypeContext) interface{} {
	return ctx.GetText()
}
