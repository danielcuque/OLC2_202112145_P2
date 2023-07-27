package interfaces

import "OLC2/chore/parser"

type Value struct {
	ParseValue interface{}
}

type Visitor struct {
	parser.SwiftVisitor
	Memory map[string]Value
}
