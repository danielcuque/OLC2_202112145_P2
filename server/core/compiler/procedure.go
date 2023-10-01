package compiler

import "fmt"

type Procedure struct {
	Name      string
	Arguments map[string]*Argument
	Code      string
}

func NewProcedure(name string) *Procedure {
	return &Procedure{
		Name:      name,
		Arguments: make(map[string]*Argument),
	}
}

func (p *Procedure) AddArguments(args []*Argument) {
	// Append arguments
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

func (p *Procedure) GetArgumentByName(name string) *Argument {
	for _, arg := range p.Arguments {
		if arg.Name == name {
			return arg
		}
	}
	return nil
}

func (p *Procedure) String() string {
	return fmt.Sprintf("void %s(){\n %s \nreturn;\n}", p.Name, p.Code)
}

type Argument struct {
	Temporal *Temporal
	Name     string
}

func (a *Argument) Tmp() string {
	return a.Temporal.String()
}
