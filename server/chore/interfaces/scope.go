package interfaces

import "fmt"

// Use nary tree to represent scopes and values

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
	Variables map[string]Variable
}

func NewScopeNode(parent *ScopeNode, scopeType ScopeType, Level int) *ScopeNode {
	return &ScopeNode{
		Parent:    parent,
		Child:     make([]*ScopeNode, 0),
		Level:     Level,
		ScopeType: scopeType,
		Variables: make(map[string]Variable),
	}
}

func (s *ScopeNode) AddVariable(name string, value Variable) {
	s.Variables[name] = value
}

func (s *ScopeNode) GetVariable(name string) interface{} {
	if val, ok := s.Variables[name]; ok {
		return val
	}
	return nil
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

func (s *ScopeTree) PushScope(scopeType ScopeType) {
	node := NewScopeNode(s.Current, scopeType, s.Current.Level+1)
	s.Current.Child = append(s.Current.Child, node)
	s.Current = node
}

func (s *ScopeTree) PopScope() {
	s.Current = s.Current.Parent
}

func (s *ScopeTree) String() string {
	return s.Root.String()
}
