package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitFunctionDeclarationStatement(ctx *parser.FunctionDeclarationStatementContext) interface{} {

	procedureName := ctx.ID().GetText()
	params := make([]*Parameter, 0)

	if ctx.FunctionParameters() != nil {
		params = c.Visit(ctx.FunctionParameters()).([]*Parameter)
	}

	statementsBlock := ctx.Block()

	staticVisitor := NewStaticVisitor(true, len(params), c.TAC)

	staticVisitor.Visit(statementsBlock)
	envFunction := staticVisitor.Env

	c.Env.Root.AppendChild(envFunction.Root)

	newProcedure := NewProcedure(procedureName)
	newProcedure.Env = envFunction

	newProcedure.AddParameters(params)

	for i, param := range params {
		newValue := envFunction.AddValue(param.InternalName, NewSimpleValue(i))
		newValue.IsRelative = true
	}

	newProcedure.SetReturnProps(
		c.TAC.NewTemporal(IntTemporal),
		c.NewLabelFlow("return", []LabelFlowType{ReturnLabel}),
	)

	c.TAC.AddProcedure(newProcedure)

	return nil
}

func (c *Compiler) VisitFunctionParameters(ctx *parser.FunctionParametersContext) interface{} {
	params := make([]*Parameter, 0)

	for _, param := range ctx.AllFunctionParameter() {
		param := c.Visit(param).(*Parameter)

		params = append(params, param)
	}

	return params
}

func (v *Compiler) VisitFunctionParameter(ctx *parser.FunctionParameterContext) interface{} {
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
