package interfaces

import (
	"fmt"
)

func GetHeader() string {
	return `
#include <stdio.h>
float stack[100000];
float heap[100000];
float P;
float H;
`
}

func GetTemps(n int) string {
	if n == 0 {
		return ""
	}

	temps := "float "

	for i := 0; i < n; i++ {
		temps += fmt.Sprintf("t%d", i)
		if i != n-1 {
			temps += ","
		}
	}

	return temps
}

// Todo:

/*
	printString
	castIntToFloat
	castFloatToInt
	castIntToString
	castFloatToString
*/
