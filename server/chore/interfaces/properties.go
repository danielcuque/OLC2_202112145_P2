package interfaces

// func (v *Visitor) VisitCallProperties(ctx *parser.CallPropertiesContext) interface{} {
// 	id := ctx.ID(0).GetText()

// 	variable, ok := v.Scope.GetVariable(id).(*Variable)

// 	if !ok {
// 		v.NewError(fmt.Sprint("La variable ", id, " no existe"), ctx.GetStart())
// 	}

// 	var params []string

// 	/*
// 		Create a recursive auxiliar function to traverse all the properties.
// 		For example, if we have a call like this:
// 		foo.bar.baz, tiene que retornar el valor de baz, pero para eso, primero
// 		tiene que retornar el valor de bar, y para eso, primero tiene que retornar
// 		el valor de foo.
// 	*/

// 	// Get all ids except the first, because the first will be the object, and we need the properties

// 	for _, id := range ctx.AllID()[1:] {
// 		params = append(params, id.GetText())
// 	}

// 	prop, okP := GetPropValue(variable, params).(*Variable)

// 	if !okP {
// 		v.NewError(fmt.Sprint("La propiedad ", params[len(params)-1], " no existe"), ctx.GetStart())
// 		return V.NewNilValue(nil)
// 	}

// 	return prop.Value

// }
