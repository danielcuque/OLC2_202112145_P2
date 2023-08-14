package interfaces

import V "OLC2/chore/values"

// This struct will handle when program founds a return, break or continue

type StackItemType string

const (
	ReturnType   StackItemType = "return"
	BreakType    StackItemType = "break"
	ContinueType StackItemType = "continue"
)

type StackItem struct {
	Name       string
	Value      V.IValue
	ValidProps []StackItemType
	Trigger    StackItemType
}

func (s *StackItem) Contains(prop StackItemType) bool {
	for _, p := range s.ValidProps {
		if p == prop {
			return true
		}
	}
	return false
}

func NewStackItem(name string, value V.IValue, validProps []StackItemType) *StackItem {
	return &StackItem{
		Name:       name,
		Value:      value,
		ValidProps: validProps,
	}
}

type Stack struct {
	Items []*StackItem
}

func NewStack() *Stack {
	return &Stack{
		Items: make([]*StackItem, 0),
	}
}

func (s *Stack) Push(item *StackItem) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() *StackItem {
	if len(s.Items) == 0 {
		return &StackItem{}
	} else {
		item := s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
		return item
	}
}

func (s *Stack) Peek() *StackItem {
	if len(s.Items) == 0 {
		return &StackItem{}
	} else {
		return s.Items[len(s.Items)-1]
	}
}
