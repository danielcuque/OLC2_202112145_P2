package compiler

import "fmt"

type Procedure struct {
	Name      string
	Labels    map[string]*Label
	Arguments map[string]*Argument
	Code      string
}

func NewProcedure(name string) *Procedure {
	return &Procedure{
		Name:      name,
		Arguments: make(map[string]*Argument),
		Labels:    make(map[string]*Label),
	}
}

func (p *Procedure) AddArguments(args []*Argument) {
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

func (p *Procedure) GetArgument(name string) *Argument {
	return p.Arguments[name]
}

func (p *Procedure) GetLabel(name string) *Label {
	return p.Labels[name]
}

func (p *Procedure) String() string {
	return fmt.Sprintf("void %s(){\n%s \nreturn;\n}", p.Name, p.Code)
}

type Argument struct {
	Temporal *Temporal
	Name     string
}

func (a *Argument) Tmp() string {
	return a.Temporal.String()
}
