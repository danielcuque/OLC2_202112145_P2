package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitCallMethods(ctx *parser.CallMethodsContext) interface{} {
	fmt.Println("CallMethods")
	return nil
}
