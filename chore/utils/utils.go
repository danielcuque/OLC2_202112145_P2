package utils

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Parse strings to numbers (int or float)
func ParseNumber(str string) interface{} {
	if strings.Contains(str, ".") {
		f, _ := strconv.ParseFloat(str, 64)
		return f
	} else {
		i, _ := strconv.ParseInt(str, 10, 64)
		return i
	}
}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	return string(content)
}
