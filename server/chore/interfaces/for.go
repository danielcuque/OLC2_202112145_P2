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

	// Now get the value to iterate
	argIterator, ok := v.Visit(ctx.Expr()).(*RangeV)

	if !ok {
		v.NewError("Error al obtener el valor de la expresión", ctx.GetStart())
		return nil
	}

	// Get data type of the iterator
	iteratorType := argIterator.GetType()

	// Check if is string, array or number
	line, column, scope := GetVariableAttr(v, ctx.GetStart())

	if iteratorType == IntType {
		newVariable := NewVariable(
			id,
			false,
			NewIntValue(0),
			iteratorType,
			line,
			column,
			scope)
		v.Scope.AddVariable(id, newVariable)

		// The array range store pointers of int values
		for _, value := range argIterator.GetValue().([]IValue) {
			// Assign the value to the variable
			v.Scope.SetVariable(id, NewIntValue(value.GetValue().(int)))

			// Visit the for loop
			v.Visit(ctx.Block())
		}

	} else {
		v.NewError(fmt.Sprintf("No se puede iterar un objeto de tipo %s", iteratorType), ctx.GetStart())
		return nil
	}

	// Pop the scope
	v.Scope.PopScope()

	return nil
}
