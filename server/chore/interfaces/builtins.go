package interfaces

import (
	"OLC2/chore/parser"
)

var internalFunction map[string]func(v *Visitor, ctx *parser.FunctionCallContext) interface{}

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
