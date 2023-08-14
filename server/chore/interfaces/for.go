package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	// Get the for loop variables
	id := ctx.ID().GetText()

	// Create a new scope
	v.Scope.PushScope(IfScope)

	v.Stack.Push(
		NewStackItem(
			"For",
			V.NewNilValue(nil),
			[]StackItemType{BreakType, ContinueType},
		))

	// Now get the value to iterate
	argIterator, ok := v.Visit(ctx.Expr()).(V.IValue)

	if !ok {
		v.NewError(InvalidExpressionError, ctx.GetStart())
		return nil
	}

	// Get data type of the iterator
	iteratorType := argIterator.GetType()

	// Check if is string, array or number
	var valuesToIterate []V.IValue

	switch iteratorType {
	case V.IntType:
		valuesToIterate = argIterator.GetValue().([]V.IValue)
	case V.StringType:
		valuesToIterate = make([]V.IValue, len(argIterator.GetValue().(string)))
		for i, char := range argIterator.GetValue().(string) {
			valuesToIterate[i] = V.NewCharValue(char)
		}
	default:
		v.NewError(fmt.Sprintf("No se puede iterar un objeto de tipo %s", iteratorType), ctx.GetStart())
		return nil
	}

	NewForVariable(v, id, valuesToIterate[0], iteratorType, ctx)

	// Execute the for loop
	v.ExecuteFor(id, valuesToIterate, ctx)

	// Pop the scope
	v.Scope.PopScope()

	return nil
}

func NewForVariable(v *Visitor, id string, value V.IValue, valueType string, ctx *parser.ForStatementContext) {
	line, column, scope := GetVariableAttr(v, ctx.GetStart())
	newVariable := NewVariable(id, true, value, valueType, line, column, scope)
	v.Scope.AddVariable(id, newVariable)
}

func (v *Visitor) ExecuteFor(id string, valuesToIterate []V.IValue, ctx *parser.ForStatementContext) {
	defer func() {
		peek, ok := recover().(*StackItem)

		if !ok {
			return
		}

		if peek.Trigger == ReturnType {
			panic(peek)
		}

		if peek.Trigger == BreakType {
			return
		}

		if peek.Trigger == ContinueType {
			v.ExecuteFor(id, valuesToIterate, ctx)
		}
	}()

	for _, value := range valuesToIterate {
		// Assign the value to the variable
		v.Scope.SetVariable(id, value)

		// Visit the for loop
		v.Visit(ctx.Block())
	}
}
