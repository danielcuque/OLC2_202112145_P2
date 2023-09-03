package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
	"strings"
)

var isDeclaringStruct bool

func (v *Visitor) HandleStructConstructor(ctx *parser.FunctionCallContext, objStruct *ObjectV) *ObjectV {
	// Here we can get struct parameters

	objStruct = objStruct.Copy().(*ObjectV)

	args := v.GetArgs(ctx)

	for _, arg := range args {
		if arg.Name == "" {
			v.NewError("El constructor de un struct no puede tener parámetros sin nombre", ctx.GetStart())
			return nil
		}
	}

	// Check if all parameters are declared
	for _, arg := range args {
		variable := objStruct.Env.Variables[arg.Name]

		if variable == nil {
			v.NewError(fmt.Sprintf("El parámetro '%s' no existe en el struct '%s'", arg.Name, objStruct.Type), ctx.GetStart())
			return nil
		}

		if variable.GetType() != arg.Value.GetType() {
			v.NewError(fmt.Sprintf("El tipo del parámetro '%s' no es válido", arg.Name), ctx.GetStart())
			return nil
		}

		// Set value to variable

		variable.SetValue(arg.Value)
	}

	// Check if all attributes are declared
	for _, attr := range objStruct.Env.Variables {
		if attr.Value.GetValue() == nil && attr.Type != V.NilType {
			v.NewError("El atributo "+attr.Name+" no tiene un valor por defecto", ctx.GetStart())
			return nil
		}
	}

	// Return a copy of the object
	return objStruct
}

func (v *Visitor) HandleVisitIdStruct(ctx *parser.IdExprContext) interface{} {
	if objectStruct == nil {
		return nil
	}

	expr := strings.Split(ctx.IdChain().GetText(), ".")[1:]

	variable := objectStruct.Env.Variables[expr[0]]

	if variable == nil {
		v.NewError(fmt.Sprintf("La variable %s no existe", expr[0]), ctx.GetStart())
		return nil
	}

	if len(expr) == 1 {
		return variable.Value
	}

	params := expr[1:]

	prop, okP := GetPropValue(variable, params).(*Variable)

	if prop == nil {
		v.NewError(fmt.Sprint("La propiedad ", params[len(params)-1], " no existe"), ctx.GetStart())
		return V.NewNilValue(nil)
	}

	if !okP {
		v.NewError(fmt.Sprint("La propiedad ", params[len(params)-1], " no existe"), ctx.GetStart())
		return V.NewNilValue(nil)
	}

	return prop.Value
}
