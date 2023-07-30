package main

import (
	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"
	"fmt"
)

func main() {

	content := U.ReadFile("./examples/test.swift")

	result := I.NewEvaluator(content)

	fmt.Println(result.Memory)
	fmt.Println(result.Errors)

}
