package main

import (
	I "OLC2/core/interfaces"
	U "OLC2/core/utils"

	"fmt"
)

func main() {

	content := U.ReadFile("./examples/test.swift")

	result := I.NewEvaluator(content)

	for _, err := range result.Errors {
		fmt.Println(err.Error())
	}

	fmt.Println(result.GetC3D())

}
