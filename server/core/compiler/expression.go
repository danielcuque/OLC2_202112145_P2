package compiler

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return &ValueResponse{
		Type:         V.IntType,
		Value:        ctx.GetText(),
		ContextValue: LiteralType,
	}
}

func (c *Compiler) VisitArithmeticExpr(ctx *parser.ArithmeticExprContext) interface{} {
	l := c.Visit(ctx.GetLeft()).(*ValueResponse)
	r := c.Visit(ctx.GetRight()).(*ValueResponse)
	op := ctx.GetOp().GetText()
	result := c.arithmeticOp(l, r, op, ctx.GetStart()).(*ValueResponse)
	c.TAC.AppendCode(fmt.Sprintf("%s = %s %s %s", result.GetValue(), l.GetValue(), op, r.GetValue()), "Arithmetic operation")
	return result
}

func (c *Compiler) arithmeticOp(l, r interface{}, op string, lc antlr.Token) interface{} {
	if l == nil || r == nil {
		return false
	}

	leftT := l.(*ValueResponse).GetType()
	rightT := r.(*ValueResponse).GetType()

	l = l.(*ValueResponse).GetValue()
	r = r.(*ValueResponse).GetValue()

	switch op {
	case "+":
		if leftT == V.IntType && rightT == V.IntType {
			return &ValueResponse{
				Type:         V.IntType,
				Value:        c.TAC.NewTemporal(fmt.Sprintf("%s + %s", l, r), IntTemporal), // Temporal
				ContextValue: TemporalType,
			}
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
	case "/":
		if rightT == V.IntType && r.(int) == 0 {
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
	case "%":
		if rightT == V.IntType && r.(int) == 0 {
			return nil
		}
		if leftT == V.IntType && rightT == V.IntType {
			return V.NewIntValue(l.(int) % r.(int))
		}
	}
	return nil
}
