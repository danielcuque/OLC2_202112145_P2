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

	// Visit a parse tree produced by SwiftParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by SwiftParser#data_type.
	VisitData_type(ctx *Data_typeContext) interface{}

	// Visit a parse tree produced by SwiftParser#Op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by SwiftParser#Digit.
	VisitDigit(ctx *DigitContext) interface{}

	// Visit a parse tree produced by SwiftParser#Id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by SwiftParser#Paren.
	VisitParen(ctx *ParenContext) interface{}
}
