package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {

	id := ctx.ID().GetText()
	evalResponse := c.Visit(ctx.Expr())

	if evalResponse == nil {
		fmt.Println("Error al declarar la variable", id)
		return nil
	}

	response := evalResponse.(*ValueResponse)

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

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	valueContent, valueToAssign := response.GetValue(), response.GetValue()

	if response.Type == StructTemporal {
		structValue := response.GetValue().([]interface{})
		valueContent = structValue[0]
		valueToAssign = structValue[1]
	}

	value.SetData(response.GetType(), valueContent)

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[(int)%v] = %v;", c.TAC.GetValueAddress(value), valueToAssign),
		fmt.Sprintf("Declaración de la variable '%s'", id),
	)

	return nil
}
