package compiler

import (
	"OLC2/core/parser"
)

func (c *Compiler) VisitMatrixDeclaration(ctx *parser.MatrixDeclarationContext) interface{} {
	// Create a matrix using a vector
	return nil
}

func (c *Compiler) VisitMatrixTypeNested(ctx *parser.MatrixTypeNestedContext) interface{} {
	return c.Visit(ctx.MatrixType())
}

func (c *Compiler) VisitMatrixTypeSingle(ctx *parser.MatrixTypeSingleContext) interface{} {
	return c.Visit(ctx.VariableType())
}

func (c *Compiler) VisitMatrixDefinition(ctx *parser.MatrixDefinitionContext) interface{} {
	return nil
}

func (v *Compiler) VisitMatrixValues(ctx *parser.MatrixValuesContext) interface{} {
	return nil
}

func (v *Compiler) VisitMatrixRepeatingDefinitionNested(ctx *parser.MatrixRepeatingDefinitionNestedContext) interface{} {
	return nil
}

func (v *Compiler) VisitMatrixRepeatingDefinitionSingle(ctx *parser.MatrixRepeatingDefinitionSingleContext) interface{} {
	return nil
}

func (v *Compiler) GetMatrixBody(ctx *parser.MatrixDeclarationContext) interface{} {
	return nil
}

func (v *Compiler) VisitMatrixAccess(ctx *parser.MatrixAccessContext) interface{} {
	return nil
}

func (v *Compiler) CheckMatrixIndexes() interface{} {
	return nil
}

func (v *Compiler) VisitMatrixAssignment(ctx *parser.MatrixAssignmentContext) interface{} {

	return nil
}

func (v *Compiler) ReplaceMatrixValue() {
}
