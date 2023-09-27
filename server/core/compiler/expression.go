package compiler

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	expr := strings.Split(ctx.IdChain().GetText(), ".")
	id := expr[len(expr)-1]

	value := c.Env.GetValue(id)

	c.TAC.AppendCode(
		fmt.Sprintf("t%d = stack[(int)%d]", c.TAC.TemporalQuantity(), value.GetAddress()),
		fmt.Sprintf("Acceso a la variable '%s'", id),
	)

	return &ValueResponse{
		Type:         V.IntType,
		Value:        c.TAC.NewTemporal(fmt.Sprintf("stack[(int)P]"), IntTemporal),
		ContextValue: TemporalType,
	}

}

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
	return c.arithmeticOp(l, r, op, ctx.GetStart())
}

func (c *Compiler) arithmeticOp(l, r interface{}, op string, lc antlr.Token) interface{} {
	leftT := l.(*ValueResponse).GetType()
	rightT := r.(*ValueResponse).GetType()

	l = l.(*ValueResponse).GetValue()
	r = r.(*ValueResponse).GetValue()

	var response *ValueResponse

	switch op {
	case "+":
		if leftT == V.IntType && rightT == V.IntType {
			response = &ValueResponse{
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
		if leftT == V.StringType && rightT == V.CharType {
			return V.NewStringValue(l.(string) + string(r.(rune)))
		}
		if leftT == V.CharType && rightT == V.StringType {
			return V.NewStringValue(string(l.(rune)) + r.(string))
		}
	case "-":
		if leftT == V.IntType && rightT == V.IntType {
			response = &ValueResponse{
				Type:         V.IntType,
				Value:        c.TAC.NewTemporal(fmt.Sprintf("%s - %s", l, r), IntTemporal), // Temporal
				ContextValue: TemporalType,
			}
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
			response = &ValueResponse{
				Type:         V.IntType,
				Value:        c.TAC.NewTemporal(fmt.Sprintf("%s * %s", l, r), IntTemporal), // Temporal
				ContextValue: TemporalType,
			}
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
		// TODO: macro to display error if division by zero

		// if rightT == V.IntType && r == "0" {
		// 	return nil
		// }

		if leftT == V.IntType && rightT == V.IntType {
			response = &ValueResponse{
				Type:         V.IntType,
				Value:        c.TAC.NewTemporal(fmt.Sprintf("%s / %s", l, r), IntTemporal), // Temporal
				ContextValue: TemporalType,
			}
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
			response = &ValueResponse{
				Type:         V.IntType,
				Value:        c.TAC.NewTemporal(fmt.Sprintf("%s %% %s", l, r), IntTemporal), // Temporal
				ContextValue: TemporalType,
			}
		}
	}

	c.TAC.AppendCode(fmt.Sprintf("%s = %s %s %s", response.GetValue(), l, op, r), fmt.Sprintf("Operación aritmética %s", op))
	return response
}
