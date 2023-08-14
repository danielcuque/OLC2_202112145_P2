package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

type IFunction interface {
	GetName() string
	GetParameters() []V.IValue
}

type Function struct {
	Name           string
	ReturnDataType string
	Parameters     []V.IValue
}

type Arguments struct {
	Value V.IValue
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetParameters() []V.IValue {
	return f.Parameters
}

func NewFunction(name string, returnDataType string, parameters []V.IValue) *Function {
	return &Function{
		Name:           name,
		ReturnDataType: returnDataType,
		Parameters:     parameters,
	}
}

// Visit methods

func (v *Visitor) VisitFunctionDeclarationStatement(ctx *parser.FunctionDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionParameters(ctx *parser.FunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionParameter(ctx *parser.FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionReturnType(ctx *parser.FunctionReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	id := ctx.ID().GetText()

	function := GetInternalBuiltinFunctions(id)

	if function != nil {
		return function(v, ctx)
	}

	// Verify if the function exists
	_, ok := v.Scope.GetFunction(id).(*Function)

	if !ok {
		v.NewError(FunctionNotFoundError, ctx.GetStart())
		return nil
	}

	// Assert to array of V.IValue
	params := v.Visit(ctx.FunctionCallParameters()).([]V.IValue)

	fmt.Println(params)

	return nil
}

func (v *Visitor) VisitFunctionCallParameters(ctx *parser.FunctionCallParametersContext) interface{} {
	// Return an array of V.IValue
	params := make([]V.IValue, 0)

	for _, param := range ctx.AllExpr() {
		param, ok := v.Visit(param).(V.IValue)

		if !ok {
			v.NewError(InvalidParameterError, ctx.GetStart())
			// Return empty array
			return make([]V.IValue, 0)
		}

		params = append(params, param)
	}

	return params
}
