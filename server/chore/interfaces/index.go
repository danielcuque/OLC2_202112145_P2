package interfaces

import (
	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		v.NewError(val.GetText(), val.GetSymbol())
		return nil
	default:
		return tree.Accept(v)
	}
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {

	// First we visit structs declarations
	for i := 0; ctx.Statement(i) != nil; i++ {
		if ctx.Statement(i).(*parser.StatementContext).StructDeclaration() != nil {
			v.Visit(ctx.Statement(i))
		}
	}

	// Second we visit the function declarations

	for i := 0; ctx.Statement(i) != nil; i++ {
		if ctx.Statement(i).(*parser.StatementContext).FunctionDeclarationStatement() != nil {
			v.Visit(ctx.Statement(i))
		}
	}

	// Then we visit the statements
	for i := 0; ctx.Statement(i) != nil; i++ {
		if ctx.Statement(i).(*parser.StatementContext).FunctionDeclarationStatement() == nil && ctx.Statement(i).(*parser.StatementContext).StructDeclaration() == nil {
			v.Visit(ctx.Statement(i))
		}
	}

	return nil
}
