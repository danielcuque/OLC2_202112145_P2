package compiler

import (
	E "OLC2/core/error"
	"OLC2/core/parser"
	"fmt"
)

type Compiler struct {
	parser.BaseSwiftVisitor
	Errors []*E.VisitorError
	Env    *EnvTree
	TAC    TAC
}

func NewCompiler() *Compiler {
	return &Compiler{
		Errors: make([]*E.VisitorError, 0),
		Env:    NewEnvTree(),
	}
}

func (c *Compiler) GetMain() string {
	var code string
	code += "int main(){\n"
	code += c.Env.GetMain()
	code += "}\n"
	return code
}

func (c *Compiler) GetTAC() string {
	code := c.GetHeader()
	code += c.TAC.GetTemporalsHeader()
	code += fmt.Sprintf("void main(){\n%s\nreturn 0;\n}", c.TAC.GetCode())
	// code += c.TAC.GetCode()
	return code
}
