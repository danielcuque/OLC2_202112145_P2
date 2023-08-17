package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
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
	ReturnValue    V.IValue
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetParameters() []Parameter {
	return f.Parameters
}

func (f *Function) GetParameter(name string) Parameter {
	for _, param := range f.Parameters {
		if param.ExternalName == name {
			return param
		}
	}

	return Parameter{}
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
	ExternalName string
	InternalName string
	Value        V.IValue
	IsINOUT      bool
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
	_, ok := v.Env.GetFunction(id).(*Function)

	if ok {
		v.NewError(FunctionAlreadyExists, ctx.GetStart())
		return nil
	}

	// Get the function parameters
	var parameters []Parameter

	// Check if parameters are nil
	if ctx.FunctionParameters() != nil {
		parameters = v.Visit(ctx.FunctionParameters()).([]Parameter)
	} else {
		parameters = make([]Parameter, 0)
	}

	// Get the function return type
	returnType := V.NilType

	// Check if return type is nil
	if ctx.FunctionReturnType() != nil {
		returnType = v.Visit(ctx.FunctionReturnType()).(string)
	}

	// if !ok {
	// 	v.NewError("No se pudo obtener el tipo de valor de retorno", ctx.GetStart())
	// 	return nil
	// }

	v.Env.AddFunction(id, NewFunction(id, returnType, parameters, ctx.Block().(*parser.BlockContext)))

	return nil
}

func (v *Visitor) VisitFunctionParameters(ctx *parser.FunctionParametersContext) interface{} {
	// Return an array of V.IValue
	params := make([]Parameter, 0)

	for _, param := range ctx.AllFunctionParameter() {
		param, ok := v.Visit(param).(Parameter)

		if !ok {
			v.NewError(InvalidParameter, ctx.GetStart())
			// Return empty array
			return make([]Parameter, 0)
		}

		params = append(params, param)
	}

	return params
}

// There are three types of parameters:

/*
	1. External name and internal name
	2. Only external name
	3. Only internal name
*/

func (v *Visitor) VisitFunctionParameter(ctx *parser.FunctionParameterContext) interface{} {
	// Get the parameter name
	var externalName string
	var internalName string

	if len(ctx.AllID()) == 2 {
		externalName = ctx.ID(0).GetText()
		internalName = ctx.ID(1).GetText()
	} else {
		externalName = ctx.ID(0).GetText()
		internalName = externalName
	}

	isINOUT := ctx.Kw_INOUT() != nil

	// Get the parameter type
	dataType, ok := v.Visit(ctx.VariableType()).(string)

	if !ok {
		v.NewError(InvalidParameterType, ctx.GetStart())
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
		ExternalName: externalName,
		InternalName: internalName,
		Value:        param,
		IsINOUT:      isINOUT,
	}
}

func (v *Visitor) VisitFunctionReturnType(ctx *parser.FunctionReturnTypeContext) interface{} {
	returnType, ok := v.Visit(ctx.VariableType()).(string)

	if !ok {
		v.NewError(InvalidReturnTypeFunction, ctx.GetStart())
		return nil
	}

	return returnType
}

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	var id string
	var ids []antlr.TerminalNode

	if ctx.IdChain() != nil {
		ids = v.Visit(ctx.IdChain()).([]antlr.TerminalNode)
		id = ids[0].GetText()
	} else {
		id = v.Visit(ctx.VariableType()).(string)
	}

	function := GetInternalBuiltinFunctions(id)

	if function != nil {
		return function(v, ctx)
	}

	var fn *Function

	// There are two cases, the function is inside the environment or inside an object

	if len(ids) == 1 {
		// Assert that the function exists
		fnt, ok := v.Env.GetFunction(id).(*Function)

		if !ok {
			v.NewError(FunctionNotFound, ctx.GetStart())
			return nil
		}
		fn = fnt
	} else {
		// Get the baseVar
		baseVar, okV := v.Env.GetVariable(id).(*Variable)

		if !okV {
			v.NewError(ObjectNotFound, ctx.GetStart())
			return nil
		}

		var props []string
		var methodName string
		var object *ObjectV

		// Check if there are not properties

		// If there are just two ids, means that second id is the method
		if len(ids) == 2 {
			// Get object
			obj, okP := baseVar.Value.(*ObjectV)

			if !okP {
				v.NewError(ObjectNotFound, ctx.GetStart())
				return nil
			}

			object = obj
			methodName = ids[1].GetText()
		} else {
			for _, id := range ids[1 : len(ids)-1] {
				props = append(props, id.GetText())
			}
			methodName = ids[len(ids)-1].GetText()

			obj, okP := GetPropValue(baseVar, props).(*Variable).Value.(*ObjectV)

			if !okP {
				v.NewError(ObjectNotFound, ctx.GetStart())
				return nil
			}

			object = obj
		}

		// Check if is a native function
		nativeFn := GetInternalBuiltinFunctions(methodName)

		if nativeFn != nil {
			return nativeFn(v, ctx)
		}

		objFn, okObjFn := object.GetMethod(methodName).(*Function)

		if !okObjFn {
			v.NewError(MethodNotFound, ctx.GetStart())
			return nil
		}

		fn = objFn
	}

	// Get the arguments
	args := make([]Argument, 0)

	// Check if arguments are nil
	if ctx.FunctionCallArguments() != nil {
		args = v.Visit(ctx.FunctionCallArguments()).([]Argument)
	}

	// Verify if the number of parameters is the same
	if len(args) != len(fn.Parameters) {
		v.NewError(InvalidNumberOfParameters, ctx.GetStart())
		return nil
	}

	/*
		There are three verifications to do:
		1. Check if is positional, for this, we need to check if the arguments have name and and if parameter external name is different to '_' (underscore)
		2. Check if the function contains external name, for this, we need to check if the arguments have name
	*/

	areArgsPositional := true

	for _, arg := range args {
		if arg.Name != "" {
			areArgsPositional = false
			break
		}
	}

	// Check if parameters have a different external name than '_', then throw an error
	for _, param := range fn.Parameters {
		if areArgsPositional && param.ExternalName != "_" {
			v.NewError("La función no acepta parámetros posicionales", ctx.GetStart())
			return nil
		}
	}

	// Create a new scope
	fnScope := v.Env.PushEnv(id)
	v.Stack.Push(NewStackItem(
		id,
		V.NewNilValue(nil),
		[]StackItemType{ReturnType},
	))

	// Now, initialize the parameters into a new Scope
	line, column, scope := GetVariableAttr(v, ctx.GetStart())

	for i, arg := range args {
		var param Parameter

		if areArgsPositional {
			param = fn.Parameters[i]
		} else {
			param = fn.GetParameter(arg.Name)
		}

		if param.Value == nil {
			v.NewError(fmt.Sprintf("El parámetro '%s' no existe", param.InternalName), ctx.GetStart())
			return false
		}

		if arg.Value.GetType() != param.Value.GetType() {
			v.NewError(fmt.Sprintf("El parámetro '%s' no es del tipo '%s'", param.InternalName, param.Value.GetType()), ctx.GetStart())
			return false
		}

		v.Env.AddVariable(param.InternalName, NewVariable(
			param.InternalName,
			false,
			arg.Value,
			arg.Value.GetType(),
			line,
			column,
			scope,
		))
	}

	// Execute the function
	v.ExecuteFunctionBody(fn, ctx, fnScope)

	var returnValue V.IValue

	if fn.ReturnValue != nil {
		returnValue = fn.ReturnValue
	} else {
		returnValue = V.NewNilValue(nil)
	}

	// Return scope to root scope
	v.Env.PopEnv()

	if returnValue.GetType() != fn.ReturnDataType {
		v.NewError(InvalidReturnTypeFunction, ctx.GetStart())
	}

	return returnValue
}

func (v *Visitor) ExecuteFunctionBody(fn *Function, ctx *parser.FunctionCallContext, calledScope *EnvNode) {
	// Create a new scope

	defer func() {
		v.Stack.Reset()
		v.Env.Current = calledScope

		peek, ok := recover().(*StackItem)

		if !ok {
			return
		}

		if peek.Trigger != ReturnType {
			v.NewError("No se pudo obtener un valor de retorno", ctx.GetStart())
			return
		}

		fn.ReturnValue = peek.Value
	}()

	// Execute the function body
	v.Visit(fn.Body)
}

func (v *Visitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	// Return an array of Argument
	args := make([]Argument, 0)

	for _, expr := range ctx.AllExpr() {
		value, ok := v.Visit(expr).(V.IValue)

		if !ok {
			v.NewError(InvalidArgument, ctx.GetStart())
			return nil
		}

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

func (v *Visitor) VisitIDChain(ctx *parser.IDChainContext) interface{} {
	return ctx.AllID()
}
