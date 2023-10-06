package compiler

import (
	"OLC2/core/parser"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) arithmeticOp(l, r interface{}, op string, lc antlr.Token) interface{} {
	leftT := l.(*ValueResponse).GetType()
	rightT := r.(*ValueResponse).GetType()

	lV := l.(*ValueResponse).GetValue()
	rV := r.(*ValueResponse).GetValue()

	fmt.Println(lV, rV)

	var response *ValueResponse

	if leftT == StringTemporal && rightT == StringTemporal {
		return c.ConcatString(l.(*ValueResponse), r.(*ValueResponse))
	}

	if leftT == IntTemporal && rightT == IntTemporal {
		response = &ValueResponse{
			Type:        IntTemporal,
			Value:       c.TAC.NewTemporal(IntTemporal), // Temporal
			ContextType: TemporalType,
		}
	}

	if leftT == FloatTemporal || rightT == FloatTemporal {
		response = &ValueResponse{
			Type:        FloatTemporal,
			Value:       c.TAC.NewTemporal(FloatTemporal), // Temporal
			ContextType: TemporalType,
		}
	}

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%s = %s %s %s;", response.GetValue(), lV, op, rV),
		},
		fmt.Sprintf("Operación aritmética %s", op))
	return response
}

func (c *Compiler) VisitArithmeticExpr(ctx *parser.ArithmeticExprContext) interface{} {
	l := c.Visit(ctx.GetLeft()).(*ValueResponse)
	r := c.Visit(ctx.GetRight()).(*ValueResponse)
	op := ctx.GetOp().GetText()
	return c.arithmeticOp(l, r, op, ctx.GetStart())
}

func (c *Compiler) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	value := "1"

	if ctx.GetText() == "false" {
		value = "0"
	}

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       value,
		ContextType: LiteralType,
	}
}

func (c *Compiler) VisitCharExpr(ctx *parser.CharExprContext) interface{} {

	newTemporal := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%s = H;", newTemporal),
			fmt.Sprintf("heap[(int)H] = %d;", ctx.GetText()[1]),
			"H = H + 1;",
		},
		fmt.Sprintf("Almacenamiento de caracter '%d' en el heap", ctx.GetText()[1]),
	)

	c.HeapPointer.AddPointer()

	return &ValueResponse{
		Type:        CharTemporal,
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitComparisonExpr(ctx *parser.ComparisonExprContext) interface{} {
	left := c.Visit(ctx.GetLeft()).(*ValueResponse)
	right := c.Visit(ctx.GetRight()).(*ValueResponse)

	op := ctx.GetOp().GetText()

	trueLabel := c.TAC.NewLabel("")
	falseLabel := c.TAC.NewLabel("")
	temporalResult := c.TAC.NewTemporal(BooleanTemporal)

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("if (%s %s %s) goto %s;", left.GetValue(), op, right.GetValue(), trueLabel.String()),
			fmt.Sprintf("%s = 0;", temporalResult),
			fmt.Sprintf("goto %s;", falseLabel.String()),
			trueLabel.Declare(),
			fmt.Sprintf("%s = 1;", temporalResult),
			falseLabel.Declare(),
		},
		fmt.Sprintf("Operación %s %s %s", left.GetValue(), op, right.GetValue()),
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       temporalResult,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return &ValueResponse{
		Type:        FloatTemporal,
		Value:       ctx.GetText(),
		ContextType: LiteralType,
	}
}

func (c *Compiler) VisitIDChain(ctx *parser.IDChainContext) interface{} {
	return ctx.AllID()
}

func (c *Compiler) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	expr := strings.Split(ctx.IdChain().GetText(), ".")
	id := expr[len(expr)-1]

	value := c.Env.GetValue(id)

	newTemporal := c.TAC.NewTemporal(value.GetType())

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%s = stack[(int)%d];", newTemporal, value.GetAddress()),
		},
		fmt.Sprintf("Acceso a la variable '%s'", id),
	)

	return &ValueResponse{
		Type:        value.GetType(),
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return &ValueResponse{
		Type:        IntTemporal,
		Value:       ctx.GetText(),
		ContextType: LiteralType,
	}
}

func (c *Compiler) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	left := c.Visit(ctx.GetLeft()).(*ValueResponse)
	right := c.Visit(ctx.GetRight()).(*ValueResponse)

	op := ctx.GetOp().GetText()

	if op == "&&" {
		return c.And(left, right)
	}

	return c.Or(left, right)
}

func (c *Compiler) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return nil
}

func (c *Compiler) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return c.Visit(ctx.Expr())
}

func (c *Compiler) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	s := strings.Trim(ctx.GetText(), "\"")

	if len(s) == 1 {
		return &ValueResponse{
			Type:        CharTemporal,
			Value:       fmt.Sprintf("%d", s[0]),
			ContextType: LiteralType,
		}
	}

	// Replace scape characters: double quote, backslash, new line, carriage return, tab
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\r", "\r")
	s = strings.ReplaceAll(s, "\\t", "\t")

	// Traverse string to insert chars in heap
	newTemporal := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%s = H;", newTemporal),
		},
		"Obtención de la posición inicial del heap",
	)

	for _, char := range s {
		c.TAC.AppendCode(
			[]string{
				fmt.Sprintf("heap[(int)H] = %d;", char),
				"H = H + 1;",
			},
			fmt.Sprintf("Almacenamiento de caracter '%s' en el heap", string(char)),
		)
		c.HeapPointer.AddPointer()
	}

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("heap[(int)H] = %d;", -1),
			"H = H + 1;",
		},
		"Almacenamiento del caracter nulo en el heap",
	)
	c.HeapPointer.AddPointer()

	return &ValueResponse{
		Type:        StringTemporal,
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	newTemporal := c.TAC.NewTemporal(response.GetType())

	c.TAC.AppendCode(
		[]string{
			// fmt.Sprintf("t%d = %s * -1;", c.TAC.TemporalQuantity(), response.GetValue()),
			fmt.Sprintf("%s = %s * -1;", newTemporal, response.GetValue()),
		},
		fmt.Sprintf("Operación aritmética %s", "-"),
	)

	return &ValueResponse{
		Type:        response.GetType(),
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitVariableType(ctx *parser.VariableTypeContext) interface{} {
	switch ctx.GetText() {
	case "Int":
		return IntTemporal
	case "Float":
		return FloatTemporal
	case "Character":
		return CharTemporal
	case "Boolean":
		return BooleanTemporal
	case "String":
		return StringTemporal
	default:
		return ctx.GetText()
	}
}
