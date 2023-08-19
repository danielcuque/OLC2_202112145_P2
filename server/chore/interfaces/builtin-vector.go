package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
)

func Append(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	// This function just appends a value to a vector

	id, _ := v.GetIds(ctx)

	// Get the vector
	baseVar, okV := v.Env.GetVariable(id).(*Variable)

	if !okV {
		v.NewError(ObjectNotFound, ctx.GetStart())
		return nil
	}

	vector, ok := baseVar.Value.(*ObjectV)

	if !ok || vector.GetType() != V.VectorType {
		v.NewError(InvalidParameter, ctx.GetStart())
		return nil
	}

	args := v.GetArgs(ctx)

	if len(args) > 1 {
		v.NewError(InvalidNumberOfParameters, ctx.GetStart())
		return nil
	}

	return V.NewNilValue(nil)
}
