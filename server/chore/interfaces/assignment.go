package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

func (v *Visitor) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	id := ctx.ID().GetText()
	value, ok := v.Visit(ctx.Expr()).(V.IValue)

	if !ok {
		v.NewError(InvalidExpression, ctx.GetStart())
		return false
	}

	variable, ok := v.Env.GetVariable(id).(*Variable) // Pointer to Variable

	if !ok {
		v.NewError(fmt.Sprintf("La variable %s no existe", id), ctx.GetStart())
		return false
	}

	if variable.IsConstant() {
		v.NewError(fmt.Sprintf("La variable %s es constante", id), ctx.GetStart())
		return false
	}

	if variable.GetType() != value.GetType() {
		if variable.GetType() == V.FloatType && value.GetType() == V.IntType {
			value = V.NewFloatValue(float64(value.GetValue().(int)))
		} else {
			v.NewError(fmt.Sprintf("El tipo de la variable %s no coincide con el valor asignado, se esperaba %s y se obtuvo %s", id, variable.GetType(), value.GetType()), ctx.GetStart())
			return false
		}
	}

	// Check the operator assignment (= += -=)
	op := ctx.GetOp().GetText()
	switch op {
	case "=":
		variable.SetValue(value)
	case "+=":
		variable.SetValue(v.arithmeticOp(variable.Value, value, "+", ctx.GetStart()).(V.IValue))
	case "-=":
		variable.SetValue(v.arithmeticOp(variable.Value, value, "-", ctx.GetStart()).(V.IValue))
	default:
		v.NewError(fmt.Sprintf("No se puede aplicar el operador %s a %s", op, variable.GetType()), ctx.GetStart())
		return false
	}

	return nil

}
