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
	//if comment is not empty
	if comment != "" {
		t.code += fmt.Sprintf("// %s\n", comment)
	}
	t.code += fmt.Sprintf("%s;\n", code)
}

func (t *TAC) GetCode() string {
	return t.code
}

func (t *TAC) NewTemporal(value, castType interface{}) *Temporal {
	temporal := NewTemporal(t.TemporalQuantity(), FloatTemporal)

	if castType != nil {
		temporal = NewTemporal(t.TemporalQuantity(), castType.(TemporalCast))
	}

	t.temporals = append(t.temporals, temporal)
	return temporal
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
