package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.VariableAssignment() != nil {
		return v.Visit(ctx.VariableAssignment())
	}
	if ctx.VariableDeclaration() != nil {
		return v.Visit(ctx.VariableDeclaration())
	}

	if ctx.IfStatement() != nil {
		return v.Visit(ctx.IfStatement())
	}

	// if ctx.WhileStatement() != nil {
	// 	return v.Visit(ctx.WhileStatement())
	// }

	// if ctx.ForStatement() != nil {
	// 	return v.Visit(ctx.ForStatement())
	// }
	return nil
}
