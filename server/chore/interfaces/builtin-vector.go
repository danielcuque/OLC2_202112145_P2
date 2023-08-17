package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

func Append(v *Visitor, ctx *parser.FunctionCallContext) interface{} {
	fmt.Println("Append")
	return V.NewNilValue(nil)
}
