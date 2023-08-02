package interfaces

import "OLC2/chore/parser"

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) Value {
	// Get the for loop variables
	return Value{ParseValue: false}
}
