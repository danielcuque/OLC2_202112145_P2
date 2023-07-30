package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())

	// Check if the variable exists
	_, ok := v.Memory[id]
	if !ok {
		v.NewError(fmt.Sprintf("Variable %s doesn't exist", id))
		fmt.Println("Variable doesn't exist", id)
		return Value{ParseValue: nil}
	}

	// Check if value is not nil
	if value.ParseValue == nil {
		v.NewError(fmt.Sprintf("Variable %s is nil", id))
		fmt.Println("Variable is nil", id)
		return Value{ParseValue: nil}
	}

	// Check if the types are the same
	if v.Memory[id].Type != value.Type {
		v.NewError(fmt.Sprintf("Variable %s is not the same type", id))
		fmt.Println("Variable is not the same type", id)
		return Value{ParseValue: nil}
	}

	// Assign the value
	v.Memory[id] = value
	return Value{ParseValue: true}
}
