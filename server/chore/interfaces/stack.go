package interfaces

// This struct will handle when program founds a return, break or continue

type StackItemType string

const (
	Return   StackItemType = "return"
	Break    StackItemType = "break"
	Continue StackItemType = "continue"
)

type StackItem struct {
	ItemType []StackItemType
	Value    IValue
	Function string
}

func NewStackItem(itemType []StackItemType, value IValue, function string) *StackItem {
	return &StackItem{
		ItemType: itemType,
		Value:    value,
		Function: function,
	}
}

type Stack struct {
	Items []StackItem
}

func NewStack() *Stack {
	return &Stack{
		Items: make([]StackItem, 0),
	}
}

func (s *Stack) Push(item StackItem) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() StackItem {
	if len(s.Items) == 0 {
		return StackItem{}
	} else {
		item := s.Items[len(s.Items)-1]
		s.Items = s.Items[:len(s.Items)-1]
		return item
	}
}

func (s *Stack) Peek() StackItem {
	if len(s.Items) == 0 {
		return StackItem{}
	} else {
		return s.Items[len(s.Items)-1]
	}
}

func (s *Stack) Remove(index int) {
	s.Items = append(s.Items[:index], s.Items[index+1:]...)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func (s *Stack) Clear() {
	s.Items = []StackItem{}
}
