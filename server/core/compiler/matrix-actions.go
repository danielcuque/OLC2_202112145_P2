package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitMatrixDeclaration(ctx *parser.MatrixDeclarationContext) interface{} {
	// Create a vector of vectors to handle the matrix

	id, _ := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	matrixType := c.Visit(ctx.MatrixType()).(TemporalCast)

	// matrixValue := c.GetMatrixBody(ctx)

	newMatrix := c.InitNewMatrix()

	c.DeclareMatrixProps(newMatrix[0], newMatrix[1], newMatrix[2])

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[(int)%v] = %v;", value.GetAddress(), newMatrix[0]), // Inicio del vector
		"Direccion de vector",
	)

	newMatrixObject := NewVector(matrixType)

	value.SetData(MatrixTemporal, newMatrixObject)

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
