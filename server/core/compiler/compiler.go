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
}

func NewCompiler() *Compiler {
	return &Compiler{
		Env: NewEnvTree(),
		TAC: *NewTAC(),
	}
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
	code += fmt.Sprintf("\nint main(){\n%s\nreturn 0;\n}", c.TAC.GetCode())
	return code
}
