package main

import "OLC2/api"

// import (
// 	I "OLC2/chore/interfaces"
// 	U "OLC2/chore/utils"
// 	"fmt"
// )

func main() {

	// content := U.ReadFile("./examples/Struct/Struct1.swift")
	// content := U.ReadFile("./examples/Struct/Struct2.swift")
	// content := U.ReadFile("./examples/Basicos/Basicas.swift")
	// content := U.ReadFile("./examples/Basicos/Intermedias.swift")
	// content := U.ReadFile("./examples/Funciones/Embebidas.swift")
	// content := U.ReadFile("./examples/Funciones/Recursivas.swift")

	// content := U.ReadFile("./examples/Arreglos/Vectores.swift")
	// content := U.ReadFile("./examples/Arreglos/Matrices.swift")

	// result := I.NewEvaluator(content)

	// for _, err := range result.Errors {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(result.GetLogs())

	api.Init()
}
