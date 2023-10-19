package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitFunctionDeclarationStatement(ctx *parser.FunctionDeclarationStatementContext) interface{} {

	id := ctx.ID().GetText()

	newProcedure := NewProcedure(id)

	parameters := make([]*Parameter, 0)

	if ctx.FunctionParameters() != nil {
		parameters = c.Visit(ctx.FunctionParameters()).([]*Parameter)
	}

	newProcedure.AddArguments(parameters)

	newProcedure.SetBasePointer(c.TAC.NewTemporal(IntTemporal))

	c.TAC.SetProcedure(newProcedure)
	c.Visit(ctx.Block())
	c.TAC.UnsetProcedure()

	// fmt.Println("FUNCTION", id, newProcedure)
	return nil
}

func (c *Compiler) VisitFunctionParameters(ctx *parser.FunctionParametersContext) interface{} {
	// Return an array of V.IValue
	params := make([]*Parameter, 0)

	for _, param := range ctx.AllFunctionParameter() {
		param := c.Visit(param).(*Parameter)

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

	isRef := ctx.Kw_INOUT() != nil

	return &Parameter{
		Temporal:     nil,
		ExternalName: externalName,
		InternalName: internalName,
		IsReference:  isRef,
	}
}
