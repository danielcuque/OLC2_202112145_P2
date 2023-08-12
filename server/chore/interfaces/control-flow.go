package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitControlBreak(ctx *parser.ControlBreakContext) interface{} {

	return nil
}

func (v *Visitor) VisitControlContinue(ctx *parser.ControlContinueContext) interface{} {
	fmt.Println("VisitControlContinue")
	return nil
}

func (v *Visitor) VisitControlReturn(ctx *parser.ControlReturnContext) interface{} {
	fmt.Println("VisitControlReturn")
	return nil
}
