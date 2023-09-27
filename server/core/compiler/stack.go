package compiler

type Stack struct {
	Pointer int
}

func NewStack() *Stack {
	return &Stack{
		Pointer: 0,
	}
}

func (s *Stack) GetPointer() int {
	return s.Pointer
}

func (s *Stack) AddPointer() int {
	s.Pointer++
	return s.Pointer
}

func (s *Stack) SubPointer() int {
	s.Pointer--
	return s.Pointer
}
