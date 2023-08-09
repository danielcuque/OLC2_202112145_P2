package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	varType := ctx.GetVarType().GetText()
	id := ctx.ID().GetText()
	value, okVal := v.Visit(ctx.Expr()).(IValue)

	if !okVal {
		return nil
	}

	_, ok := v.Scope.GetVariable(id).(Variable)

	if ok {
		v.NewError(fmt.Sprintf("Error: La variable %s ya existe", id))
		return false
	}

	newVariable := NewVariable(id, varType == "let", value, value.GetType())

	v.Scope.AddVariable(id, newVariable)

	return true
}
