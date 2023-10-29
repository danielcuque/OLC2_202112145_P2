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

	c.GetMatrixBody(ctx)

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
	if ctx.MatrixValues() != nil {
		return c.Visit(ctx.MatrixValues())
	}

	return c.Visit(ctx.Expr())
}

func (v *Compiler) VisitMatrixValues(ctx *parser.MatrixValuesContext) interface{} {
	expr, ok := v.Visit(ctx.MatrixDefinition(0)).(*ValueResponse)

	if !ok {
		fmt.Println("Error al obtener el valor de la matriz")
		return expr
	}

	initMatrix := v.InitNewMatrix()

	if expr.GetType() == MatrixTemporal {
		value := expr.GetValue().(map[string]interface{})
		v.AppendVectorValue(value["temporals"].([]*Temporal)[0])
	} else {
		v.AppendVectorValue(expr.GetValue())
	}

	// 0 = matrix pointer
	// 1 = matrix size
	// 2 = isEmpty

	baseNode := NewMatrix(expr.GetType())

	for _, matrixDef := range ctx.AllMatrixDefinition()[1:] {
		matrixDef := v.Visit(matrixDef).(*ValueResponse)

		if matrixDef.GetType() == MatrixTemporal {
			value := expr.GetValue().(map[string]interface{})
			v.AppendVectorValue(value["temporals"].([]*Temporal)[0])
			baseNode.AddValue(value["matrix"])
		} else {
			v.AppendVectorValue(matrixDef.GetValue())
			baseNode.AddValue(matrixDef.GetValue())
		}

	}

	v.TAC.AppendInstruction(
		fmt.Sprintf("%v = %v;", initMatrix[1], len(ctx.AllMatrixDefinition())),
		"Contador de vector",
	)

	v.DeclareMatrixProps(initMatrix[0], initMatrix[1], initMatrix[2])

	return &ValueResponse{
		Value:       map[string]interface{}{"matrix": baseNode, "temporals": initMatrix},
		Type:        MatrixTemporal,
		ContextType: TemporalType,
	}
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
		body = v.Visit(ctx.MatrixDefinition())
	} else {
		body = v.Visit(ctx.MatrixRepeatingDefinition())
	}

	return body
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
