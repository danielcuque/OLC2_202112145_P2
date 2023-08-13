package interfaces

import (
	"OLC2/chore/parser"
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
			NewNilValue(nil),
			[]StackItemType{BreakType, ContinueType},
		))

	// Now get the value to iterate
	argIterator, ok := v.Visit(ctx.Expr()).(IValue)

	if !ok {
		v.NewError("Error al obtener el valor de la expresión", ctx.GetStart())
		return nil
	}

	// Get data type of the iterator
	iteratorType := argIterator.GetType()

	// Check if is string, array or number
	var valuesToIterate []IValue

	switch iteratorType {
	case IntType:
		valuesToIterate = argIterator.GetValue().([]IValue)
	case StringType:
		valuesToIterate = make([]IValue, len(argIterator.GetValue().(string)))
		for i, char := range argIterator.GetValue().(string) {
			valuesToIterate[i] = NewCharValue(char)
		}
	default:
		v.NewError(fmt.Sprintf("No se puede iterar un objeto de tipo %s", iteratorType), ctx.GetStart())
		return nil
	}

	NewForVariable(v, id, valuesToIterate[0], iteratorType, ctx)

	// Execute the for loop
	v.ExecuteFor(id, valuesToIterate, ctx)

	// for _, value := range valuesToIterate {
	// 	// Assign the value to the variable
	// 	v.Scope.SetVariable(id, value)

	// 	// Visit the for loop
	// 	v.Visit(ctx.Block())
	// }

	// Pop the scope
	v.Scope.PopScope()

	return nil
}

func NewForVariable(v *Visitor, id string, value IValue, valueType string, ctx *parser.ForStatementContext) {
	line, column, scope := GetVariableAttr(v, ctx.GetStart())
	newVariable := NewVariable(id, true, value, valueType, line, column, scope)
	v.Scope.AddVariable(id, newVariable)
}

func (v *Visitor) ExecuteFor(id string, valuesToIterate []IValue, ctx *parser.ForStatementContext) {
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
