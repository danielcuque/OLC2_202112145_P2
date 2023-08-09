package main

import (
	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"

	"fmt"
)

func main() {

	content := U.ReadFile("./examples/test.swift")

	result := I.NewEvaluator(content)

	fmt.Println(result.Errors)

	variables := result.Scope.Current.Variables

	for k, v := range variables {
		fmt.Println("Key:", k, "Value:", v.GetValue(), "Type:", v.IsConstant(), "Value Type: ", v.GetType())
	}

}
