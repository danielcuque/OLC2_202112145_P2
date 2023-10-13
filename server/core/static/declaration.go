package static

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *StaticVisitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	// c.Env.AddValue(ctx.ID().GetText())
	return nil

}

func (c *StaticVisitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	fmt.Println(id)
	return nil

}

func (c *StaticVisitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	fmt.Println(id)
	return nil
}
