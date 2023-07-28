package utils

import (
	"io"
	"os"
	"strconv"
	"strings"
)

// Parse strings to numbers (int or float) if possible
func ParseNumber(str string) interface{} {
	if strings.Contains(str, ".") {
		if f, err := strconv.ParseFloat(str, 64); err == nil {
			return f
		}
	} else {
		if i, err := strconv.ParseInt(str, 10, 64); err == nil {
			return i
		}
	}

	return str
}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	return string(content)
}
