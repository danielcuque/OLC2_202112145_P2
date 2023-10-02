package compiler

import (
	"OLC2/core/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	id, args := c.GetIds(ctx)

	fnc := GetInternalBuiltinFunctions(id)

	if fnc != nil {
		return fnc(c, ctx)
	}

	fmt.Println("CALL", id, args)

	return nil
}

func (c *Compiler) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	// Unamed arguments
	args := make([]Argument, 0)

	for _, arg := range ctx.AllExpr() {
		value := c.Visit(arg).(*ValueResponse)

		// TODO: Check if value is a pointer
		args = append(args, Argument{
			Value:     value,
			Name:      "",
			IsPointer: false,
		})
	}

	return args
}

// Utils

func (c *Compiler) GetIds(ctx *parser.FunctionCallContext) (string, []antlr.TerminalNode) {
	var id string
	var ids []antlr.TerminalNode

	if ctx.IdChain() != nil {
		ids = c.Visit(ctx.IdChain()).([]antlr.TerminalNode)
		id = ids[0].GetText()
		return id, ids
	}

	id = c.Visit(ctx.VariableType()).(string)

	return id, ids
}

func (c *Compiler) GetArgs(ctx *parser.FunctionCallContext) []Argument {
	args := make([]Argument, 0)

	if ctx.FunctionCallArguments() != nil {
		args = c.Visit(ctx.FunctionCallArguments()).([]Argument)
	}

	return args
}
