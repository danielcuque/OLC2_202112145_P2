package interfaces

import (
	"OLC2/chore/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Value struct {
	value interface{}
}

type Visitor struct {
	parser.SwiftVisitor
	Memory map[string]Value
}

func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch value := tree.(type) {
	case *parser.OpContext:
		return v.VisitOp(value)
	case *parser.DigitContext:
		return v.VisitDigit(value)
	case *parser.ParenContext:
		return v.VisitParen(value)
	default:
		panic("Unknown context")
	}
}

func (v *Visitor) VisitOp(ctx *parser.OpContext) Value {

	left := v.Visit(ctx.GetLeft()).value.(int64)
	right := v.Visit(ctx.GetRight()).value.(int64)
	op := ctx.GetOp().GetText()

	switch op {
	case "+":
		return Value{
			value: left + right,
		}
	case "-":
		return Value{
			value: left - right,
		}
	case "*":
		return Value{
			value: left * right,
		}
	case "/":
		return Value{
			value: left / right,
		}
	default:
		return Value{
			value: false,
		}
	}
}

func (v *Visitor) VisitDigit(ctx *parser.DigitContext) Value {
	i1, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Value{
		value: i1,
	}
}

func (v *Visitor) VisitParen(ctx *parser.ParenContext) Value {
	tar := v.Visit(ctx.Expr())
	return tar
}
