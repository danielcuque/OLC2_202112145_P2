package compiler

import (
	"OLC2/core/parser"

	"github.com/antlr4-go/antlr/v4"
)

type StaticVisitor struct {
	parser.BaseSwiftVisitor
	Env        *EnvTree
	Counter    int
	IsRelative bool
	Offset     int
	Pointer    *Temporal
}

func NewStaticVisitor(IsRelative bool, Offset int, Pointer *Temporal) *StaticVisitor {
	return &StaticVisitor{
		Env:        NewEnvTree(),
		IsRelative: IsRelative,
		Offset:     Offset,
	}
}

func (c *StaticVisitor) SetEnv(envType string, ctx *parser.BlockContext) {
	c.Env.PushEnv(envType)
	c.Visit(ctx)
	c.Env.PopEnv()
}

func (c *StaticVisitor) NewValue(name string) {
	value := c.Env.AddValue(name, NewSimpleValue(c.Counter))
	value.IsRelative = c.IsRelative
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
