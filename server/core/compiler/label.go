package compiler

import "fmt"

type Label struct {
	Name string
	ID   int
}

func NewLabel(id int, name string) *Label {
	return &Label{
		ID:   id,
		Name: name,
	}
}

func (l *Label) String() string {
	return fmt.Sprintf("L%d", l.ID)
}

func (l *Label) Declare() string {
	return fmt.Sprintf("%s:", l.String())
}
