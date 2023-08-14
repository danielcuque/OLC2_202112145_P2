package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

type IFunction interface {
	GetName() string
	GetParameters() []IValue
	Evaluate() IValue
}

type Function struct {
	Name       string
	Parameters []IValue
}

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

	// Verify if the function exists
	_, ok := v.Scope.GetFunction(id).(*Function)

	if !ok {
		v.NewError("La funcion no existe en este ambiente", ctx.GetStart())
		return nil
	}

	// Assert to array of IValue
	params, ok := v.Visit(ctx.FunctionCallParameters()).([]IValue)

	if !ok {
		v.NewError("No se pudieron obtener los parametros de la funcion", ctx.GetStart())
		return nil
	}

	fmt.Println(params)

	return nil
}

func (v *Visitor) VisitFunctionCallParameters(ctx *parser.FunctionCallParametersContext) interface{} {
	// Return an array of IValue
	params := make([]IValue, 0)

	for _, param := range ctx.AllExpr() {
		param, ok := v.Visit(param).(IValue)

		if !ok {
			v.NewError("No se pudo obtener el parametro", ctx.GetStart())
			// Return empty array
			return make([]IValue, 0)
		}

		params = append(params, param)
	}

	return params
}
