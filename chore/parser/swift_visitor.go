// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SwiftParser.
type SwiftVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SwiftParser#data_type.
	VisitData_type(ctx *Data_typeContext) interface{}

	// Visit a parse tree produced by SwiftParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SwiftParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SwiftParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SwiftParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SwiftParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by SwiftParser#ifstmt.
	VisitIfstmt(ctx *IfstmtContext) interface{}

	// Visit a parse tree produced by SwiftParser#whilestmt.
	VisitWhilestmt(ctx *WhilestmtContext) interface{}

	// Visit a parse tree produced by SwiftParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by SwiftParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}
}
