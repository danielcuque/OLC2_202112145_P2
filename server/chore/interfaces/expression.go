package interfaces

import (
	"strconv"
	"strings"

	"OLC2/chore/parser"
	V "OLC2/chore/values"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	i, _ := strconv.Atoi(ctx.GetText())
	return V.NewIntValue(i)
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	f, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return V.NewFloatValue(f)
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	// Check if is posible char or string
	s := strings.Trim(ctx.GetText(), "\"")

	if len(s) == 1 {
		return V.NewCharValue([]rune(s)[0])
	}

	// Replace scape characters: double quote, backslash, new line, carriage return, tab
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\r", "\r")
	s = strings.ReplaceAll(s, "\\t", "\t")
	return V.NewStringValue(s)
}

func (v *Visitor) VisitCharExpr(ctx *parser.CharExprContext) interface{} {
	c := []rune(strings.Trim(ctx.GetText(), "'"))[0]
	return V.NewCharValue(c)
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	b, _ := strconv.ParseBool(ctx.GetText())
	return V.NewBooleanValue(b)
}

func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	return V.NewNilValue(nil)
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
	op := v.Visit(ctx.Expr()).(V.IValue)
	value := op.GetValue()
	opType := op.GetType()

	if opType == V.IntType {
		return V.NewIntValue(-value.(int))
	} else if opType == V.FloatType {
		return V.NewFloatValue(-value.(float64))
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

	leftT := l.(V.IValue).GetType()
	rightT := r.(V.IValue).GetType()

	l = l.(V.IValue).GetValue()
	r = r.(V.IValue).GetValue()

	switch op {
	case "+":

		if leftT == V.IntType && rightT == V.IntType {
			return V.NewIntValue(l.(int) + r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewFloatValue(l.(float64) + r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewFloatValue(l.(float64) + float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewFloatValue(float64(l.(int)) + r.(float64))
		}
		if leftT == V.StringType && rightT == V.StringType {
			return V.NewStringValue(l.(string) + r.(string))
		}
		// Sumamos string con char
		if leftT == V.StringType && rightT == V.CharType {
			return V.NewStringValue(l.(string) + string(r.(rune)))
		}
		// Sumamos char con string
		if leftT == V.CharType && rightT == V.StringType {
			return V.NewStringValue(string(l.(rune)) + r.(string))
		}
		v.NewError("No se puede sumar "+leftT+" con "+rightT, lc)
	case "-":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewIntValue(l.(int) - r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewFloatValue(l.(float64) - r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewFloatValue(l.(float64) - float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewFloatValue(float64(l.(int)) - r.(float64))
		}
		v.NewError("No se puede restar "+leftT+" con "+rightT, lc)
	case "*":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewIntValue(l.(int) * r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewFloatValue(l.(float64) * r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewFloatValue(l.(float64) * float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewFloatValue(float64(l.(int)) * r.(float64))
		}
		v.NewError("No se puede multiplicar "+leftT+" con "+rightT, lc)
	case "/":
		if rightT == V.IntType && r.(int) == 0 {
			v.NewError("No se puede dividir entre 0", lc)
			return nil
		}

		if leftT == V.IntType && rightT == V.IntType {
			return V.NewIntValue(l.(int) / r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewFloatValue(l.(float64) / r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewFloatValue(l.(float64) / float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewFloatValue(float64(l.(int)) / r.(float64))
		}
		v.NewError("No se puede dividir "+leftT+" con "+rightT, lc)
	case "%":
		if rightT == V.IntType && r.(int) == 0 {
			v.NewError("No se puede dividir entre 0", lc)
			return nil
		}
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewIntValue(l.(int) % r.(int))
		}
		v.NewError("No se puede modular "+leftT+" con "+rightT, lc)
	}

	return nil
}

func (v *Visitor) VisitComparasionExpr(ctx *parser.ComparasionExprContext) interface{} {

	// Parse left and right
	lV, okL := v.Visit(ctx.GetLeft()).(V.IValue)
	rV, okR := v.Visit(ctx.GetRight()).(V.IValue)

	if !okL || !okR {
		v.NewError("No es posible realizar la comparación", ctx.GetStart())
		return nil
	}

	l := lV.GetValue()
	r := rV.GetValue()

	op := ctx.GetOp().GetText()

	// Get types
	leftT := v.Visit(ctx.GetLeft()).(V.IValue).GetType()
	rightT := v.Visit(ctx.GetRight()).(V.IValue).GetType()

	// Compare
	switch op {
	case "==":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewBooleanValue(l.(int) == r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewBooleanValue(l.(float64) == r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewBooleanValue(l.(float64) == float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewBooleanValue(float64(l.(int)) == r.(float64))
		}
		if leftT == V.StringType && rightT == V.StringType {
			return V.NewBooleanValue(l.(string) == r.(string))
		}
		if leftT == V.CharType && rightT == V.CharType {
			return V.NewBooleanValue(l.(rune) == r.(rune))
		}

	case "!=":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewBooleanValue(l.(int) != r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewBooleanValue(l.(float64) != r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewBooleanValue(l.(float64) != float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewBooleanValue(float64(l.(int)) != r.(float64))
		}
		if leftT == V.StringType && rightT == V.StringType {
			return V.NewBooleanValue(l.(string) != r.(string))
		}
		if leftT == V.CharType && rightT == V.CharType {
			return V.NewBooleanValue(l.(rune) != r.(rune))
		}
	case ">":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewBooleanValue(l.(int) > r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewBooleanValue(l.(float64) > r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewBooleanValue(l.(float64) > float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewBooleanValue(float64(l.(int)) > r.(float64))
		}
		if leftT == V.StringType && rightT == V.StringType {
			return V.NewBooleanValue(l.(string) > r.(string))
		}
		if leftT == V.CharType && rightT == V.CharType {
			return V.NewBooleanValue(l.(rune) > r.(rune))
		}
	case "<":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewBooleanValue(l.(int) < r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewBooleanValue(l.(float64) < r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewBooleanValue(l.(float64) < float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewBooleanValue(float64(l.(int)) < r.(float64))
		}
		if leftT == V.StringType && rightT == V.StringType {
			return V.NewBooleanValue(l.(string) < r.(string))
		}
	case ">=":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewBooleanValue(l.(int) >= r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewBooleanValue(l.(float64) >= r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewBooleanValue(l.(float64) >= float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewBooleanValue(float64(l.(int)) >= r.(float64))
		}
		if leftT == V.StringType && rightT == V.StringType {
			return V.NewBooleanValue(l.(string) >= r.(string))
		}
		if leftT == V.CharType && rightT == V.CharType {
			return V.NewBooleanValue(l.(rune) >= r.(rune))
		}
	case "<=":
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewBooleanValue(l.(int) <= r.(int))
		}
		if leftT == V.FloatType && rightT == V.FloatType {
			return V.NewBooleanValue(l.(float64) <= r.(float64))
		}
		if leftT == V.FloatType && rightT == V.IntType {
			return V.NewBooleanValue(l.(float64) <= float64(r.(int)))
		}
		if leftT == V.IntType && rightT == V.FloatType {
			return V.NewBooleanValue(float64(l.(int)) <= r.(float64))
		}
		if leftT == V.StringType && rightT == V.StringType {
			return V.NewBooleanValue(l.(string) <= r.(string))
		}
		if leftT == V.CharType && rightT == V.CharType {
			return V.NewBooleanValue(l.(rune) <= r.(rune))
		}
	}

	v.NewError("No se puede comparar "+leftT+" con "+rightT, ctx.GetStart())
	return nil
}

// Logical operators

func (v *Visitor) VisitTernaryExpr(ctx *parser.TernaryExprContext) interface{} {
	cond := v.Visit(ctx.GetCondition()).(V.IValue).GetValue().(bool)
	if cond {
		return v.Visit(ctx.GetCTrue())
	} else {
		return v.Visit(ctx.GetCFalse())
	}

}

func (v *Visitor) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	l := v.Visit(ctx.GetLeft()).(V.IValue).GetValue().(bool)
	r := v.Visit(ctx.GetRight()).(V.IValue).GetValue().(bool)
	op := ctx.GetOp().GetText()

	operators := map[string]func(bool, bool) bool{
		"&&": func(a, b bool) bool { return a && b },
		"||": func(a, b bool) bool { return a || b },
	}

	if f, ok := operators[op]; ok {
		return V.NewBooleanValue(f(l, r))
	}

	return nil
}

// Not operator
func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	value := v.Visit(ctx.GetRight()).(V.IValue).GetValue().(bool)
	return V.NewBooleanValue(!value)
}

// Range operator
func (v *Visitor) VisitRangeExpr(ctx *parser.RangeExprContext) interface{} {
	// Catch if the left and right are not integers
	l, okL := v.Visit(ctx.GetLeft()).(V.IValue).GetValue().(int)
	r, okR := v.Visit(ctx.GetRight()).(V.IValue).GetValue().(int)

	if !okL || !okR {
		v.NewError("Los valores de rango deben ser enteros", ctx.GetStart())
		return nil
	}

	if l > r {
		v.NewError("El valor izquierdo del rango debe ser menor o igual al valor derecho", ctx.GetStart())
		return nil
	}

	newVal := make([]V.IValue, r-l+1)

	for i := l; i <= r; i++ {
		newVal[i-l] = V.NewIntValue(i)
	}

	return V.NewRangeValue(newVal)
}

func (v *Visitor) VisitVariableType(ctx *parser.VariableTypeContext) interface{} {
	return ctx.GetText()
}
