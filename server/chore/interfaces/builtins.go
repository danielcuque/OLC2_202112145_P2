package interfaces

import (
	"OLC2/chore/parser"
)

var internalFunction map[string]func(v *Visitor, ctx *parser.FunctionCallContext) interface{}
var internalProperty map[string]func(v *Visitor, ctx *parser.CallPropertiesContext) interface{}

func GetInternalBuiltinFunctions(name string) func(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	if internalFunction == nil {
		internalFunction = map[string]func(v *Visitor, ctx *parser.FunctionCallContext) interface{}{
			"print":  Print,
			"Int":    Int,
			"Float":  Float,
			"String": String,
		}
	}
	return internalFunction[name]
}

func GetInternalBuiltinProperties(name string) func(v *Visitor, ctx *parser.CallPropertiesContext) interface{} {
	if internalFunction == nil {
		internalProperty = map[string]func(v *Visitor, ctx *parser.CallPropertiesContext) interface{}{
			"count": CountBuiltin,
		}
	}
	return internalProperty[name]
}
