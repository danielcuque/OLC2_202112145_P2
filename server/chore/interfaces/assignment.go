package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(IValue)

	// Check if the variable exists
	_, ok := v.Memory[id]
	if !ok {
		v.NewError(fmt.Sprintf("La variable %s no existe", id))
		fmt.Println("Variable doesn't exist", id)
		return nil
	}

	// Check if the types are the same
	if v.Memory[id].GetType() != value.GetType() {
		v.NewError(fmt.Sprintf("La variable %s, es de tipo %s y se le quiere asignar un valor de tipo %s", id, v.Memory[id].GetType(), value.GetType()))
		fmt.Println("Variable is not the same type", id)
		return nil
	}

	// Assign the value
	v.Memory[id] = value
	return true
}
