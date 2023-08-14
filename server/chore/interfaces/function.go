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
	Parameters     []Parameter
	Body           *parser.BlockContext
}

type Parameter struct {
	Name  string
	Value V.IValue
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetParameters() []Parameter {
	return f.Parameters
}

func NewFunction(name string, returnDataType string, parameters []Parameter, body *parser.BlockContext) *Function {
	return &Function{
		Name:           name,
		ReturnDataType: returnDataType,
		Parameters:     parameters,
	}
}

// Visit methods

func (v *Visitor) VisitFunctionDeclarationStatement(ctx *parser.FunctionDeclarationStatementContext) interface{} {

	// Get the function name
	id := ctx.ID().GetText()

	// Verify if the function already exists
	_, ok := v.Scope.GetFunction(id).(*Function)

	if ok {
		v.NewError(FunctionAlreadyExistsError, ctx.GetStart())
		return nil
	}

	// Get the function parameters
	parameters := v.Visit(ctx.FunctionParameters()).([]Parameter)

	// Get the function return type
	returnType, ok := v.Visit(ctx.FunctionReturnType()).(string)

	if !ok {
		v.NewError("No se pudo obtener el tipo de valor de retorno", ctx.GetStart())
		return nil
	}

	v.Scope.AddFunction(id, NewFunction(id, returnType, parameters, ctx.Block().(*parser.BlockContext)))

	return nil
}

func (v *Visitor) VisitFunctionParameters(ctx *parser.FunctionParametersContext) interface{} {
	// Return an array of V.IValue
	params := make([]Parameter, 0)

	for _, param := range ctx.AllFunctionParameter() {
		param, ok := v.Visit(param).(Parameter)

		if !ok {
			v.NewError(InvalidParameterError, ctx.GetStart())
			// Return empty array
			return make([]Parameter, 0)
		}

		params = append(params, param)
	}

	return params
}

func (v *Visitor) VisitFunctionParameter(ctx *parser.FunctionParameterContext) interface{} {
	// Get the parameter name
	id := ctx.ID().GetText()

	// Get the parameter type
	dataType, ok := v.Visit(ctx.VariableType()).(string)

	if !ok {
		v.NewError(InvalidParameterTypeError, ctx.GetStart())
		return nil
	}

	var param V.IValue

	// Create the parameter
	switch dataType {
	case "Int":
		param = V.NewIntValue(0)
	case "Float":
		param = V.NewFloatValue(0.0)
	case "String":
		param = V.NewStringValue("")
	case "Bool":
		param = V.NewBooleanValue(false)
	case "Char":
		param = V.NewCharValue(' ')
	default:
		param = V.NewNilValue(nil)
	}

	return Parameter{
		Name:  id,
		Value: param,
	}
}

func (v *Visitor) VisitFunctionReturnType(ctx *parser.FunctionReturnTypeContext) interface{} {
	returnType, ok := v.Visit(ctx.VariableType()).(string)

	if !ok {
		v.NewError(InvalidReturnTypeDeclarationError, ctx.GetStart())
		return ""
	}

	return returnType
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
