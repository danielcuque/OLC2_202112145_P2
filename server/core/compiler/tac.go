package compiler

import "fmt"

type TAC struct {
	temporals []*Temporal
	labels    []*Label
	code      string
}

func NewTAC() *TAC {
	return &TAC{
		code:      "",
		temporals: make([]*Temporal, 0),
		labels:    make([]*Label, 0),
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
		code += fmt.Sprintf("t%d", i+1)
		if i != t.TemporalQuantity()-1 {
			code += ", "
		}
	}
	code += ";\n"
	return code
}

func (t *TAC) TemporalQuantity() int {
	return len(t.temporals) + 1
}

func (t *TAC) LabelQuantity() int {
	return len(t.labels) + 1
}

func (t *TAC) NewLabel() *Label {
	label := NewLabel(
		t.LabelQuantity(),
	)
	t.labels = append(t.labels, label)
	return label
}

func (t *TAC) AppendCode(instrucions []string, comment string) {

	if comment != "" {
		t.code += fmt.Sprintf("// %s\n", comment)
	}

	for _, instruction := range instrucions {
		t.code += fmt.Sprintf("%s\n", instruction)
	}
}

func (t *TAC) GetCode() string {
	return t.code
}

func (t *TAC) NewTemporal(value, castType interface{}) *Temporal {
	temporal := NewTemporal(t.TemporalQuantity(), FloatTemporal)

	if castType != nil {
		temporal.Type = castType.(TemporalCast)
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
