package static

import (
	"OLC2/core/parser"
)

func (c *StaticVisitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	c.NewValue(id)
	return nil

}

func (c *StaticVisitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	c.NewValue(id)
	return nil

}

func (c *StaticVisitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	id := ctx.ID().GetText()
	c.NewValue(id)
	return nil
}
