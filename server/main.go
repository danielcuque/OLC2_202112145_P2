package main

import (
	"OLC2/api"

	I "OLC2/core/interfaces"
	U "OLC2/core/utils"

	"fmt"
)

func main() {

	content := U.ReadFile("./examples/test.swift")
	// content := U.ReadFile("./examples/Struct/Struct2.swift")
	// content := U.ReadFile("./examples/Basicos/Basicas.swift")
	// content := U.ReadFile("./examples/Basicos/Intermedias.swift")
	// content := U.ReadFile("./examples/Funciones/Embebidas.swift")
	// content := U.ReadFile("./examples/Funciones/Recursivas.swift")

	// content := U.ReadFile("./examples/Arreglos/Vectores.swift")
	// content := U.ReadFile("./examples/Arreglos/Matrices.swift")

	result := I.NewEvaluator(content)

	for _, err := range result.Errors {
		fmt.Println(err.Error())
	}

	fmt.Println(result.GetLogs())

	api.Init()
}
