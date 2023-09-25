package c3d

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

type TokenSymbol struct {
	Env      string
	Type     string
	Name     string
	DataType string
	Value    string
	Params   []string
}

func NewTokenSymbol(env, tokenType, name, dataType, value string, params []string) *TokenSymbol {
	return &TokenSymbol{
		Env:      env,
		Type:     tokenType,
		Name:     name,
		DataType: dataType,
		Value:    value,
		Params:   params,
	}
}

type EnvNode struct {
	Parent    *EnvNode
	Child     []*EnvNode
	Level     int
	ScopeType string
	Values    map[string]interface{}
}

func NewEnvNode(parent *EnvNode, envType string, Level int) *EnvNode {
	return &EnvNode{
		Parent:    parent,
		Child:     make([]*EnvNode, 0),
		Level:     Level,
		ScopeType: envType,
	}
}

func (s *EnvNode) GetType() string {
	return string(s.ScopeType)
}

func (s *EnvNode) String() string {
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

func (s *EnvNode) Copy() *EnvNode {
	newNode := NewEnvNode(nil, s.ScopeType, s.Level)
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
	// Traverse tree to get symbol table
	return s.Root.GetAllVariables()
}

func (s *EnvTree) GetCurrentScope() *EnvNode {
	return s.Current
}

func (s *EnvTree) String() string {
	return s.Root.String()
}

// ApiObject is a struct to represent variables in api
type ApiObject struct {
	Name       string
	Value      interface{}
	Type       string
	Line       int
	Column     int
	Scope      string
	Params     string
	ReturnType string
}

func NewApiObject(name string, value interface{}, tokenType string, line, column int, scope string, params string, returnType string) ApiObject {
	return ApiObject{
		Name:       name,
		Value:      value,
		Type:       tokenType,
		Line:       line,
		Column:     column,
		Scope:      scope,
		Params:     params,
		ReturnType: returnType,
	}
}

func (s *EnvNode) GetAllVariables() []ApiObject {
	allVariables := make([]ApiObject, 0)
	// s.collectObjects(&allVariables)
	return allVariables
}
