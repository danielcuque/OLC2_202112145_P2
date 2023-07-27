package interfaces

import "OLC2/chore/parser"

type Value struct {
	ParseValue interface{}
}

type Visitor struct {
	parser.SwiftVisitor
	Memory map[string]Value
	Error  []error
}

func NewVisitor() *Visitor {
	return &Visitor{
		Memory: make(map[string]Value),
		Error:  make([]error, 0),
	}
}
