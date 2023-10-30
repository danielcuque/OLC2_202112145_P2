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
	args := c.GetArgs(ctx)

	baseTemporal := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = H;", baseTemporal),
			fmt.Sprintf("heap[(int)H] = %v", len(args)),
			c.HeapPointer.IncreasePointer(),
		},
		"Instancia de struct",
	)

	for _, arg := range args {
		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("heap[(int)H] = %v", arg.Value.GetValue()),
				c.HeapPointer.IncreasePointer(),
			},
			"",
		)
	}

	return &ValueResponse{
		Value:       baseTemporal,
		Type:        StructTemporal,
		ContextType: PointerType,
	}
}
