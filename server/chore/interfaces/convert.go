package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"math"
	"strconv"
)

func Int(v *Visitor, ctx *parser.FunctionCallContext) interface{} {

	args, ok := v.Visit(ctx.FunctionCallArguments()).([]Argument)

	if !ok {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return nil
	}

	if len(args) != 1 {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return V.NewNilValue(nil)

	}

	value := args[0].Value

	switch value.GetType() {
	case V.StringType:
		stringValue := value.GetValue().(string)

		// Intentar convertir a float primero y luego a int
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err == nil {
			return V.NewIntValue(int(math.Trunc(floatValue)))
		} else {
			v.NewError(InvalidParameterError, ctx.GetStart())
			return V.NewNilValue(nil)
		}

	case V.FloatType:
		return V.NewIntValue(int(math.Trunc(value.GetValue().(float64))))

	default:
		v.NewError(InvalidParameterError, ctx.GetStart())
		return V.NewNilValue(nil)
	}
}

func Float(v *Visitor, ctx *parser.FunctionCallContext) interface{} {

	args, ok := v.Visit(ctx.FunctionCallArguments()).([]Argument)

	if !ok {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return nil
	}

	if len(args) != 1 {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return V.NewNilValue(nil)

	}

	value := args[0].Value

	switch value.GetType() {
	case V.StringType:
		stringValue := value.GetValue().(string)

		// Intentar convertir a float primero y luego a int
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err == nil {
			return V.NewFloatValue(floatValue)
		} else {
			v.NewError(InvalidParameterError, ctx.GetStart())
			return V.NewNilValue(nil)
		}

	case V.IntType:
		return V.NewFloatValue(float64(value.GetValue().(int)))

	default:
		v.NewError(InvalidParameterError, ctx.GetStart())
		return V.NewNilValue(nil)
	}
}

func String(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	/*
		Esta función es la contraparte de las dos anteriores, es decir, toma como parámetro un valor numérico y retorna una cadena de tipo String. Además sí recibe un valor Bool lo convierte en "true" o "false". Para valores tipo Float la cantidad de números después del punto decimal queda a discreción del estudiante.
	*/

	args, ok := v.Visit(ctx.FunctionCallArguments()).([]Argument)

	if !ok {
		v.NewError(InvalidParameterError, ctx.GetStart())
		return nil
	}

	if len(args) != 1 {
		v.NewError(InvalidParameterError, ctx.GetStart())
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

	default:
		v.NewError(InvalidParameterError, ctx.GetStart())
		return V.NewNilValue(nil)
	}
}
