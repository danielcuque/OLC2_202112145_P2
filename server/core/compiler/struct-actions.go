package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {

	id := ctx.ID().GetText()

	newStruct := c.Visit(ctx.StructBody()).(*Object)

	newStruct.Type = id

	c.TAC.AddStruct(id, newStruct)

	return nil
}

func (c *Compiler) VisitStructBody(ctx *parser.StructBodyContext) interface{} {

	object := NewObject("struct", nil)

	for i, prop := range ctx.AllStructProperty() {
		varTypeDeclaration := prop.VariableDeclaration().(*parser.TypeDeclarationContext)
		id := varTypeDeclaration.ID().GetText()
		newValue := NewSimpleValue(i)
		newValue.SetData(IntTemporal, "")
		object.AddProp(id, newValue)
	}

	return object
}
