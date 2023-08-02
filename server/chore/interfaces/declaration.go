package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	isIdExist, ok := v.Memory[id]

	if ok && isIdExist.ParseValue != nil {
		v.NewError(fmt.Sprintf("Variable %s already exist", id))
		return Value{ParseValue: nil}
	}

	v.Memory[id] = value
	return Value{ParseValue: true}
}
