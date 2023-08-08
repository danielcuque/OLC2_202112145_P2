package interfaces

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
	Variables map[string]IValue
}

func NewScopeNode(parent *ScopeNode) *ScopeNode {
	return &ScopeNode{
		Parent:    parent,
		Child:     make([]*ScopeNode, 0),
		Variables: make(map[string]IValue),
	}
}
