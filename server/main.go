package main

import (
	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"

	"fmt"
)

func main() {

	content := U.ReadFile("./examples/test.swift")

	result := I.NewEvaluator(content)

	for _, err := range result.Errors {
		fmt.Println(err.Error())
	}

	variables := result.Scope.GetSymbolTable()

	fmt.Println()
	for k, v := range variables {
		fmt.Println("Key:", k, "Value:", v.GetValue(), "Is constant:", v.IsConstant(), "Value Type: ", v.GetType())
	}

	// api.Init()
}
