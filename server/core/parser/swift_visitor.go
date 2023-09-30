// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Swift.
type SwiftVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Swift#IDChain.
	VisitIDChain(ctx *IDChainContext) interface{}

	// Visit a parse tree produced by Swift#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by Swift#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by Swift#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Swift#variableType.
	VisitVariableType(ctx *VariableTypeContext) interface{}

	// Visit a parse tree produced by Swift#TypeValueDeclaration.
	VisitTypeValueDeclaration(ctx *TypeValueDeclarationContext) interface{}

	// Visit a parse tree produced by Swift#ValueDeclaration.
	VisitValueDeclaration(ctx *ValueDeclarationContext) interface{}

	// Visit a parse tree produced by Swift#TypeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by Swift#functionDeclarationStatement.
	VisitFunctionDeclarationStatement(ctx *FunctionDeclarationStatementContext) interface{}

	// Visit a parse tree produced by Swift#functionParameters.
	VisitFunctionParameters(ctx *FunctionParametersContext) interface{}

	// Visit a parse tree produced by Swift#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) interface{}

	// Visit a parse tree produced by Swift#FunctionParameterCompoundNested.
	VisitFunctionParameterCompoundNested(ctx *FunctionParameterCompoundNestedContext) interface{}

	// Visit a parse tree produced by Swift#FunctionParameterCompoundSingle.
	VisitFunctionParameterCompoundSingle(ctx *FunctionParameterCompoundSingleContext) interface{}

	// Visit a parse tree produced by Swift#functionReturnType.
	VisitFunctionReturnType(ctx *FunctionReturnTypeContext) interface{}

	// Visit a parse tree produced by Swift#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by Swift#Arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by Swift#NamedArguments.
	VisitNamedArguments(ctx *NamedArgumentsContext) interface{}

	// Visit a parse tree produced by Swift#variableAssignment.
	VisitVariableAssignment(ctx *VariableAssignmentContext) interface{}

	// Visit a parse tree produced by Swift#variableAssignmentObject.
	VisitVariableAssignmentObject(ctx *VariableAssignmentObjectContext) interface{}

	// Visit a parse tree produced by Swift#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by Swift#ifTail.
	VisitIfTail(ctx *IfTailContext) interface{}

	// Visit a parse tree produced by Swift#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by Swift#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by Swift#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by Swift#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by Swift#switchDefault.
	VisitSwitchDefault(ctx *SwitchDefaultContext) interface{}

	// Visit a parse tree produced by Swift#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by Swift#guardStatement.
	VisitGuardStatement(ctx *GuardStatementContext) interface{}

	// Visit a parse tree produced by Swift#VectorTypeValue.
	VisitVectorTypeValue(ctx *VectorTypeValueContext) interface{}

	// Visit a parse tree produced by Swift#VectorStructValue.
	VisitVectorStructValue(ctx *VectorStructValueContext) interface{}

	// Visit a parse tree produced by Swift#VectorListValue.
	VisitVectorListValue(ctx *VectorListValueContext) interface{}

	// Visit a parse tree produced by Swift#VectorSingleValue.
	VisitVectorSingleValue(ctx *VectorSingleValueContext) interface{}

	// Visit a parse tree produced by Swift#vectorValues.
	VisitVectorValues(ctx *VectorValuesContext) interface{}

	// Visit a parse tree produced by Swift#objectChain.
	VisitObjectChain(ctx *ObjectChainContext) interface{}

	// Visit a parse tree produced by Swift#vectorAccess.
	VisitVectorAccess(ctx *VectorAccessContext) interface{}

	// Visit a parse tree produced by Swift#vectorAssignment.
	VisitVectorAssignment(ctx *VectorAssignmentContext) interface{}

	// Visit a parse tree produced by Swift#matrixDeclaration.
	VisitMatrixDeclaration(ctx *MatrixDeclarationContext) interface{}

	// Visit a parse tree produced by Swift#MatrixTypeNested.
	VisitMatrixTypeNested(ctx *MatrixTypeNestedContext) interface{}

	// Visit a parse tree produced by Swift#MatrixTypeSingle.
	VisitMatrixTypeSingle(ctx *MatrixTypeSingleContext) interface{}

	// Visit a parse tree produced by Swift#matrixDefinition.
	VisitMatrixDefinition(ctx *MatrixDefinitionContext) interface{}

	// Visit a parse tree produced by Swift#matrixValues.
	VisitMatrixValues(ctx *MatrixValuesContext) interface{}

	// Visit a parse tree produced by Swift#MatrixRepeatingDefinitionNested.
	VisitMatrixRepeatingDefinitionNested(ctx *MatrixRepeatingDefinitionNestedContext) interface{}

	// Visit a parse tree produced by Swift#MatrixRepeatingDefinitionSingle.
	VisitMatrixRepeatingDefinitionSingle(ctx *MatrixRepeatingDefinitionSingleContext) interface{}

	// Visit a parse tree produced by Swift#matrixAccess.
	VisitMatrixAccess(ctx *MatrixAccessContext) interface{}

	// Visit a parse tree produced by Swift#matrixAssignment.
	VisitMatrixAssignment(ctx *MatrixAssignmentContext) interface{}

	// Visit a parse tree produced by Swift#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by Swift#structBody.
	VisitStructBody(ctx *StructBodyContext) interface{}

	// Visit a parse tree produced by Swift#structProperty.
	VisitStructProperty(ctx *StructPropertyContext) interface{}

	// Visit a parse tree produced by Swift#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by Swift#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by Swift#NilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by Swift#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by Swift#ComparisonExpr.
	VisitComparisonExpr(ctx *ComparisonExprContext) interface{}

	// Visit a parse tree produced by Swift#RangeExpr.
	VisitRangeExpr(ctx *RangeExprContext) interface{}

	// Visit a parse tree produced by Swift#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by Swift#CharExpr.
	VisitCharExpr(ctx *CharExprContext) interface{}

	// Visit a parse tree produced by Swift#VectorAccessExpr.
	VisitVectorAccessExpr(ctx *VectorAccessExprContext) interface{}

	// Visit a parse tree produced by Swift#FunctionCallExpr.
	VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{}

	// Visit a parse tree produced by Swift#ArithmeticExpr.
	VisitArithmeticExpr(ctx *ArithmeticExprContext) interface{}

	// Visit a parse tree produced by Swift#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by Swift#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by Swift#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by Swift#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by Swift#LogicalExpr.
	VisitLogicalExpr(ctx *LogicalExprContext) interface{}

	// Visit a parse tree produced by Swift#TernaryExpr.
	VisitTernaryExpr(ctx *TernaryExprContext) interface{}

	// Visit a parse tree produced by Swift#ObjectChainExpr.
	VisitObjectChainExpr(ctx *ObjectChainExprContext) interface{}

	// Visit a parse tree produced by Swift#MatrixAccessExpr.
	VisitMatrixAccessExpr(ctx *MatrixAccessExprContext) interface{}

	// Visit a parse tree produced by Swift#ControlBreak.
	VisitControlBreak(ctx *ControlBreakContext) interface{}

	// Visit a parse tree produced by Swift#ControlContinue.
	VisitControlContinue(ctx *ControlContinueContext) interface{}

	// Visit a parse tree produced by Swift#ControlReturn.
	VisitControlReturn(ctx *ControlReturnContext) interface{}
}
