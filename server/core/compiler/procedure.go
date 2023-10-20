package compiler

import "fmt"

type Procedure struct {
	Name           string
	Labels         map[string]*Label
	Parameters     map[string]*Parameter
	Code           string
	Env            *EnvTree
	ReturnTemporal *Temporal
	ReturnLabel    *Label
}

func NewProcedure(name string) *Procedure {
	return &Procedure{
		Name:       name,
		Parameters: make(map[string]*Parameter),
		Labels:     make(map[string]*Label),
	}
}

func (p *Procedure) AddParameters(args []*Parameter) {
	for _, arg := range args {
		p.Parameters[arg.ExternalName] = arg
	}
}

func (p *Procedure) AddCode(instructions []string, comment string) {
	if comment != "" {
		p.Code += fmt.Sprintf("// %s\n", comment)
	}

	for _, instruction := range instructions {
		p.Code += fmt.Sprintf("%s\n", instruction)
	}
}

func (p *Procedure) AddLabels(labels []*Label) {
	for _, label := range labels {
		p.Labels[label.Name] = label
	}
}

func (p *Procedure) GetParameter(name string) *Parameter {
	return p.Parameters[name]
}

func (p *Procedure) GetLabel(name string) *Label {
	return p.Labels[name]
}

func (p *Procedure) String() string {
	return fmt.Sprintf("void %s(){\n%s \nreturn;\n}\n", p.Name, p.Code)
}

func (p *Procedure) SetReturnProps(returnTemporal *Temporal, label *Label) {
	p.ReturnTemporal = returnTemporal
	p.ReturnLabel = label
}

func (p *Procedure) ParamSize() int {
	return len(p.Parameters)
}

type Parameter struct {
	Temporal     *Temporal
	ExternalName string
	InternalName string
	IsReference  bool
	Type         TemporalCast
}

func (p *Parameter) Tmp() string {
	return p.Temporal.String()
}

type Argument struct {
	Name      string
	Value     *ValueResponse
	IsPointer bool
}
