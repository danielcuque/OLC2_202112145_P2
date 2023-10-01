package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	return c.DeclareValue(id, response)
}

func (c *Compiler) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	return c.DeclareValue(id, response)
}

func (c *Compiler) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	// The value will be nil
	id := ctx.ID().GetText()

	response := &ValueResponse{
		Type:        ctx.VariableType().GetText(),
		Value:       DefaultNil(),
		ContextType: LiteralType,
	}

	return c.DeclareValue(id, response)
}

func (c *Compiler) DeclareValue(id string, response *ValueResponse) *Value {

	if response.GetContextValue() == LiteralType {

		c.TAC.NewTemporal(response.GetValue(), nil)
		c.TAC.AppendCode(
			[]string{
				fmt.Sprintf("t%d = %s;", c.TAC.TemporalQuantity(), response.GetValue()),
			},
			"",
		)

		response.SetValue(fmt.Sprintf("t%d", c.TAC.TemporalQuantity()))
	}

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("stack[(int)P] = %s;", response.GetValue()),
			fmt.Sprintf("P = P + 1;"),
		},
		fmt.Sprintf("Declaración de la variable '%s'", id),
	)

	newValue := NewValue(response.GetValue(), c.StackPointer.GetPointer(), TemporalCast(response.GetType()))

	c.StackPointer.AddPointer()

	c.Env.AddValue(id, newValue)

	return newValue
}
