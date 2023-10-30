package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {

	id := ctx.ID().GetText()

	fmt.Println("StructDeclaration: ", id)

	return nil
}
