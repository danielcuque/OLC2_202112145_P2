package interfaces

import (
	"OLC2/chore/parser"
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
	peek := v.Stack.Peek()

	if !peek.Contains(ReturnType) {
		v.NewError("No se puede usar return en este contexto", ctx.GetStart())
		return nil
	}

	peek.Trigger = ReturnType

	// Throw panic to stop execution and return to the top of the stack
	panic(peek)
}
