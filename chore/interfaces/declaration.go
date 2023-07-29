package interfaces

import (
	"OLC2/chore/parser"
)

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	isIdExist, ok := v.Memory[id]

	if ok && isIdExist.ParseValue != nil {
		panic("Variable already exist")
	}

	v.Memory[id] = value
	return Value{ParseValue: true}
}
