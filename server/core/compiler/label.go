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

type LabelStack struct {
	// Here, we save TrueLabels, FalseLabels, ExitLabels
	TrueLabel  []*Label
	FalseLabel []*Label
	ExitLabel  []*Label
}

func NewLabelStack() *LabelStack {
	return &LabelStack{
		TrueLabel:  []*Label{},
		FalseLabel: []*Label{},
		ExitLabel:  []*Label{},
	}
}

func (l *LabelStack) PushTrueLabel(label *Label) {
	l.TrueLabel = append(l.TrueLabel, label)
}

func (l *LabelStack) PushFalseLabel(label *Label) {
	l.FalseLabel = append(l.FalseLabel, label)
}

func (l *LabelStack) PushExitLabel(label *Label) {
	l.ExitLabel = append(l.ExitLabel, label)
}

func (l *LabelStack) PopTrueLabel() *Label {
	if len(l.TrueLabel) == 0 {
		return nil
	}

	label := l.TrueLabel[len(l.TrueLabel)-1]
	l.TrueLabel = l.TrueLabel[:len(l.TrueLabel)-1]
	return label
}

func (l *LabelStack) PopFalseLabel() *Label {
	if len(l.FalseLabel) == 0 {
		return nil
	}

	label := l.FalseLabel[len(l.FalseLabel)-1]
	l.FalseLabel = l.FalseLabel[:len(l.FalseLabel)-1]
	return label
}

func (l *LabelStack) PopExitLabel() *Label {
	if len(l.ExitLabel) == 0 {
		return nil
	}

	label := l.ExitLabel[len(l.ExitLabel)-1]
	l.ExitLabel = l.ExitLabel[:len(l.ExitLabel)-1]
	return label
}

func (l *LabelStack) GetTrueLabel() *Label {
	if len(l.TrueLabel) == 0 {
		return nil
	}

	return l.TrueLabel[len(l.TrueLabel)-1]
}

func (l *LabelStack) GetFalseLabel() *Label {
	if len(l.FalseLabel) == 0 {
		return nil
	}

	return l.FalseLabel[len(l.FalseLabel)-1]
}

func (l *LabelStack) GetExitLabel() *Label {
	if len(l.ExitLabel) == 0 {
		return nil
	}

	return l.ExitLabel[len(l.ExitLabel)-1]
}

func (l *LabelStack) GetTrueLabelCount() int {
	return len(l.TrueLabel)
}

func (l *LabelStack) GetFalseLabelCount() int {
	return len(l.FalseLabel)
}

func (l *LabelStack) GetExitLabelCount() int {
	return len(l.ExitLabel)
}
