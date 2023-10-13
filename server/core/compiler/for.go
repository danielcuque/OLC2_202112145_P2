package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

func (c *Compiler) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	fmt.Println("VisitForStatement")
	return nil
}
