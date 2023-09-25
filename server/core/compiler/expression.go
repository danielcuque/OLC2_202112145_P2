package compiler

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
)

func (c *Compiler) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return &ValueResponse{
		Type:         V.IntType,
		Value:        ctx.GetText(),
		ContextValue: LiteralType,
	}
}
