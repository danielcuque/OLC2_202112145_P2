package compiler

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

func (c *StaticVisitor) VisitVectorTypeValue(ctx *parser.VectorTypeValueContext) interface{} {
	id := ctx.ID().GetText()
	c.NewValue(id)
	return nil
}

func (c *StaticVisitor) VisitVectorStructValue(ctx *parser.VectorStructValueContext) interface{} {
	id := ctx.ID(0).GetText()
	c.NewValue(id)
	return nil
}

func (c *StaticVisitor) VisitMatrixDeclaration(ctx *parser.MatrixDeclarationContext) interface{} {
	id := ctx.IdChain().GetText()
	c.NewValue(id)
	return nil
}
