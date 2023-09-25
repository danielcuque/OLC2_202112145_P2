package c3d

import (
	E "OLC2/core/error"
	"OLC2/core/parser"
)

type Compiler struct {
	parser.BaseSwiftVisitor
	Errors []*E.VisitorError
	Env    *EnvTree
}

func NewCompiler() *Compiler {
	return &Compiler{}
}
