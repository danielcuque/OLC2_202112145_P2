package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
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

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetParameters() []Parameter {
	return f.Parameters
}

func (f *Function) GetParameter(name string) *Parameter {
	for _, param := range f.Parameters {
		if param.Name == name {
			return &param
		}
	}

	return nil
}

func (f *Function) GetBody() *parser.BlockContext {
	return f.Body
}

func NewFunction(name string, returnDataType string, parameters []Parameter, body *parser.BlockContext) *Function {
	return &Function{
		Name:           name,
		ReturnDataType: returnDataType,
		Parameters:     parameters,
		Body:           body,
	}
}

type Parameter struct {
	Name  string
	Value V.IValue
}

type Argument struct {
	Name  string
	Value V.IValue
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
		return nil
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
	fn, ok := v.Scope.GetFunction(id).(*Function)

	if !ok {
		v.NewError(FunctionNotFoundError, ctx.GetStart())
		return nil
	}

	// Get the arguments
	args := v.Visit(ctx.FunctionCallArguments()).([]Argument)

	// Verify if the number of parameters is the same
	if len(args) != len(fn.Parameters) {
		v.NewError(InvalidNumberOfParametersError, ctx.GetStart())
		return nil
	}

	isAllPositional := true

	// Check if arguments have name, if not, is a positional argument
	for _, arg := range args {
		if arg.Name == "" {
			continue
		} else {
			// Check if the argument name exists in the function parameters
			exists := false
			for _, param := range fn.Parameters {
				if param.Name == arg.Name {
					exists = true
					break
				}
			}

			if !exists {
				v.NewError(InvalidArgumentNameError, ctx.GetStart())
				return nil
			}

			isAllPositional = false
		}
	}

	// Create a new scope
	fnScope := v.Scope.PushScope(id)

	if isAllPositional {
		// Add the arguments to the scope by index
		for i, arg := range args {
			var param = fn.Parameters[i]
			fnScope.AddVariable(param.Name, NewVariable(
				param.Name,
				false,
				arg.Value,
				arg.Value.GetType(),
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				v.Scope.GetCurrentScope(),
			))
		}
	} else {
		// Add the arguments to the scope
		for _, arg := range args {
			var param = fn.GetParameter(arg.Name)
			fnScope.AddVariable(param.Name, NewVariable(
				param.Name,
				false,
				arg.Value,
				arg.Value.GetType(),
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				v.Scope.GetCurrentScope(),
			))
		}
	}

	// Execute the function
	body := fn.GetBody()

	if body == nil {
		v.NewError(InvalidFunctionBodyError, ctx.GetStart())
		v.Scope.PopScope()
		return nil
	}

	v.Visit(body)

	// Pop the scope
	v.Scope.PopScope()

	// if retunrValue.GetType() != fn.ReturnDataType {
	// 	v.NewError(InvalidReturnTypeFunctionError, ctx.GetStart())
	// }

	return nil
}

func (v *Visitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	// Return an array of Argument
	args := make([]Argument, 0)

	for _, expr := range ctx.AllExpr() {
		value := v.Visit(expr).(V.IValue)
		args = append(args, Argument{Value: value, Name: ""})
	}

	return args
}

func (v *Visitor) VisitNamedArguments(ctx *parser.NamedArgumentsContext) interface{} {
	// Return an array of Argument
	args := make([]Argument, 0)

	for i, expr := range ctx.AllExpr() {
		value := v.Visit(expr).(V.IValue)
		name := ctx.ID(i).GetText()
		args = append(args, Argument{Name: name, Value: value})
	}

	return args
}
