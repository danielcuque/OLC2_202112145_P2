package compiler

import "fmt"

type Label struct {
	ID int
}

func NewLabel(id int) *Label {
	return &Label{
		ID: id,
	}
}

func (l *Label) String() string {
	return fmt.Sprintf("L%d", l.ID)
}
