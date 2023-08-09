package interfaces

import (
	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		v.NewError(val.GetText(), val.GetSymbol())
		return nil
	default:
		return tree.Accept(v)
	}
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	for i := 0; ctx.Statement(i) != nil; i++ {
		v.Visit(ctx.Statement(i))
	}
	return nil
}
