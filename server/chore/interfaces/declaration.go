package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	value, okVal := v.Visit(ctx.Expr()).(IValue)

	if !okVal {
		return nil
	}

	_, ok := v.Memory[id]

	if ok {
		v.NewError(fmt.Sprintf("Variable %s already exist", id))
		return false
	}

	v.Memory[id] = value

	return true
}
