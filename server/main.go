package main

// import "OLC2/api"

// func main() {
// 	api.Init()
// }

import (
	"OLC2/core/interpreter"
	"OLC2/core/utils"
	"fmt"
)

func main() {

	content := utils.ReadFile("./examples/test.swift")
	compiler, checker := interpreter.NewEvaluator(content)

	for _, err := range checker.Errors {
		fmt.Println(err.Error())
	}

	if compiler != nil {
		fmt.Println(compiler)
	}

}
