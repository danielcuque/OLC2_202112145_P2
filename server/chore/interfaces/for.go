package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	// Get the for loop variables
	id := ctx.ID().GetText()

	// Now get the value to iterate
	argIterator, ok := v.Visit(ctx.Expr()).(V.IValue)

	if !ok {
		v.NewError(InvalidExpression, ctx.GetStart())
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
	case V.VectorType:
		valuesToIterate = argIterator.GetValue().(*VectorV).Body
	default:
		v.NewError(fmt.Sprintf("No se puede iterar un objeto de tipo %s", iteratorType), ctx.GetStart())
		return nil
	}

	// Create a new scope
	v.Env.PushEnv(ForEnv)

	v.Stack.Push(
		NewStackItem(
			"For",
			V.NewNilValue(nil),
			[]StackItemType{BreakType, ContinueType},
		))

	if len(valuesToIterate) != 0 {
		v.Env.AddVariable(id, NewVariable(v, id, true, valuesToIterate[0], iteratorType, ctx.GetStart()))

		v.ExecuteFor(id, valuesToIterate, ctx)
	}

	// Pop the scope
	v.Env.PopEnv()

	return nil
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
		v.Env.SetVariable(id, value)

		// Visit the for loop
		v.Visit(ctx.Block())
	}
}
