package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitDecl(ctx *parser.DeclarationContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.Memory[id] = value
	return Value{ParseValue: true}
}
