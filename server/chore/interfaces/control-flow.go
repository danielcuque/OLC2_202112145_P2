package interfaces

import (
	"OLC2/chore/parser"
	"fmt"
)

func (v *Visitor) VisitControlBreak(ctx *parser.ControlBreakContext) interface{} {

	peek := v.Stack.Peek()

	if !peek.Contains(BreakType) {
		v.NewError("No se puede usar break en este contexto", ctx.GetStart())
		return nil
	}

	peek.Trigger = BreakType

	// Throw panic to stop execution and return to the top of the stack
	panic(peek)
}

func (v *Visitor) VisitControlContinue(ctx *parser.ControlContinueContext) interface{} {
	peek := v.Stack.Peek()

	if !peek.Contains(ContinueType) {
		v.NewError("No se puede usar continue en este contexto", ctx.GetStart())
		return nil
	}

	peek.Trigger = ContinueType

	// Throw panic to stop execution and return to the top of the stack
	panic(peek)
}

func (v *Visitor) VisitControlReturn(ctx *parser.ControlReturnContext) interface{} {
	fmt.Println("VisitControlReturn")
	return nil
}
