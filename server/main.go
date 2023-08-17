package main

import (
	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"
	"fmt"
)

func main() {

	content := U.ReadFile("./examples/recursivas.swift")

	result := I.NewEvaluator(content)

	for _, err := range result.Errors {
		fmt.Println(err.Error())
	}

	fmt.Println(result.GetLogs())

	// api.Init()
}
