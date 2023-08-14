package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

// We extends the visitor to add the print function, create PrintContext and add it to the scope

func Print(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	params, ok := v.Visit(ctx.FunctionCallParameters()).([]IValue)

	if !ok {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return nil
	}

	var log string

	for _, param := range params {
		if !IsBaseType(param) {
			v.NewError("El parámetro no es de tipo primitivo", ctx.GetStart())
			return nil
		}

		log += fmt.Sprintf("%v ", param.GetValue())
	}

	v.NewLog(log)

	return nil

}
