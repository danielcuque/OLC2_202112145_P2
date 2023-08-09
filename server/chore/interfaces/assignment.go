package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(IValue)

	variable, ok := v.Scope.GetVariable(id).(*Variable) // Pointer to Variable

	if !ok {
		v.NewError(fmt.Sprintf("La variable %s no existe", id), ctx.GetStart())
		return false
	}

	if variable.IsConstant() {
		v.NewError(fmt.Sprintf("La variable %s es constante", id), ctx.GetStart())
		return false
	}

	if variable.Value.GetType() != value.GetType() {
		v.NewError(fmt.Sprintf("No se puede asignar %s a %s", value.GetType(), variable.Value.GetType()), ctx.GetStart())
		return false
	}

	variable.SetValue(value)

	return nil
}
