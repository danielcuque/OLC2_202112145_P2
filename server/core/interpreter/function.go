package interpreter

import (
	E "OLC2/core/error"
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

type IFunction interface {
	GetName() string
	GetParameters() []V.IValue
}

type Function struct {
	IsMutating     bool
	IsStructMethod bool
	Name           string
	ReturnDataType string
	Parameters     []Parameter
	Body           *parser.BlockContext
	ReturnValue    V.IValue
	Line           int
	Column         int
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

func (f *Function) GetParamsString() string {
	str := ""

	for i, param := range f.Parameters {
		str += param.String()

		if i != len(f.Parameters)-1 {
			str += ", "
		}
	}

	return str
}

func (f *Function) GetBody() *parser.BlockContext {
	return f.Body
}

func NewFunction(IsStructMethod, IsMutating bool, name string, returnDataType string, parameters []Parameter, body *parser.BlockContext, lc antlr.Token) *Function {
	return &Function{
		IsStructMethod: IsStructMethod,
		IsMutating:     IsMutating,
		Name:           name,
		ReturnDataType: returnDataType,
		Parameters:     parameters,
		Body:           body,
		Line:           lc.GetLine(),
		Column:         lc.GetColumn(),
	}
}

type Parameter struct {
	ExternalName string
	InternalName string
	Value        V.IValue
	IsINOUT      bool
}

func (p *Parameter) String() string {
	return p.InternalName
}

type Argument struct {
	Name      string
	Value     V.IValue
	IsPointer bool
}

// Visit methods

func (v *Visitor) VisitFunctionDeclarationStatement(ctx *parser.FunctionDeclarationStatementContext) interface{} {

	// Get the function name
	id := ctx.ID().GetText()

	// Verify if the function already exists
	function := v.Env.GetFunction(id)

	if function != nil {
		v.NewError(E.FunctionAlreadyExists, ctx.GetStart())
		return nil
	}

	// Get the function parameters
	parameters := make([]Parameter, 0)

	// Check if parameters are nil
	if ctx.FunctionParameters() != nil {
		parameters = v.Visit(ctx.FunctionParameters()).([]Parameter)
	}

	// Get the function return type
	returnType := V.NilType

	// Check if return type is nil
	if ctx.FunctionReturnType() != nil {
		RT, ok := v.Visit(ctx.FunctionReturnType()).(string)

		if !ok {
			v.NewError(E.InvalidReturnTypeFunction, ctx.GetStart())
			return nil
		}

		returnType = RT
	}

	isMutating := ctx.Kw_MUTATING() != nil

	if isMutating && !isDeclaringStruct {
		v.NewError("No se puede definir una función mutante fuera de un struct", ctx.GetStart())
		return nil
	}

	newFunction := NewFunction(isDeclaringStruct, isMutating, id, returnType, parameters, ctx.Block().(*parser.BlockContext), ctx.GetStart())

	if !isDeclaringStruct {
		v.Env.AddFunction(id, newFunction)
		return nil
	}

	return newFunction
}

func (v *Visitor) VisitFunctionParameters(ctx *parser.FunctionParametersContext) interface{} {
	// Return an array of V.IValue
	params := make([]Parameter, 0)

	for _, param := range ctx.AllFunctionParameter() {
		param, ok := v.Visit(param).(Parameter)

		if !ok {
			v.NewError(E.InvalidParameter, ctx.GetStart())
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

	dataType := V.NilType

	if ctx.VariableType() != nil {
		dataType = v.Visit(ctx.VariableType()).(string)
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

	if ctx.MatrixType() != nil {
		if ctx.MatrixType().GetText()[1] == '[' {
			param = NewMatrix(V.MatrixType, NewMatrixNode(v.Visit(ctx.MatrixType()).(string), []V.IValue{}))
		} else {
			param = NewVector(V.VectorType, []V.IValue{})
		}
	}

	return Parameter{
		ExternalName: externalName,
		InternalName: internalName,
		Value:        param,
		IsINOUT:      isINOUT,
	}
}

func (v *Visitor) VisitFunctionParameterCompoundSingle(ctx *parser.FunctionParameterCompoundSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFunctionReturnType(ctx *parser.FunctionReturnTypeContext) interface{} {
	returnType, ok := v.Visit(ctx.VariableType()).(string)

	if !ok {
		v.NewError(E.InvalidReturnTypeFunction, ctx.GetStart())
		return nil
	}

	return returnType
}

var IsStructMethodRunning bool
var IsStructMethodMutating bool
var objectStruct *ObjectV

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	id, ids := v.GetIds(ctx)

	// First check if is a native function, as print, etc
	function := GetInternalBuiltinFunctions(id)

	if function != nil {
		return function(v, ctx)
	}

	var fn *Function

	// There are two cases, the function is inside the environment or inside an object

	if len(ids) == 1 {
		// Assert that the function exists
		fnt := v.Env.GetFunction(id)

		if fnt == nil {
			objStruct := v.Env.GetStruct(id)

			if objStruct != nil {
				return v.HandleStructConstructor(ctx, objStruct)
			}

			v.NewError(E.FunctionNotFound, ctx.GetStart())
			return nil
		}
		fn = fnt
	} else {

		// Get the object
		object, okV := v.LookUpObject(id, ids, ctx.GetStart())

		if !okV {
			return nil
		}

		methodName := ids[1].GetText()

		if len(ids) > 2 {
			methodName = ids[len(ids)-1].GetText()
		}

		// Check if is a native function
		nativeFn := GetInternalBuiltinFunctions(methodName)

		if nativeFn != nil {
			return nativeFn(v, ctx)
		}

		objFn, okObjFn := object.GetMethod(methodName).(*Function)

		if !okObjFn || objFn == nil {
			v.NewError(fmt.Sprintf(`%s: '%s'`, E.MethodNotFound, methodName), ctx.GetStart())
			return nil
		}

		fn = objFn

		if fn.IsStructMethod {
			IsStructMethodRunning = true
			IsStructMethodMutating = fn.IsMutating
			objectStruct = object
		}
	}

	// Get the arguments
	args := v.GetArgs(ctx)

	// Verify if the number of parameters is the same

	if len(args) != len(fn.Parameters) {
		v.NewError(E.InvalidNumberOfParameters, ctx.GetStart())
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

	for i, arg := range args {
		var param Parameter

		if areArgsPositional {
			param = fn.Parameters[i]
		} else {
			param = fn.GetParameter(arg.Name)
		}

		if param.Value == nil {
			v.NewError(fmt.Sprintf("El parámetro '%s' no existe", param.InternalName), ctx.GetStart())
			v.Env.PopEnv()
			return false
		}

		if param.IsINOUT && !CheckIsPointer(arg.Value) {
			v.NewError(fmt.Sprintf("El parámetro '%s' es de tipo INOUT, por lo tanto debe ser un puntero", param.InternalName), ctx.GetStart())
			v.Env.PopEnv()
			return false
		}

		if !param.IsINOUT && CheckIsPointer(arg.Value) {
			v.NewError(fmt.Sprintf("El parámetro '%s' no es de tipo INOUT, por lo tanto no debe ser un puntero", param.InternalName), ctx.GetStart())
			v.Env.PopEnv()
			return false
		}

		if arg.Value.GetType() != param.Value.GetType() {
			v.NewError(fmt.Sprintf("El parámetro '%s' no es del tipo '%s'", param.InternalName, param.Value.GetType()), ctx.GetStart())
			v.Env.PopEnv()
			return false
		}

		newValue := arg.Value

		if !param.IsINOUT {
			newValue = arg.Value.Copy()
		}

		v.Env.AddVariable(param.InternalName, NewVariable(
			v,
			param.InternalName,
			false,
			newValue,
			param.Value.GetType(),
			ctx.GetStart(),
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
		v.NewError(E.InvalidReturnTypeFunction, ctx.GetStart())
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
			v.NewError(E.InvalidArgument, ctx.GetStart())
			return make([]Argument, 0)
		}

		args = append(args, Argument{Value: value, Name: "", IsPointer: CheckIsPointer(value)})
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

func (v *Visitor) GetIds(ctx *parser.FunctionCallContext) (string, []antlr.TerminalNode) {
	var id string
	var ids []antlr.TerminalNode

	if ctx.IdChain() != nil {
		ids = v.Visit(ctx.IdChain()).([]antlr.TerminalNode)
		id = ids[0].GetText()
		return id, ids
	}

	id = v.Visit(ctx.VariableType()).(string)

	return id, ids
}

func (v *Visitor) GetArgs(ctx *parser.FunctionCallContext) []Argument {
	args := make([]Argument, 0)

	if ctx.FunctionCallArguments() != nil {
		args = v.Visit(ctx.FunctionCallArguments()).([]Argument)
	}

	return args
}

func (v *Visitor) LookUpObject(id string, tokenProps []antlr.TerminalNode, lc antlr.Token) (*ObjectV, bool) {
	variable := v.Env.GetVariable(id)

	if variable == nil {
		v.NewError(E.ObjectNotFound, lc)
		return nil, false
	}

	object, ok := GetObject(variable).(*ObjectV)

	props := v.GetProps(tokenProps)

	if len(props) > 0 {
		object, ok = GetPropValue(variable, props).(*Variable).Value.(*ObjectV)

		if !ok {
			v.NewError(E.ObjectNotFound, lc)
			return nil, false
		}
	}

	if !ok {
		v.NewError(E.ObjectNotFound, lc)
		return nil, false
	}

	return object, true
}

func (v *Visitor) GetProps(ids []antlr.TerminalNode) []string {

	props := make([]string, 0)

	if len(ids) > 2 {
		props = make([]string, len(ids)-2)

		for i, id := range ids[1 : len(ids)-1] {
			props[i] = id.GetText()
		}
	}

	return props

}

func CheckIsPointer(value interface{}) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(&Variable{})
}
