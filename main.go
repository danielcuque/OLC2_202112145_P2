package main

import (
	"OLC2/chore/parser"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseSwiftVisitor
}

func (v *Visitor) Visit(tree antlr.ParseTree) float64 {
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

func (v *Visitor) VisitOp(ctx *parser.OpContext) float64 {
	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()

	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		return 0
	}
}

func (v *Visitor) VisitDigit(ctx *parser.DigitContext) float64 {
	i1, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return i1
}

func (v *Visitor) VisitParen(ctx *parser.ParenContext) float64 {
	tar := v.Visit(ctx.Expr())
	return tar
}

func main() {
	expression := "3*(5+4)+"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewSwiftLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSwiftParser(stream)
	p.BuildParseTrees = true
	tree := p.Expr()
	var visitor = Visitor{}
	var result = visitor.Visit(tree)
	fmt.Println(expression, "=", result)

}
