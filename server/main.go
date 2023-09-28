package main

import (
	I "OLC2/core/interpreter"
	U "OLC2/core/utils"
	"fmt"
)

func main() {

	content := U.ReadFile("./examples/test.swift")

	compiler, checker := I.NewEvaluator(content)

	for _, err := range checker.Errors {
		fmt.Println(err.Error())
	}

	if compiler != nil {
		fmt.Println(compiler.GetTAC())
		// fmt.Println(len(compiler.TAC.GetTemporals()))
	}

}
