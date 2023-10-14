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
	GuardEnv  = "Guard"
	SwitchEnv = "Switch"
	StructEnv = "Struct"
)

type EnvNode struct {
	Parent     *EnvNode
	Child      []*EnvNode
	EnvType    string
	Values     map[string]*Value
	FlowLabels []*Label
}

func NewEnvNode(parent *EnvNode, envType string) *EnvNode {
	return &EnvNode{
		Parent:  parent,
		Child:   make([]*EnvNode, 0),
		EnvType: envType,
	}
}

func (s *EnvNode) GetType() string {
	return string(s.EnvType)
}

func (s *EnvNode) String() string {
	result := ""

	result += fmt.Sprintf("%s\n", s.EnvType)
	for _, v := range s.Child {
		result += v.String()
	}
	return result
}

func (s *EnvNode) Copy() *EnvNode {
	newNode := NewEnvNode(nil, s.EnvType)
	for _, child := range s.Child {
		newNode.Child = append(newNode.Child, child.Copy())
	}
	return newNode
}

func (s *EnvNode) AddValue(key string, value *Value) {
	if s.Values == nil {
		s.Values = make(map[string]*Value)
	}
	s.Values[key] = value
}

func (s *EnvNode) AddLabel(value *Label) {
	s.FlowLabels = append(s.FlowLabels, value)
}

func (s *EnvNode) GetLabel(labelType LabelFlowType) *Label {
	for node := s; node != nil; node = node.Parent {
		for _, label := range node.FlowLabels {
			for _, lblType := range label.Type {
				if lblType == labelType {
					return label
				}
			}
		}
	}

	return nil
}

// EnvTree is a nary tree to represent scopes
type EnvTree struct {
	Root     *EnvNode
	Current  *EnvNode
	StackEnv []*EnvNode
}

func NewEnvTree() *EnvTree {
	root := NewEnvNode(nil, RootEnv)
	return &EnvTree{
		Root:    root,
		Current: root,
	}
}

func (s *EnvTree) PushEnv(scopeType string) *EnvNode {
	node := NewEnvNode(s.Current, scopeType)
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

func (s *EnvTree) AddValue(key string, value *Value) {
	s.Current.AddValue(key, value)
}

func (s *EnvTree) GetValue(key string) *Value {
	for node := s.Current; node != nil; node = node.Parent {
		if value, ok := node.Values[key]; ok {
			return value
		}
	}
	return nil
}

func (s *EnvTree) AddLabel(value *Label) {
	s.Current.AddLabel(value)
}

func (s *EnvTree) GetLabel(labelType LabelFlowType) *Label {
	return s.Current.GetLabel(labelType)
}

func (s *EnvTree) GetMain() string {
	return ""
}

func (s *EnvTree) Next() {
	if len(s.StackEnv) == 0 {
		// Si la pila está vacía, no hay más nodos para recorrer.
		return
	}

	// Toma el primer nodo de la pila.
	node := s.StackEnv[0]
	s.StackEnv = s.StackEnv[1:]

	// Actualiza el nodo actual.
	s.Current = node

	// Agrega los hijos del nodo actual a la pila (en el orden deseado).
	s.StackEnv = append(s.StackEnv, node.Child...)
}
