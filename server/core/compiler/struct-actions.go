package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {

	id := ctx.ID().GetText()

	newStruct := c.Visit(ctx.StructBody()).(*Object)

	newStruct.Type = id

	newValue := NewSimpleValue(0)

	newValue.SetData(StructTemporal, newStruct)

	c.Env.AddValue(id, newValue)

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

func (c *Compiler) HandleStructInstance(ctx *parser.FunctionCallContext) interface{} {
	fmt.Println("Struct instance")
	return nil
}
