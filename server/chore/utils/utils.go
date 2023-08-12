package utils

import (
	"io"
	"os"
	"reflect"
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

func GetType(v interface{}) string {
	return reflect.TypeOf(v).String()
}
