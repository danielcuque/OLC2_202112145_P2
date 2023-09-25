package compiler

import "fmt"

type TAC struct {
	temporals []*Temporal
	code      string
}

func NewTAC() *TAC {
	return &TAC{
		code: "",
	}
}

func (t *TAC) GetTemporals() []*Temporal {
	return t.temporals
}

func (t *TAC) GetTemporalsHeader() string {
	if t.TemporalQuantity() == 0 {
		return ""
	}

	var code string
	code += "float "
	for i := 0; i < t.TemporalQuantity(); i++ {
		// This looks like float t0, t1, t2, t3, t4, t5, t6, t7, t8, t9;
		code += fmt.Sprintf("t%d", i)
		if i != t.TemporalQuantity()-1 {
			code += ", "
		}
	}
	code += ";\n"
	return code
}

func (t *TAC) TemporalQuantity() int {
	return len(t.temporals)
}

func (t *TAC) AppendCode(code string, comment string) {
	t.code += fmt.Sprintf("%s // %s\n", code, comment)
}

func (t *TAC) GetCode() string {
	return t.code
}

func (t *TAC) AppendNewTemporal(comment string, temporal *Temporal) {
	t.AppendCode(fmt.Sprintf("H = H + 1; heap[(int)H] = stack[(int)%s];", temporal.String()), comment)
	t.temporals = append(t.temporals, temporal)
}

func (c *Compiler) GetHeader() string {
	return `
#include <stdio.h>
float stack[100000];
float heap[100000];
float P;
float H;
`
}
