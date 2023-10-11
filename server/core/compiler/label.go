package compiler

import "fmt"

type LabelFlowType string

const (
	ContinueLabel LabelFlowType = "continue"
	BreakLabel    LabelFlowType = "break"
	ReturnLabel   LabelFlowType = "return"
)

type Label struct {
	Name string
	ID   int
	Type []LabelFlowType
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
