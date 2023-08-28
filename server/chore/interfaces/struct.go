package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

var isDeclaringStruct bool

/*
- Si un atributo no posee un valor por defecto, entonces se debe establecer dicho
valor de forma obligatoria en el constructor, en caso contrario será un error ya que
no será posible que hayan atributos sin valor.
- Los atributos pueden ser mutables (var) o inmutables(let), si un atributo es inmutable
sólo se podrá asignarle un valor de alguna de las siguientes formas:
- En la declaración del atributo
- En el constructor
*/

func (v *Visitor) HandleStructConstructor(ctx *parser.FunctionCallContext, objectStruct *ObjectV) interface{} {
	// Here we can get struct parameters

	args := v.GetArgs(ctx)

	for _, arg := range args {
		if arg.Name == "" {
			v.NewError("El constructor de un struct no puede tener parámetros sin nombre", ctx.GetStart())
			return nil
		}
	}

	// Check if all parameters are declared
	for _, arg := range args {
		variable := objectStruct.Env.Variables[arg.Name]

		if variable == nil {
			v.NewError(fmt.Sprintf("El parámetro '%s' no existe en este struct", arg.Name), ctx.GetStart())
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
	for _, attr := range objectStruct.Env.Variables {
		if attr.Value.GetValue() == nil && attr.Type != V.NilType {
			v.NewError("El atributo "+attr.Name+" no tiene un valor por defecto", ctx.GetStart())
			return nil
		}
	}

	// Return a copy of the object
	return objectStruct.Copy()
}
