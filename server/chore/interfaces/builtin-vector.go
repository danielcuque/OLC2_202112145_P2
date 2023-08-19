package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
)

func Append(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	// This function just appends a value to a vector

	id, _ := v.GetIds(ctx)

	// Get the vector
	vectorObj, _ := v.LookUpObject(id, nil, ctx)

	if vectorObj.GetType() != V.VectorType {
		v.NewError(InvalidParameter, ctx.GetStart())
		return nil
	}

	args := v.GetArgs(ctx)

	if len(args) > 1 {
		v.NewError(InvalidNumberOfParameters, ctx.GetStart())
		return nil
	}

	// Change object props as isEmpty and count

	vector := vectorObj.GetValue().(*VectorV)

	vector.Append(args[0].Value)

	// Change props
	vectorObj.SetPropValue("isEmpty", V.NewBooleanValue(vector.IsEmpty()))
	vectorObj.SetPropValue("count", V.NewIntValue(vector.Count()))

	return V.NewNilValue(nil)
}
