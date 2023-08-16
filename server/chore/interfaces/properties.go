package interfaces

import (
	"OLC2/chore/parser"
	V "OLC2/chore/values"
	"fmt"
)

func (v *Visitor) VisitCallProperties(ctx *parser.CallPropertiesContext) interface{} {
	id := ctx.ID(0).GetText()

	variable, ok := v.Scope.GetVariable(id).(*Variable)

	if !ok {
		v.NewError(fmt.Sprint("La variable ", id, " no existe"), ctx.GetStart())
	}

	var params []string

	/*
		Create a recursive auxiliar function to traverse all the properties.
		For example, if we have a call like this:
		foo.bar.baz, tiene que retornar el valor de baz, pero para eso, primero
		tiene que retornar el valor de bar, y para eso, primero tiene que retornar
		el valor de foo.
	*/

	// Get all ids except the first, because the first will be the object, and we need the properties

	for _, id := range ctx.AllID()[1:] {
		params = append(params, id.GetText())
	}

	fmt.Println("params: ", params)
	fmt.Println("variable: ", variable.Name)

	return V.NewNilValue(nil)
}

// func getProperty(obj interface{}, properties []string) (interface{}, error) {
// 	if len(properties) == 0 {
// 		return obj, nil
// 	}

// 	propertyName := properties[0]

// 	// Aquí debes implementar la lógica para obtener la propiedad "propertyName" del objeto "obj"

// 	// Luego, recursivamente llamamos a getProperty con el objeto de la propiedad y las propiedades restantes
// 	return getProperty(newObj, properties[1:])
// }
