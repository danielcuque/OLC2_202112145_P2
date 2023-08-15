package utils

import (
	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

// Create a new type of visitor to traverse antlr parse tree and return a graphviz dot string

type GraphvizVisitor struct {
	*parser.BaseSwiftVisitor
	output string
}

func NewGraphvizVisitor() *GraphvizVisitor {
	return &GraphvizVisitor{
		output: "digraph G {\n",
	}
}

func (g *GraphvizVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch tree.(type) {
	case *antlr.ErrorNodeImpl:
		return nil
	default:
		node := tree.Accept(g)
		return node
	}
}

func (g *GraphvizVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	g.Visit(ctx.Block())
	return nil
}

func (g *GraphvizVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	for i := 0; ctx.Statement(i) != nil; i++ {
		g.Visit(ctx.Statement(i))
	}
	return nil
}

func (g *GraphvizVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.VariableAssignment() != nil {
		return g.Visit(ctx.VariableAssignment())
	}
	if ctx.VariableDeclaration() != nil {
		return g.Visit(ctx.VariableDeclaration())
	}

	if ctx.IfStatement() != nil {
		return g.Visit(ctx.IfStatement())
	}

	return nil
}
