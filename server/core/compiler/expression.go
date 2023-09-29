package compiler

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

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

		// if (leftT == V.IntType || leftT == V.FloatType) && (rightT == V.IntType || rightT == V.FloatType) {
		// }
		response = &ValueResponse{
			Type:         V.FloatType,
			Value:        c.TAC.NewTemporal(fmt.Sprintf("%s + %s", l, r), FloatTemporal), // Temporal
			ContextValue: TemporalType,
		}

	case "-":
		if leftT == V.IntType && rightT == V.IntType {
			response = &ValueResponse{
				Type:         V.IntType,
				Value:        c.TAC.NewTemporal(fmt.Sprintf("%s - %s", l, r), IntTemporal), // Temporal
				ContextValue: TemporalType,
			}
		}

		response = &ValueResponse{
			Type:         V.FloatType,
			Value:        c.TAC.NewTemporal(fmt.Sprintf("%s - %s", l, r), FloatTemporal), // Temporal
			ContextValue: TemporalType,
		}
		// if leftT == V.FloatType && rightT == V.FloatType {
		// 	return V.NewFloatValue(l.(float64) - r.(float64))
		// }
		// if leftT == V.FloatType && rightT == V.IntType {
		// 	return V.NewFloatValue(l.(float64) - float64(r.(int)))
		// }
		// if leftT == V.IntType && rightT == V.FloatType {
		// 	return V.NewFloatValue(float64(l.(int)) - r.(float64))
		// }
	case "*":
		if leftT == V.IntType && rightT == V.IntType {
			response = &ValueResponse{
				Type:         V.IntType,
				Value:        c.TAC.NewTemporal(fmt.Sprintf("%s * %s", l, r), IntTemporal), // Temporal
				ContextValue: TemporalType,
			}
		}
		// if leftT == V.FloatType && rightT == V.FloatType {
		// 	return V.NewFloatValue(l.(float64) * r.(float64))
		// }
		// if leftT == V.FloatType && rightT == V.IntType {
		// 	return V.NewFloatValue(l.(float64) * float64(r.(int)))
		// }
		// if leftT == V.IntType && rightT == V.FloatType {
		// 	return V.NewFloatValue(float64(l.(int)) * r.(float64))
		// }

		response = &ValueResponse{
			Type:         V.FloatType,
			Value:        c.TAC.NewTemporal(fmt.Sprintf("%s * %s", l, r), FloatTemporal), // Temporal
			ContextValue: TemporalType,
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

		response = &ValueResponse{
			Type:         V.FloatType,
			Value:        c.TAC.NewTemporal(fmt.Sprintf("%s / %s", l, r), FloatTemporal), // Temporal
			ContextValue: TemporalType,
		}
		// if leftT == V.FloatType && rightT == V.FloatType {
		// 	return V.NewFloatValue(l.(float64) / r.(float64))
		// }
		// if leftT == V.FloatType && rightT == V.IntType {
		// 	return V.NewFloatValue(l.(float64) / float64(r.(int)))
		// }
		// if leftT == V.IntType && rightT == V.FloatType {
		// 	return V.NewFloatValue(float64(l.(int)) / r.(float64))
		// }
	case "%":
		// if rightT == V.IntType && r.(int) == 0 {
		// 	return nil
		// }

		// if leftT == V.IntType && rightT == V.IntType {
		response = &ValueResponse{
			Type:         V.IntType,
			Value:        c.TAC.NewTemporal(fmt.Sprintf("%s %% %s", l, r), IntTemporal), // Temporal
			ContextValue: TemporalType,
		}
		// }
	}

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%s = %s %s %s", response.GetValue(), l, op, r),
		},
		fmt.Sprintf("Operación aritmética %s", op))
	return response
}

func (c *Compiler) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return &ValueResponse{
		Type:         V.FloatType,
		Value:        ctx.GetText(),
		ContextValue: LiteralType,
	}
}

func (c *Compiler) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	value := "1"

	if ctx.GetText() == "false" {
		value = "0"
	}

	return &ValueResponse{
		Type:         V.BooleanType,
		Value:        value,
		ContextValue: LiteralType,
	}
}

func (c *Compiler) VisitIDChain(ctx *parser.IDChainContext) interface{} {
	return ctx.AllID()
}

func (c *Compiler) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	expr := strings.Split(ctx.IdChain().GetText(), ".")
	id := expr[len(expr)-1]

	value := c.Env.GetValue(id)

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("t%d = stack[(int)%d]", c.TAC.TemporalQuantity(), value.GetAddress()),
		},
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

func (c *Compiler) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return c.Visit(ctx.Expr())
}

func (c *Compiler) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	s := strings.Trim(ctx.GetText(), "\"")

	if len(s) == 1 {
		return &ValueResponse{
			Type:         V.CharType,
			Value:        fmt.Sprintf("%d", s[0]),
			ContextValue: LiteralType,
		}
	}

	// Replace scape characters: double quote, backslash, new line, carriage return, tab
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\r", "\r")
	s = strings.ReplaceAll(s, "\\t", "\t")

	// Traverse string to insert chars in heap
	initialPointer := c.HeapPointer.GetPointer()

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("t%d = H", c.TAC.TemporalQuantity()),
		},
		"Obtención de la posición inicial del heap",
	)

	for _, char := range s {
		c.TAC.AppendCode(
			[]string{
				fmt.Sprintf("heap[(int)H] = %d", char),
				fmt.Sprintf("H = H + 1"),
			},
			fmt.Sprintf("Almacenamiento de caracter '%s' en el heap", string(char)),
		)
		c.HeapPointer.AddPointer()
	}

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("heap[(int)H] = %d", -1),
			fmt.Sprintf("H = H + 1"),
		},
		fmt.Sprintf("Almacenamiento del caracter nulo en el heap"),
	)
	c.HeapPointer.AddPointer()

	return &ValueResponse{
		Type:         V.StringType,
		Value:        c.TAC.NewTemporal(fmt.Sprintf("%d", initialPointer), IntTemporal),
		ContextValue: TemporalType,
	}
}
