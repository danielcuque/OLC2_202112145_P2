package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	ids := v.Visit(ctx.IdChain()).([]antlr.TerminalNode)
	id := ids[0].GetText()

	if IsStructMethodRunning && !IsStructMethodMutating && id == "self" {
		v.NewError("No se puede modificar un struct desde un método que no sea mutante", ctx.GetStart())
		return false
	}

	// TODO: verify if declaration is in struct
	value, ok := v.Visit(ctx.Expr()).(V.IValue)

	if !ok {
		v.NewError(InvalidExpression, ctx.GetStart())
		return false
	}

	variable := v.Env.GetVariable(id)

	if IsStructMethodRunning && id == "self" {
		variable = objectStruct.Env.Variables[ids[1].GetText()]
	}

	if variable == nil {
		v.NewError(fmt.Sprintf("La variable %s no existe", id), ctx.GetStart())
		return false
	}

	if len(ids) > 1 && ids[0].GetText() != "self" {
		props := []string{ids[1].GetText()}

		if len(ids) > 2 {
			props = v.GetProps(ids)
		}

		prop, ok := GetPropValue(variable, props).(*Variable)

		if !ok {
			v.NewError(fmt.Sprintf("La variable %s no es un objeto", id), ctx.GetStart())
			return false
		}

		if prop == nil {
			v.NewError(fmt.Sprintf("La propiedad %s no existe en %s", ids[1].GetText(), id), ctx.GetStart())
			return false
		}

		variable = prop
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

	if CheckIsPointer(variable.Value) {
		variable = variable.Value.(*Variable)
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
