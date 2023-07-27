package utils

import (
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
