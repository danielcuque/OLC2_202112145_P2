package interfaces

import (
	"fmt"
	"strconv"
	"strings"

	"OLC2/core/parser"
	V "OLC2/core/values"

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

// Returns pointer or value, if amper is true returns pointer
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	expr := strings.Split(ctx.IdChain().GetText(), ".")

	if IsStructMethodRunning && expr[0] == "self" {
		return v.HandleVisitIdStruct(ctx)
	}

	variable := v.Env.GetVariable(expr[0])

	if variable == nil {
		v.NewError(fmt.Sprintf("La variable %s no existe", expr[0]), ctx.GetStart())
		return nil
	}

	amper := ctx.Kw_AMPER() != nil

	if len(expr) == 1 {
		if amper {
			return variable
		}
		return variable.Value
	}

	// Check if there is more than one id

	params := expr[1:]

	prop, okP := GetPropValue(variable, params).(*Variable)

	if prop == nil {
		v.NewError(fmt.Sprint("La propiedad ", params[len(params)-1], " no existe"), ctx.GetStart())
		return V.NewNilValue(nil)
	}

	if !okP {
		v.NewError(fmt.Sprint("La propiedad ", params[len(params)-1], " no existe"), ctx.GetStart())
		return V.NewNilValue(nil)
	}

	if amper {
		return prop
	}
	return prop.Value
}

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *Visitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	expr := v.Visit(ctx.Expr()).(V.IValue)
	value := expr.GetValue()
	valueType := expr.GetType()

	if valueType == V.IntType {
		return V.NewIntValue(-value.(int))
	} else if valueType == V.FloatType {
		return V.NewFloatValue(-value.(float64))
	}

	v.NewError("No se puede aplicar el operador unario - a "+valueType, ctx.GetStart())
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
			return V.NewIntValue(l.(int) * r.(int))
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

func (v *Visitor) VisitComparisonExpr(ctx *parser.ComparisonExprContext) interface{} {

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
	value, ok := v.Visit(ctx.Expr()).(V.IValue)

	if !ok {
		v.NewError("No se puede aplicar el operador not", ctx.GetStart())
		return nil
	}

	if value.GetType() != V.BooleanType {
		v.NewError("No se puede aplicar el operador not a "+value.GetType(), ctx.GetStart())
		return nil
	}

	return V.NewBooleanValue(!value.GetValue().(bool))
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

	switch ctx.GetText() {
	case "Int":
		return V.IntType
	case "Float":
		return V.FloatType
	case "String":
		return V.StringType
	case "Character":
		return V.CharType
	case "Bool":
		return V.BooleanType
	case "nil":
		return V.NilType
	}

	// Check if is a struct
	if structType := v.Env.GetStruct(ctx.GetText()); structType != nil {
		return structType.GetType()
	}

	v.NewError(fmt.Sprintf("El tipo %s no existe", ctx.GetText()), ctx.GetStart())
	return nil
}

func (v *Visitor) VisitFunctionCallExpr(ctx *parser.FunctionCallExprContext) interface{} {
	return v.Visit(ctx.FunctionCall())
}

func (v *Visitor) VisitVectorAccessExpr(ctx *parser.VectorAccessExprContext) interface{} {
	dict, ok := v.Visit(ctx.VectorAccess()).(map[string]interface{})

	if !ok {
		v.NewError("No se puede acceder al vector", ctx.GetStart())
		return nil
	}

	return dict["value"].(V.IValue).Copy()
}

func (v *Visitor) VisitMatrixAccessExpr(ctx *parser.MatrixAccessExprContext) interface{} {
	dict, ok := v.Visit(ctx.MatrixAccess()).(map[string]interface{})

	if !ok {
		v.NewError("No se puede acceder a la matriz", ctx.GetStart())
		return nil
	}

	return dict["value"]
}

func (v *Visitor) VisitObjectChain(ctx *parser.ObjectChainContext) interface{} {
	dict, ok := v.Visit(ctx.VectorAccess()).(map[string]interface{})

	if !ok {
		v.NewError("No se puede acceder al objeto", ctx.GetStart())
		return nil
	}

	value := dict["value"].(*ObjectV)

	props := make([]string, 0)

	for _, prop := range ctx.AllID() {
		props = append(props, prop.GetText())
	}

	prop, okP := GetObjectPropValue(value, props).(*Variable)

	if !okP {
		v.NewError("No se puede acceder a la propiedad", ctx.GetStart())
		return nil
	}

	return prop
}

func (v *Visitor) VisitObjectChainExpr(ctx *parser.ObjectChainExprContext) interface{} {
	return v.Visit(ctx.ObjectChain()).(V.IValue).Copy()
}
