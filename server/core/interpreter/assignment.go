package interfaces

import (
	E "OLC2/core/error"
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) handleVariableAssignment(ctx antlr.ParserRuleContext, variable *Variable, value V.IValue, op string) interface{} {
	if variable == nil {
		v.NewError("No se puede asignar a un objeto nulo", ctx.GetStart())
		return nil
	}

	if variable.GetType() != value.GetType() {
		if variable.GetType() == V.FloatType && value.GetType() == V.IntType {
			value = V.NewFloatValue(float64(value.GetValue().(int)))
		} else {
			v.NewError(fmt.Sprintf("El tipo del objeto no coincide con el valor asignado, se esperaba %s y se obtuvo %s", variable.GetType(), value.GetType()), ctx.GetStart())
			return nil
		}
	}

	var newValue V.IValue
	switch op {
	case "=":
		newValue = value.Copy()
	case "+=":
		newValue = v.arithmeticOp(variable.Value, value, "+", ctx.GetStart()).(V.IValue)
	case "-=":
		newValue = v.arithmeticOp(variable.Value, value, "-", ctx.GetStart()).(V.IValue)
	default:
		v.NewError(fmt.Sprintf("No se puede aplicar el operador %s a %s", op, variable.GetType()), ctx.GetStart())
		return nil
	}

	variable.SetValue(newValue)
	return nil
}

func (v *Visitor) VisitVariableAssignment(ctx *parser.VariableAssignmentContext) interface{} {
	ids := v.Visit(ctx.IdChain()).([]antlr.TerminalNode)
	id := ids[0].GetText()

	if IsStructMethodRunning && !IsStructMethodMutating && id == "self" {
		v.NewError("No se puede modificar un struct desde un método que no sea mutante", ctx.GetStart())
		return nil
	}

	value, ok := v.Visit(ctx.Expr()).(V.IValue)
	if !ok {
		v.NewError(E.InvalidExpression, ctx.GetStart())
		return nil
	}

	variable := v.Env.GetVariable(id)

	if IsStructMethodRunning && id == "self" {
		variable = objectStruct.Env.Variables[ids[1].GetText()]
	}

	if variable == nil {
		v.NewError(fmt.Sprintf("La variable %s no existe", id), ctx.GetStart())
		return nil
	}

	if len(ids) > 1 && ids[0].GetText() != "self" {
		props := []string{ids[1].GetText()}

		if len(ids) > 2 {
			props = v.GetProps(ids)
		}

		prop, ok := GetPropValue(variable, props).(*Variable)

		if !ok {
			v.NewError(fmt.Sprintf("La variable %s no es un objeto", id), ctx.GetStart())
			return nil
		}

		if prop == nil {
			v.NewError(fmt.Sprintf("La propiedad %s no existe en %s", ids[1].GetText(), id), ctx.GetStart())
			return nil
		}

		variable = prop
	}

	if variable.IsConstant() {
		v.NewError(fmt.Sprintf("La variable %s es constante", id), ctx.GetStart())
		return nil
	}

	if CheckIsPointer(variable.Value) {
		variable = variable.Value.(*Variable)
	}

	op := ctx.GetOp().GetText()
	return v.handleVariableAssignment(ctx, variable, value, op)
}

func (v *Visitor) VisitVariableAssignmentObject(ctx *parser.VariableAssignmentObjectContext) interface{} {
	variable := v.Visit(ctx.ObjectChain()).(*Variable)

	value, ok := v.Visit(ctx.Expr()).(V.IValue)
	if !ok {
		v.NewError(E.InvalidExpression, ctx.GetStart())
		return nil
	}

	op := ctx.GetOp().GetText()
	return v.handleVariableAssignment(ctx, variable, value, op)
}
