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

	// variables := result.Scope.GetSymbolTable()

	// fmt.Println()
	// for _, v := range variables {
	// 	fmt.Println(v)
	// }

	fmt.Println(result.GetLogs())

	// api.Init()
}
