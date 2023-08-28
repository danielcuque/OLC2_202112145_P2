package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	return string(content)
}

func GetLexer() string {
	lexer, err := json.Marshal(ReadFile("grammar/SwiftLexer.g4"))
	if err != nil {
		fmt.Println("Error marshalling lexer:", err)
		return ""
	}

	return string(lexer)
}

func GetParser() string {
	parser, err := json.Marshal(ReadFile("grammar/Swift.g4"))
	if err != nil {
		fmt.Println("Error marshalling parser:", err)
		return ""
	}

	return string(parser)
}

func FormatInput(input string) string {
	formatted, err := json.Marshal(input)

	if err != nil {
		fmt.Println("Error formatting input:", err)
		return ""
	}

	return string(formatted)
}
