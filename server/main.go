package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"OLC2/api"
	"OLC2/core/interpreter"
	"OLC2/core/utils"
	"fmt"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mode := os.Getenv("MODE")

	if mode != "DEVELOP" {
		api.Init()
		return
	}

	content := utils.ReadFile("./examples/test.swift")
	compiler, checker := interpreter.NewEvaluator(content)

	for _, err := range checker.Errors {
		fmt.Println(err.Error())
	}

	if compiler != nil {
		fmt.Println(compiler)
	}
}
