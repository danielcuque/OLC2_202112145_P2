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
		newTemporal := c.TAC.NewTemporal(response.GetValue(), nil)
		c.TAC.AppendCode(
			[]string{
				fmt.Sprintf("%s = %s;", newTemporal, response.GetValue()),
			},
			"",
		)

		response.SetValue(newTemporal)
	}

	if response.GetContextValue() == LabelType {
		// This response have true and false labels

		result := c.TAC.NewTemporal("", BooleanTemporal)

		labelStack := response.GetValue().(*LabelStack)
		exitLabel := c.TAC.NewLabel("")

		c.TAC.AppendCode(
			[]string{
				c.DeclareLabels(labelStack.TrueLabel),
				fmt.Sprintf("%s = 1;", result),
				fmt.Sprintf("goto %s;", exitLabel),
				c.DeclareLabels(labelStack.FalseLabel),
				fmt.Sprintf("%s = 0;", result),
				exitLabel.Declare(),
			},
			"",
		)

		response.SetValue(result)
	}

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("stack[(int)P] = %s;", response.GetValue()),
			"P = P + 1;",
		},
		fmt.Sprintf("Declaración de la variable '%s'", id),
	)

	newValue := NewValue(response.GetValue(), c.StackPointer.GetPointer(), TemporalCast(response.GetType()))

	c.StackPointer.AddPointer()

	c.Env.AddValue(id, newValue)

	return newValue
}
