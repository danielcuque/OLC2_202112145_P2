// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SwiftParser.
type SwiftVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SwiftParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SwiftParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#variableType.
	VisitVariableType(ctx *VariableTypeContext) interface{}

	// Visit a parse tree produced by SwiftParser#TypeValueDeclaration.
	VisitTypeValueDeclaration(ctx *TypeValueDeclarationContext) interface{}

	// Visit a parse tree produced by SwiftParser#ValueDeclaration.
	VisitValueDeclaration(ctx *ValueDeclarationContext) interface{}

	// Visit a parse tree produced by SwiftParser#TypeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by SwiftParser#variableAssignment.
	VisitVariableAssignment(ctx *VariableAssignmentContext) interface{}

	// Visit a parse tree produced by SwiftParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#ifTail.
	VisitIfTail(ctx *IfTailContext) interface{}

	// Visit a parse tree produced by SwiftParser#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by SwiftParser#switchDefault.
	VisitSwitchDefault(ctx *SwitchDefaultContext) interface{}

	// Visit a parse tree produced by SwiftParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#guardStatement.
	VisitGuardStatement(ctx *GuardStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#RangeExpr.
	VisitRangeExpr(ctx *RangeExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#CharExpr.
	VisitCharExpr(ctx *CharExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#ArithmeticExpr.
	VisitArithmeticExpr(ctx *ArithmeticExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#ComparasionExpr.
	VisitComparasionExpr(ctx *ComparasionExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#LogicalExpr.
	VisitLogicalExpr(ctx *LogicalExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#TernaryExpr.
	VisitTernaryExpr(ctx *TernaryExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#ControlBreak.
	VisitControlBreak(ctx *ControlBreakContext) interface{}

	// Visit a parse tree produced by SwiftParser#ControlContinue.
	VisitControlContinue(ctx *ControlContinueContext) interface{}

	// Visit a parse tree produced by SwiftParser#ControlReturn.
	VisitControlReturn(ctx *ControlReturnContext) interface{}
}
