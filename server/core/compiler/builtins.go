package compiler

import "OLC2/core/parser"

var internalFunction map[string]func(c *Compiler, ctx *parser.FunctionCallContext) interface{}

func GetInternalBuiltinFunctions(name string) func(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	if internalFunction == nil {
		internalFunction = map[string]func(c *Compiler, ctx *parser.FunctionCallContext) interface{}{
			"print":      Print,
			"Int":        Int,
			"Float":      Float,
			"String":     String,
			"append":     Append,
			"removeLast": RemoveLast,
			"remove":     Remove,
		}
	}
	return internalFunction[name]
}
