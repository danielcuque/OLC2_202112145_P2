package static

import (
	"OLC2/core/compiler"
	"OLC2/core/parser"

	"github.com/antlr4-go/antlr/v4"
)

type StaticVisitor struct {
	parser.BaseSwiftVisitor
	Env     *compiler.EnvTree
	Counter int
}

func NewStaticVisitor() *StaticVisitor {
	return &StaticVisitor{
		Env: compiler.NewEnvTree(),
	}
}

func (c *StaticVisitor) SetEnv(envType string, ctx *parser.BlockContext) {
	c.Env.PushEnv(envType)
	c.Visit(ctx)
	c.Env.PopEnv()
}

func (c *StaticVisitor) NewValue(name string) {
	c.Env.AddValue(name, compiler.NewSimpleValue(c.Counter))
	c.Counter++
}

func (c *StaticVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch tree.(type) {
	case *antlr.ErrorNodeImpl:
		return nil
	default:
		return tree.Accept(c)
	}
}
