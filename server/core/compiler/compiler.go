package compiler

import (
	"OLC2/core/parser"
	"fmt"
)

type Compiler struct {
	parser.BaseSwiftVisitor
	Env          *EnvTree
	TAC          TAC
	HeapPointer  Heap
	StackPointer Stack
	StaticVars   int
}

func NewCompiler(staticVars int) *Compiler {
	newCompiler := &Compiler{
		TAC:        *NewTAC(),
		StaticVars: staticVars,
		Env:        NewEnvTree(),
	}

	newCompiler.GenerateVariables()
	return newCompiler
}

func (c *Compiler) GetMain() string {
	var code string
	code += "int main(){\n"
	code += c.Env.GetMain()
	code += "}\n"
	return code
}

func (c *Compiler) String() string {
	code := c.GetHeader()
	code += c.TAC.GetTemporalsHeader()
	code += c.TAC.GetProcudres()
	code += fmt.Sprintf("\nint main(){\n%s\nreturn 0;\n}", c.TAC.String())
	return code
}
