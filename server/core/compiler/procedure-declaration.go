package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitFunctionDeclarationStatement(ctx *parser.FunctionDeclarationStatementContext) interface{} {

	id := ctx.ID().GetText()

	parameters := make([]Parameter, 0)

	if ctx.FunctionParameters() != nil {
		parameters = c.Visit(ctx.FunctionParameters()).([]Parameter)
	}

	fmt.Println("FUNCTION", id, parameters)
	return nil
}

func (c *Compiler) VisitFunctionParameters(ctx *parser.FunctionParametersContext) interface{} {
	// Return an array of V.IValue
	params := make([]Parameter, 0)

	for _, param := range ctx.AllFunctionParameter() {
		param := c.Visit(param).(Parameter)

		params = append(params, param)
	}

	return params
}

func (v *Compiler) VisitFunctionParameter(ctx *parser.FunctionParameterContext) interface{} {
	// Get the parameter name
	externalName := ctx.ID(0).GetText()
	internalName := externalName

	if len(ctx.AllID()) == 2 {
		externalName = ctx.ID(0).GetText()
		internalName = ctx.ID(1).GetText()
	}

	fmt.Println(externalName, internalName)
	return nil
}
