package interfaces

// Use nary tree to represent scopes and values

type ScopeNode struct {
	Parent    *ScopeNode
	Child     []*ScopeNode
	Level     int
	Variables map[string]Value
}

func NewScopeNode(parent *ScopeNode) *ScopeNode {
	return &ScopeNode{
		Parent:    parent,
		Child:     make([]*ScopeNode, 0),
		Variables: make(map[string]Value),
	}
}

func (s *ScopeNode) DeclareVariable(name string, value Value) {
	s.Variables[name] = value
}
