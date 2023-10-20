package compiler

import (
	"OLC2/core/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	id, _ := c.GetIds(ctx)

	fnc := GetInternalBuiltinFunctions(id)

	if fnc != nil {
		return fnc(c, ctx)
	}

	procedure := c.TAC.GetProcedure(id)

	if procedure == nil {
		return nil
	}

	args := c.GetArgs(ctx)

	if len(args) > 0 {

		c.AllocateStack(len(args))

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
		value := c.Visit(arg).(*ValueResponse)

		// TODO: Check if value is a pointer
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

func (c *Compiler) GetIds(ctx *parser.FunctionCallContext) (string, []antlr.TerminalNode) {
	var id string
	var ids []antlr.TerminalNode

	if ctx.IdChain() != nil {
		ids = c.Visit(ctx.IdChain()).([]antlr.TerminalNode)
		id = ids[0].GetText()
		return id, ids
	}

	id = c.Visit(ctx.VariableType()).(string)

	return id, ids
}

func (c *Compiler) GetArgs(ctx *parser.FunctionCallContext) []*Argument {
	args := make([]*Argument, 0)

	if ctx.FunctionCallArguments() != nil {
		args = c.Visit(ctx.FunctionCallArguments()).([]*Argument)
	}

	return args
}
