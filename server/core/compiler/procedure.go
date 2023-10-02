package compiler

import "fmt"

type Procedure struct {
	Name      string
	Labels    map[string]*Label
	Arguments map[string]*Parameter
	Code      string
}

func NewProcedure(name string) *Procedure {
	return &Procedure{
		Name:      name,
		Arguments: make(map[string]*Parameter),
		Labels:    make(map[string]*Label),
	}
}

func (p *Procedure) AddArguments(args []*Parameter) {
	for _, arg := range args {
		p.Arguments[arg.Name] = arg
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

func (p *Procedure) GetArgument(name string) *Parameter {
	return p.Arguments[name]
}

func (p *Procedure) GetLabel(name string) *Label {
	return p.Labels[name]
}

func (p *Procedure) String() string {
	return fmt.Sprintf("void %s(){\n%s \nreturn;\n}\n", p.Name, p.Code)
}

type Parameter struct {
	Temporal *Temporal
	Name     string
}

func (p *Parameter) Tmp() string {
	return p.Temporal.String()
}

type Argument struct {
	Name      string
	Value     *ValueResponse
	IsPointer bool
}
