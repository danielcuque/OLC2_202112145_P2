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

	// Visit a parse tree produced by SwiftParser#SimpleIfStatement.
	VisitSimpleIfStatement(ctx *SimpleIfStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#IfElseStatement.
	VisitIfElseStatement(ctx *IfElseStatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#IfElseIfStatement.
	VisitIfElseIfStatement(ctx *IfElseIfStatementContext) interface{}

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
}
