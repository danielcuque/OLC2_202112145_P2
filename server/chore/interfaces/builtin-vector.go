package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"

	"github.com/antlr4-go/antlr/v4"
)

func Append(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	// This function just appends a value to a vector
	var id string
	var ids []antlr.TerminalNode

	if ctx.IdChain() != nil {
		ids = v.Visit(ctx.IdChain()).([]antlr.TerminalNode)
		id = ids[0].GetText()
	} else {
		id = v.Visit(ctx.VariableType()).(string)
	}

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

	// Get the value to append

	args, ok := v.Visit(ctx.FunctionCallArguments()).([]Argument)

	if !ok {
		v.NewError(InvalidParameter, ctx.GetStart())
		return nil
	}

	if len(args) > 1 {
		v.NewError(InvalidNumberOfParameters, ctx.GetStart())
		return nil
	}

	return V.NewNilValue(nil)
}
