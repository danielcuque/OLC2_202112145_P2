package interfaces

import (
	"fmt"

	"OLC2/chore/parser"
	V "OLC2/chore/values"
)

// We extends the visitor to add the print function, create PrintContext and add it to the scope

func Print(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	params, ok := v.Visit(ctx.FunctionCallParameters()).([]V.IValue)

	if !ok {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return nil
	}

	var log string

	for _, param := range params {
		if !V.IsBaseType(param) {
			v.NewError("El parámetro no es de tipo primitivo", ctx.GetStart())
			return nil
		}

		log += fmt.Sprintf("%v ", param.GetValue())
	}

	v.NewLog(log)

	return nil

}
