// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import "github.com/antlr4-go/antlr/v4"

type BaseSwiftVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSwiftVisitor) VisitIDChain(ctx *IDChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVariableType(ctx *VariableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitTypeValueDeclaration(ctx *TypeValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitValueDeclaration(ctx *ValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionDeclarationStatement(ctx *FunctionDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionParameters(ctx *FunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionParameterCompoundNested(ctx *FunctionParameterCompoundNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionParameterCompoundSingle(ctx *FunctionParameterCompoundSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionReturnType(ctx *FunctionReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitNamedArguments(ctx *NamedArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVariableAssignment(ctx *VariableAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVariableAssignmentObject(ctx *VariableAssignmentObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitIfTail(ctx *IfTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitElseStatement(ctx *ElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitSwitchDefault(ctx *SwitchDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitGuardStatement(ctx *GuardStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorTypeValue(ctx *VectorTypeValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorStructValue(ctx *VectorStructValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorListValue(ctx *VectorListValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorSingleValue(ctx *VectorSingleValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorValues(ctx *VectorValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitObjectChain(ctx *ObjectChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorAccess(ctx *VectorAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorAssignment(ctx *VectorAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixDeclaration(ctx *MatrixDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixTypeNested(ctx *MatrixTypeNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixTypeSingle(ctx *MatrixTypeSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixDefinition(ctx *MatrixDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixValues(ctx *MatrixValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixRepeatingDefinitionNested(ctx *MatrixRepeatingDefinitionNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixRepeatingDefinitionSingle(ctx *MatrixRepeatingDefinitionSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixAccess(ctx *MatrixAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixAssignment(ctx *MatrixAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitStructProperty(ctx *StructPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitComparisonExpr(ctx *ComparisonExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitRangeExpr(ctx *RangeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitCharExpr(ctx *CharExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVectorAccessExpr(ctx *VectorAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitArithmeticExpr(ctx *ArithmeticExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitStrExpr(ctx *StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitLogicalExpr(ctx *LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitObjectChainExpr(ctx *ObjectChainExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitMatrixAccessExpr(ctx *MatrixAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitControlBreak(ctx *ControlBreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitControlContinue(ctx *ControlContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitControlReturn(ctx *ControlReturnContext) interface{} {
	return v.VisitChildren(ctx)
}
