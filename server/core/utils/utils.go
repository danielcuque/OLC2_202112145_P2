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
	return ParseToJSON(ReadFile("grammar/SwiftLexer.g4"))
}

func GetParser() string {
	return ParseToJSON(ReadFile("grammar/Swift.g4"))
}

func FormatInput(input string) string {
	return ParseToJSON(input)
}

func ParseToJSON(input string) string {
	formatted, err := json.Marshal(input)

	if err != nil {
		fmt.Println("Error formatting input:", err)
		return ""
	}

	return string(formatted)
}
