package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

func Count(v *Visitor, ctx *parser.CallPropertiesContext) interface{} {
	fmt.Println("Count")
	return V.NewNilValue(nil)
}
