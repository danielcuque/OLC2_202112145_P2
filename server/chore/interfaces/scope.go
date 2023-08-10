package interfaces

import "fmt"

type ScopeType string

const (
	RootScope  ScopeType = "Root"
	FuncScope  ScopeType = "Func"
	WhileScope ScopeType = "While"
	ForScope   ScopeType = "For"
	IfScope    ScopeType = "If"
)

type ScopeNode struct {
	Parent    *ScopeNode
	Child     []*ScopeNode
	Level     int
	ScopeType ScopeType
	Variables map[string]*Variable
}

func NewScopeNode(parent *ScopeNode, scopeType ScopeType, Level int) *ScopeNode {
	return &ScopeNode{
		Parent:    parent,
		Child:     make([]*ScopeNode, 0),
		Level:     Level,
		ScopeType: scopeType,
		Variables: make(map[string]*Variable),
	}
}

func (s *ScopeNode) AddVariable(name string, value *Variable) {
	s.Variables[name] = value
}

func (s *ScopeNode) GetVariable(name string) interface{} {
	if val, ok := s.Variables[name]; ok {
		return val
	}
	return nil
}

func (s *ScopeNode) SetVariable(name string, value IValue) {
	if val, ok := s.Variables[name]; ok {
		val.SetValue(value)
	}
}

// Get variables from current scope and all children scopes
func (s *ScopeNode) GetVariables() map[string]*Variable {
	result := make(map[string]*Variable)
	for k, v := range s.Variables {
		result[k] = v
	}
	for _, v := range s.Child {
		for k, v := range v.GetVariables() {
			result[k] = v
		}
	}
	return result
}

func (s *ScopeNode) String() string {
	result := ""
	for i := 0; i < s.Level; i++ {
		result += "\t"
	}
	result += fmt.Sprintf("%s\n", s.ScopeType)
	for _, v := range s.Child {
		result += v.String()
	}
	return result
}

// ScopeTree is a nary tree to represent scopes
type ScopeTree struct {
	Root    *ScopeNode
	Current *ScopeNode
}

func NewScopeTree() *ScopeTree {
	root := NewScopeNode(nil, RootScope, 0)
	return &ScopeTree{
		Root:    root,
		Current: root,
	}
}

func (s *ScopeTree) AddVariable(name string, value *Variable) {
	s.Current.AddVariable(name, value)
}

func (s *ScopeTree) GetVariable(name string) interface{} {
	// Check if var exists in current scope
	// if not, check in parent scope until root

	for node := s.Current; node != nil; node = node.Parent {
		if val, ok := node.Variables[name]; ok {
			return val
		}
	}
	return nil
}

func (s *ScopeTree) SetVariable(name string, value IValue) {
	s.Current.SetVariable(name, value)
}

func (s *ScopeTree) PushScope(scopeType ScopeType) *ScopeNode {
	node := NewScopeNode(s.Current, scopeType, s.Current.Level+1)
	s.Current.Child = append(s.Current.Child, node)
	s.Current = node
	return node
}

func (s *ScopeTree) PopScope() {
	s.Current = s.Current.Parent
}

func (s *ScopeTree) GetSymbolTable() *ScopeNode {
	// Traverse tree to get symbol table
	return s.Root
}

func (s *ScopeTree) String() string {
	return s.Root.String()
}
