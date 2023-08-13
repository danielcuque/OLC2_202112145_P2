package main

import (
	"OLC2/api"
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
	for _, v := range variables {
		fmt.Println(v)
	}

	api.Init()
}
