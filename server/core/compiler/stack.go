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
	s.Pointer += amount
	return fmt.Sprintf("P = P + %d;", s.Pointer)
}

func (s *Stack) IncreasePointerByOne() string {
	s.Pointer++
	return "P = P + 1;"
}
