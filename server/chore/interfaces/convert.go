package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"math"
	"strconv"
)

func Int(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	args := v.GetArgs(ctx)

	if len(args) != 1 {
		v.NewError(InvalidNumberOfParameters, ctx.GetStart())
		return V.NewNilValue(nil)
	}

	value := args[0].Value

	switch value.GetType() {
	case V.StringType:
		floatValue, err := strconv.ParseFloat(value.GetValue().(string), 64)
		if err == nil {
			return V.NewIntValue(int(math.Trunc(floatValue)))
		}

	case V.FloatType:
		return V.NewIntValue(int(math.Trunc(value.GetValue().(float64))))
	}

	return v.handleInvalidParameter(ctx)
}

func Float(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	args := v.GetArgs(ctx)

	if len(args) != 1 {
		v.NewError(InvalidNumberOfParameters, ctx.GetStart())
		return V.NewNilValue(nil)
	}

	value := args[0].Value

	switch value.GetType() {
	case V.StringType:
		floatValue, err := strconv.ParseFloat(value.GetValue().(string), 64)
		if err == nil {
			return V.NewFloatValue(floatValue)
		}

	case V.IntType:
		return V.NewFloatValue(float64(value.GetValue().(int)))
	}

	return v.handleInvalidParameter(ctx)
}

func String(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	args := v.GetArgs(ctx)

	if len(args) != 1 {
		v.NewError(InvalidNumberOfParameters, ctx.GetStart())
		return V.NewNilValue(nil)
	}

	value := args[0].Value

	switch value.GetType() {
	case V.IntType:
		return V.NewStringValue(strconv.Itoa(value.GetValue().(int)))

	case V.FloatType:
		return V.NewStringValue(strconv.FormatFloat(value.GetValue().(float64), 'f', -1, 64))

	case V.BooleanType:
		if value.GetValue().(bool) {
			return V.NewStringValue("true")
		} else {
			return V.NewStringValue("false")
		}
	}

	return v.handleInvalidParameter(ctx)
}

func (v *Visitor) handleInvalidParameter(ctx *parser.FunctionCallContext) V.IValue {
	v.NewError(InvalidParameter, ctx.GetStart())
	return V.NewNilValue(nil)
}
