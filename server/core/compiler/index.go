package compiler

import (
	"OLC2/core/parser"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) Visit(tree antlr.ParseTree) interface{} {
	switch tree.(type) {
	case *antlr.ErrorNodeImpl:
		return nil
	default:
		return tree.Accept(c)
	}
}

func (c *Compiler) VisitProgram(ctx *parser.ProgramContext) interface{} {
	return c.Visit(ctx.Block())
}

func (c *Compiler) VisitBlock(ctx *parser.BlockContext) interface{} {

	// First we visit structs declarations
	for i := 0; ctx.Statement(i) != nil; i++ {
		if ctx.Statement(i).(*parser.StatementContext).StructDeclaration() != nil {
			c.Visit(ctx.Statement(i))
		}
	}

	// Second we visit the function declarations

	for i := 0; ctx.Statement(i) != nil; i++ {
		if ctx.Statement(i).(*parser.StatementContext).FunctionDeclarationStatement() != nil {
			c.Visit(ctx.Statement(i))
		}
	}

	// Then we visit the statements
	for i := 0; ctx.Statement(i) != nil; i++ {
		if ctx.Statement(i).(*parser.StatementContext).FunctionDeclarationStatement() == nil && ctx.Statement(i).(*parser.StatementContext).StructDeclaration() == nil {
			c.Visit(ctx.Statement(i))
		}
	}

	return nil
}
