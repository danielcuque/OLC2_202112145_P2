package main

import (
	A "OLC2/api"
	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"

	"fmt"
)

func main() {

	content := U.ReadFile("./examples/test.swift")

	result := I.NewEvaluator(content)

	fmt.Println(result.Memory)
	fmt.Println(result.Errors)

	A.Init()

}
