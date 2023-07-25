// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import "github.com/antlr4-go/antlr/v4"

type BaseSwiftVisitor struct {
	*antlr.BaseParseTreeVisitor
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

func (v *BaseSwiftVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitData_type(ctx *Data_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitDigit(ctx *DigitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitParen(ctx *ParenContext) interface{} {
	return v.VisitChildren(ctx)
}
