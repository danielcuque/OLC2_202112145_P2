// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import "github.com/antlr4-go/antlr/v4"

type BaseSwiftVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSwiftVisitor) VisitData_type(ctx *Data_typeContext) interface{} {
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

func (v *BaseSwiftVisitor) VisitVariable_assignment(ctx *Variable_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitVariable_declaration(ctx *Variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitIfstmt(ctx *IfstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitWhilestmt(ctx *WhilestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitDoubleExpr(ctx *DoubleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
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

func (v *BaseSwiftVisitor) VisitComparasionExpr(ctx *ComparasionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitLogicalExpr(ctx *LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSwiftVisitor) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}
