package compiler

import (
	"fmt"
)

const (
	RootEnv   = "Root"
	FuncEnv   = "Func"
	WhileEnv  = "While"
	ForEnv    = "For"
	IfEnv     = "If"
	ElseEnv   = "Else"
	SwitchEnv = "Switch"
	StructEnv = "Struct"
)

type EnvNode struct {
	Parent  *EnvNode
	Child   []*EnvNode
	Level   int
	EnvType string
	Values  map[string]*Value
}

func NewEnvNode(parent *EnvNode, envType string, Level int) *EnvNode {
	return &EnvNode{
		Parent:  parent,
		Child:   make([]*EnvNode, 0),
		Level:   Level,
		EnvType: envType,
	}
}

func (s *EnvNode) GetType() string {
	return string(s.EnvType)
}

func (s *EnvNode) String() string {
	result := ""
	for i := 0; i < s.Level; i++ {
		result += "\t"
	}
	result += fmt.Sprintf("%s\n", s.EnvType)
	for _, v := range s.Child {
		result += v.String()
	}
	return result
}

func (s *EnvNode) Copy() *EnvNode {
	newNode := NewEnvNode(nil, s.EnvType, s.Level)
	for _, child := range s.Child {
		newNode.Child = append(newNode.Child, child.Copy())
	}
	return newNode
}

// EnvTree is a nary tree to represent scopes
type EnvTree struct {
	Root    *EnvNode
	Current *EnvNode
}

func NewEnvTree() *EnvTree {
	root := NewEnvNode(nil, RootEnv, 0)
	return &EnvTree{
		Root:    root,
		Current: root,
	}
}

func (s *EnvTree) PushEnv(scopeType string) *EnvNode {
	node := NewEnvNode(s.Current, scopeType, s.Current.Level+1)
	s.Current.Child = append(s.Current.Child, node)
	s.Current = node
	return node
}

func (s *EnvTree) PopEnv() {
	s.Current = s.Current.Parent
}

func (s *EnvTree) GetSymbolTable() []ApiObject {
	return s.Root.GetAllVariables()
}

func (s *EnvTree) GetCurrentScope() *EnvNode {
	return s.Current
}

func (s *EnvTree) String() string {
	return s.Root.String()
}

func (s *EnvTree) GetMain() string {
	return ""
}
