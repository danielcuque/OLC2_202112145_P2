package c3d

import "fmt"

// Generate code like this

/*
#include <stdio.h>
float stack[100000];
float heap[100000];
float P;
float H;
*/

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
	temps := "float "

	for i := 0; i < n; i++ {
		temps += fmt.Sprintf("t%d", i)
		if i != n-1 {
			temps += ","
		}
	}

	return temps
}
