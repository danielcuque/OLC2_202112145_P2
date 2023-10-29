package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitMatrixDeclaration(ctx *parser.MatrixDeclarationContext) interface{} {
	id, _ := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	matrixType := c.Visit(ctx.MatrixType()).(TemporalCast)

	initVector := c.GetMatrixBody(ctx).(*ValueResponse)

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[(int)%v] = %v;", value.GetAddress(), initVector.GetValue()), // Inicio del vector
		"Direccion de la matriz",
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

	if ctx.MatrixValues() != nil {
		newMatrix := c.NewInitMatrix()

		responses := make([]*ValueResponse, 0)

		allMatrixDefinition := ctx.MatrixValues().AllMatrixDefinition()

		for _, matrixDef := range allMatrixDefinition {
			response := c.GetMatrixValues(matrixDef.(*parser.MatrixDefinitionContext))

			responses = append(responses, response)
		}

		c.DeclareNewMatrix(newMatrix)

		for _, response := range responses {
			c.AppendVectorValue(response.GetValue())
		}

		c.TAC.AppendInstruction(
			fmt.Sprintf("%v = %v;", newMatrix.Counter, len(allMatrixDefinition)),
			"Contador de vector",
		)

		c.DeclareMatrixProps(newMatrix.InitVector, newMatrix.Counter, newMatrix.IsEmptyAddress)

		return &ValueResponse{
			Value:       newMatrix.InitVector,
			ContextType: TemporalType,
			Type:        MatrixTemporal,
		}

	}

	return nil
}

// This function will be recursive
func (c *Compiler) GetMatrixValues(ctx *parser.MatrixDefinitionContext) *ValueResponse {
	if ctx.MatrixValues() != nil {
		newMatrix := c.NewInitMatrix()
		responses := make([]*ValueResponse, 0)

		allMatrixDefinition := ctx.MatrixValues().AllMatrixDefinition()

		for _, matrixDef := range allMatrixDefinition {

			response := c.GetMatrixValues(matrixDef.(*parser.MatrixDefinitionContext))
			responses = append(responses, response)
		}

		c.DeclareNewMatrix(newMatrix)

		for _, response := range responses {
			c.AppendVectorValue(response.GetValue())
		}

		c.TAC.AppendInstruction(
			fmt.Sprintf("%v = %v;", newMatrix.Counter, len(allMatrixDefinition)),
			"Contador de vector",
		)

		c.DeclareMatrixProps(newMatrix.InitVector, newMatrix.Counter, newMatrix.IsEmptyAddress)

		return &ValueResponse{
			Value: newMatrix.InitVector,
		}

	}

	if ctx.Expr() != nil {
		return c.Visit(ctx.Expr()).(*ValueResponse)
	}

	return nil
}

func (v *Compiler) VisitMatrixRepeatingDefinitionNested(ctx *parser.MatrixRepeatingDefinitionNestedContext) interface{} {
	return nil
}

func (v *Compiler) VisitMatrixRepeatingDefinitionSingle(ctx *parser.MatrixRepeatingDefinitionSingleContext) interface{} {
	return nil
}

func (v *Compiler) GetMatrixBody(ctx *parser.MatrixDeclarationContext) interface{} {

	// The body can be defined explicitly or implicitly

	var body interface{}

	if ctx.MatrixDefinition() != nil {
		// Convert node to array
		body = v.Visit(ctx.MatrixDefinition()).(*ValueResponse)
	} else {
		body = v.Visit(ctx.MatrixRepeatingDefinition()).(*ValueResponse)
	}

	return body
}

func (c *Compiler) VisitMatrixAccess(ctx *parser.MatrixAccessContext) interface{} {
	id, _ := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	firstIndex := c.Visit(ctx.Expr(0)).(*ValueResponse)

	matrixValueAddress := c.VectorAccess(
		c.TAC.GetValueAddress(value),
		firstIndex,
		true,
		value.GetType(),
	)

	matrixValue := c.GetVectorValue(matrixValueAddress)

	for _, expr := range ctx.AllExpr()[1:] {
		index := c.Visit(expr).(*ValueResponse)

		exprValueAddress := c.VectorAccess(
			fmt.Sprintf("%v", matrixValue.GetValue()),
			index,
			false,
			value.GetType(),
		)

		exprValue := c.GetVectorValue(exprValueAddress)
		matrixValue = exprValue
	}

	return matrixValue
}

func (v *Compiler) CheckMatrixIndexes() interface{} {
	return nil
}

func (v *Compiler) VisitMatrixAssignment(ctx *parser.MatrixAssignmentContext) interface{} {

	return nil
}

func (v *Compiler) ReplaceMatrixValue() {
}
