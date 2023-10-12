package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	if response == nil {
		fmt.Println("Error al declarar la variable")
		return nil
	}

	return c.DeclareValue(id, response)
}

func (c *Compiler) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	if response == nil {
		return nil
	}

	return c.DeclareValue(id, response)
}

func (c *Compiler) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	// The value will be nil
	id := ctx.ID().GetText()

	response := &ValueResponse{
		Type:        c.Visit(ctx.VariableType()).(TemporalCast),
		Value:       DefaultNil(),
		ContextType: LiteralType,
	}

	return c.DeclareValue(id, response)
}

func (c *Compiler) DeclareValue(id string, response *ValueResponse) *Value {

	if response.GetContextValue() == LiteralType {
		newTemporal := c.TAC.NewTemporal(response.GetType())

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("%s = %s;", newTemporal, response.GetValue()),
			},
			"",
		)

		response.SetValue(newTemporal)
	}

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("stack[(int)P] = %s;", response.GetValue()),
			"P = P + 1;",
		},
		fmt.Sprintf("Declaración de la variable '%s'", id),
	)

	newValue := NewValue(response.GetValue(), c.StackPointer.GetPointer(), response.GetType())

	c.StackPointer.AddPointer()

	c.Env.AddValue(id, newValue)

	return newValue
}
