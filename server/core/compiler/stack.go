package compiler

import "fmt"

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

func (s *Stack) SubPointer() int {
	s.Pointer--
	return s.Pointer
}

func (s *Stack) IncreasePointer(amount int) string {
	return fmt.Sprintf("P = P + %d;", amount)
}

func (s *Stack) IncreasePointerByOne() string {
	s.Pointer++
	return "P = P + 1;"
}

type StackIndex struct {
	Index interface{}
}

func NewStackIndex(Index interface{}) *StackIndex {
	return &StackIndex{
		Index: Index,
	}
}

func (s *StackIndex) String() string {
	return fmt.Sprintf("%v", s.Index)
}
