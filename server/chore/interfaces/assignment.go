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
		if variable.Value.GetType() == floatT && value.GetType() == intT {
			value = NewFloatValue(float64(value.GetValue().(int)))
		} else {
			v.NewError(fmt.Sprintf("El tipo de la variable %s no coincide con el valor asignado, se esperaba %s y se obtuvo %s", id, variable.Value.GetType(), value.GetType()), ctx.GetStart())
			return false
		}
	}

	// Check the operator assignment (= += -=)
	op := ctx.GetOp().GetText()

	if op != "=" {
		if op == "+=" {
			variable.SetValue(v.arithmeticOp(variable.Value, value, "+", ctx.GetStart()).(IValue))
		} else if op == "-=" {
			variable.SetValue(v.arithmeticOp(variable.Value, value, "-", ctx.GetStart()).(IValue))
		} else {
			v.NewError(fmt.Sprintf("No se puede aplicar el operador %s a %s", op, variable.Value.GetType()), ctx.GetStart())
			return false
		}
	} else {
		variable.SetValue(value)
	}

	return nil
}
