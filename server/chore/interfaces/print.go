package interfaces

import (
	"fmt"

	"OLC2/chore/parser"
	V "OLC2/chore/values"
)

// We extends the visitor to add the print function, create PrintContext and add it to the scope

func Print(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	args, ok := v.Visit(ctx.FunctionCallArguments()).([]Argument)

	if !ok {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return nil
	}

	var log string

	for _, arg := range args {
		if !V.IsBaseType(arg.Value) {
			v.NewError("El parámetro no es de tipo primitivo", ctx.GetStart())
			return nil
		}

		log += fmt.Sprintf("%v ", arg.Value.String())
	}

	v.NewLog(log)

	return nil

}
