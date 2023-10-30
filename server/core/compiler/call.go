package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	_, props := c.GetIds(ctx)

	functionName := props[len(props)-1]

	fnc := GetInternalBuiltinFunctions(functionName)

	if fnc != nil {
		return fnc(c, ctx)
	}

	if c.Env.GetValue(functionName) != nil {
		return c.HandleStructInstance(ctx)
	}

	procedure := c.TAC.GetProcedure(functionName)

	if procedure == nil {
		return nil
	}

	args := c.GetArgs(ctx)

	c.AllocateStack(procedure.Env.GetCountValues())

	if len(args) > 0 {

		procedureEnv := procedure.Env

		arePositional := args[0].Name == ""

		// Traverse array of arguments and match with procedure parameters
		for i, arg := range args {

			parameter := procedure.Parameters[i]

			if !arePositional {
				parameter = procedure.GetParameter(arg.Name)
			}

			if parameter == nil {
				fmt.Println("Parameter not found", arg.Name)
				continue
			}

			parameterName := parameter.InternalName

			parameterValue := procedureEnv.GetValue(parameterName)

			if parameterValue == nil {
				fmt.Print("Parameter value not found", parameterName)
				continue
			}

			parameterAddress := c.TAC.GetValueAddress(parameterValue)

			c.TAC.AppendInstruction(
				fmt.Sprintf(
					"stack[(int)%s]= %s;",
					parameterAddress,
					arg.Value.GetValue(),
				),
				"",
			)
		}
	}

	c.TAC.AppendInstruction(
		procedure.Call(),
		"",
	)

	return &ValueResponse{
		Type:        procedure.ReturnTemporal.Type,
		Value:       procedure.ReturnTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	// Unamed arguments
	args := make([]*Argument, 0)

	for _, arg := range ctx.AllExpr() {
		evalValue := c.Visit(arg)

		if evalValue == nil {
			return nil
		}

		value := evalValue.(*ValueResponse)

		args = append(args, &Argument{
			Value:     value,
			Name:      "",
			IsPointer: false,
		})
	}

	return args
}

func (c *Compiler) VisitNamedArguments(ctx *parser.NamedArgumentsContext) interface{} {
	// Return an array of Argument
	args := make([]*Argument, 0)

	for i, expr := range ctx.AllExpr() {
		value := c.Visit(expr).(*ValueResponse)
		name := ctx.ID(i).GetText()
		args = append(args, &Argument{Name: name, Value: value})
	}

	return args
}

// Utils

func (c *Compiler) GetIds(ctx *parser.FunctionCallContext) (string, []string) {
	id := ""
	ids := make([]string, 0)

	if ctx.IdChain() != nil {
		propId, props := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))
		id = propId
		ids = props
		return id, ids
	}

	id = ctx.VariableType().GetText()

	ids = append(ids, id)

	return id, ids
}

func (c *Compiler) GetArgs(ctx *parser.FunctionCallContext) []*Argument {
	args := make([]*Argument, 0)

	if ctx.FunctionCallArguments() != nil {
		evalArgs := c.Visit(ctx.FunctionCallArguments())

		if evalArgs == nil {
			return args
		}

		args = evalArgs.([]*Argument)
	}

	return args
}
