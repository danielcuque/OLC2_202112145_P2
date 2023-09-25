package interfaces

import (
	E "OLC2/core/error"
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"
	"math"
	"strconv"
)

func Int(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	args := v.GetArgs(ctx)

	if len(args) != 1 {
		v.NewError(E.InvalidNumberOfParameters, ctx.GetStart())
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

	return v.handleInvalidParameter(value.GetType(), "Int", ctx)
}

func Float(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	args := v.GetArgs(ctx)

	if len(args) != 1 {
		v.NewError(E.InvalidNumberOfParameters, ctx.GetStart())
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

	return v.handleInvalidParameter(value.GetType(), "Float", ctx)
}

func String(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	args := v.GetArgs(ctx)

	if len(args) != 1 {
		v.NewError(E.InvalidNumberOfParameters, ctx.GetStart())
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

	return v.handleInvalidParameter(value.GetType(), "String", ctx)
}

func (v *Visitor) handleInvalidParameter(erroType, expectedType string, ctx *parser.FunctionCallContext) V.IValue {
	v.NewError(fmt.Sprintf("No se puede convertir un objeto de tipo %s a %s", erroType, expectedType), ctx.GetStart())
	return V.NewNilValue(nil)
}
