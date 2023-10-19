package static

import (
	"OLC2/core/compiler"
	"OLC2/core/parser"
)

func (c *StaticVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	return c.Visit(ctx.Block())
}

func (c *StaticVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	for i := 0; ctx.Statement(i) != nil; i++ {
		c.Visit(ctx.Statement(i))
	}

	return nil
}

func (c *StaticVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	// Just visit variable declarations, possible scopes such as if, for, while, etc

	if ctx.VectorDeclaration() != nil {
		return c.Visit(ctx.VectorDeclaration())
	}

	if ctx.MatrixDeclaration() != nil {
		return c.Visit(ctx.MatrixDeclaration())
	}

	if ctx.VariableDeclaration() != nil {
		c.Visit(ctx.VariableDeclaration())
	}

	if ctx.IfStatement() != nil {
		c.Visit(ctx.IfStatement())
	}

	if ctx.ForStatement() != nil {
		c.Visit(ctx.ForStatement())
	}

	if ctx.WhileStatement() != nil {
		c.Visit(ctx.WhileStatement())
	}

	if ctx.GuardStatement() != nil {
		c.Visit(ctx.GuardStatement())
	}

	if ctx.SwitchStatement() != nil {
		c.Visit(ctx.SwitchStatement())
	}

	// if ctx.FunctionDeclarationStatement() != nil {
	// 	c.Visit(ctx.FunctionDeclarationStatement())
	// }

	return nil
}

func (c *StaticVisitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	c.SetEnv(compiler.WhileEnv, ctx.Block().(*parser.BlockContext))
	return nil
}

func (c *StaticVisitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	c.Env.PushEnv(compiler.ForEnv)
	c.Visit(ctx.Block())
	c.NewValue(ctx.ID().GetText())
	c.Env.PopEnv()
	return nil
}

func (c *StaticVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {

	for _, ifStmt := range ctx.AllIfTail() {
		ifStatement := ifStmt.(*parser.IfTailContext)
		c.SetEnv(compiler.IfEnv, ifStatement.Block().(*parser.BlockContext))
	}

	if ctx.ElseStatement() != nil {
		c.SetEnv(compiler.ElseEnv, ctx.ElseStatement().(*parser.ElseStatementContext).Block().(*parser.BlockContext))
	}

	return nil
}

func (c *StaticVisitor) VisitGuardStatement(ctx *parser.GuardStatementContext) interface{} {
	c.SetEnv(compiler.GuardEnv, ctx.Block().(*parser.BlockContext))
	return nil
}

func (c *StaticVisitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {

	for _, switchCase := range ctx.AllSwitchCase() {
		c.SetEnv(compiler.SwitchEnv, switchCase.(*parser.SwitchCaseContext).Block().(*parser.BlockContext))
	}

	if ctx.SwitchDefault() != nil {
		c.SetEnv(compiler.SwitchEnv, ctx.SwitchDefault().(*parser.SwitchDefaultContext).Block().(*parser.BlockContext))
	}

	return nil
}

// func (c *StaticVisitor) VisitFunctionDeclarationStatement(ctx *parser.FunctionDeclarationStatementContext) interface{} {

// 	id := ctx.ID().GetText()

// 	params := ctx.FunctionParameters().GetText()

// 	fmt.Println("FUNCTION", id, params)

// 	return nil
// }
