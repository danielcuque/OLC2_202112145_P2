package main

import (
	I "OLC2/core/interpreter"
	U "OLC2/core/utils"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	content := U.ReadFile("./examples/test.swift")

	compiler, checker := I.NewEvaluator(content)

	for _, err := range checker.Errors {
		fmt.Println(err.Error())
	}

	if compiler != nil {
		tac := compiler.GetTAC()

		fmt.Println(tac)
		err := os.WriteFile(
			"./out/main.c",
			[]byte(tac),
			os.ModePerm,
		)

		if err != nil {
			fmt.Println(err.Error())
		}

		cmd := exec.Command("gcc", "./out/main.c", "-o", "./out/main")
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(string(out))
		} else {
			fmt.Println(string(out))
		}

	}

}
