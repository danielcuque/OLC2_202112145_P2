package compiler

import "fmt"

type TAC struct {
	temporals []*Temporal
	labels    []*Label
	standar   map[string]*Procedure
	code      string
	procedure *Procedure
}

func NewTAC() *TAC {
	return &TAC{
		code:      "",
		temporals: make([]*Temporal, 0),
		labels:    make([]*Label, 0),
		standar:   make(map[string]*Procedure),
		procedure: nil,
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

func (c *Compiler) NewLabelFlow(name string, Type []LabelFlowType) *Label {
	label := c.TAC.NewLabel(name)
	label.Type = Type
	c.Env.AddLabel(label)
	return label
}

func (t *TAC) AppendInstructions(instructions []string, comment string) {

	if t.procedure != nil {
		t.procedure.AddCode(instructions, comment)
		return
	}

	if comment != "" {
		t.code += fmt.Sprintf("// %s\n", comment)
	}

	for _, instruction := range instructions {
		t.code += fmt.Sprintf("%s\n", instruction)
	}
}

func (t *TAC) AppendInstruction(instruction string, comment string) {
	t.AppendInstructions([]string{instruction}, comment)
}

func (t *TAC) String() string {
	return t.code
}

func (t *TAC) NewTemporal(castType interface{}) *Temporal {
	temporal := NewTemporal(t.TemporalQuantity()+1, FloatTemporal)

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

func (t *TAC) SetProcedure(procedure *Procedure) {
	t.procedure = procedure
	t.AddProcedure(procedure)
}

func (t *TAC) UnsetProcedure() {
	t.procedure = nil
}

func (t *TAC) GetProcudres() string {
	var code string
	for _, procedure := range t.standar {
		code += fmt.Sprintf("%s\n", procedure)
	}
	return code
}

func (c *Compiler) GetHeader() string {
	return "#include <stdio.h>\nfloat stack[100000];\nfloat heap[100000];\nfloat P;\nfloat H;\n"
}

func (c *Compiler) GenerateVariables() {
	counter := c.StaticVars
	for i := 0; i < counter; i++ {
		c.TAC.AppendInstructions(
			[]string{
				"stack[(int)P] = 0;",
				"P = P + 1;",
			},
			"",
		)
	}
}
