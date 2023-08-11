package main

import (
	"OLC2/api"
)

func main() {

	// content := U.ReadFile("./examples/test.swift")

	// result := I.NewEvaluator(content)

	// for _, err := range result.Errors {
	// 	fmt.Println(err.Error())
	// }

	// variables := result.Scope.GetSymbolTable()

	// fmt.Println()
	// for _, v := range variables {
	// 	fmt.Println("Name:", v.GetName(), "Value:", v.GetValue(), "Is constant:", v.IsConstant(), "Value Type: ", v.GetType(), "Scope: ", v.GetScopeName(), "Line: ", v.GetLine(), "Column: ", v.GetColumn())
	// }

	api.Init()
}
