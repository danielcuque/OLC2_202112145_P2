package main

import (
	"fmt"

	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"
)

func main() {

	// content := U.ReadFile("./examples/Struct/Struct1.swift")
	// content := U.ReadFile("./examples/Struct/Struct2.swift")
	// content := U.ReadFile("./examples/Basicos/Basicas.swift")
	// content := U.ReadFile("./examples/Basicos/Intermedias.swift")
	content := U.ReadFile("./examples/Basicos/Intermedias.swift")

	result := I.NewEvaluator(content)

	for _, err := range result.Errors {
		fmt.Println(err.Error())
	}

	fmt.Println(result.GetLogs())

	// api.Init()
}
