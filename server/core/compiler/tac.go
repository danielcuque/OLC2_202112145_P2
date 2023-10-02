package compiler

import "fmt"

type TAC struct {
	temporals []*Temporal
	labels    []*Label
	standar   map[string]*Procedure
	code      string
}

func NewTAC() *TAC {
	return &TAC{
		code:      "",
		temporals: make([]*Temporal, 0),
		labels:    make([]*Label, 0),
		standar:   make(map[string]*Procedure),
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
	return code + "\n"
}

func (t *TAC) TemporalQuantity() int {
	return len(t.temporals)
}

func (t *TAC) LabelQuantity() int {
	return len(t.labels) + 1
}

func (t *TAC) NewLabel(name string) *Label {
	label := NewLabel(
		t.LabelQuantity(),
		name,
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

func (t *TAC) String() string {
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

func (t *TAC) AddProcedure(procedure *Procedure) {
	t.standar[procedure.Name] = procedure
}

func (t *TAC) GetStandar(name string) *Procedure {
	return t.standar[name]
}

func (t *TAC) GetProcudres() string {
	var code string
	for _, procedure := range t.standar {
		code += fmt.Sprintf("%s\n", procedure)
	}
	return code
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
