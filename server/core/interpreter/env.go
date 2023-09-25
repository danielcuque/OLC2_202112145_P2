package interfaces

import (
	V "OLC2/core/values"
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
	Variables map[string]*Variable
	Functions map[string]*Function
	Structs   map[string]*ObjectV
}

func NewEnvNode(parent *EnvNode, envType string, Level int) *EnvNode {
	return &EnvNode{
		Parent:    parent,
		Child:     make([]*EnvNode, 0),
		Level:     Level,
		ScopeType: envType,
		Variables: make(map[string]*Variable),
		Functions: make(map[string]*Function),
		Structs:   make(map[string]*ObjectV),
	}
}

func (s *EnvNode) AddVariable(name string, value *Variable) {
	s.Variables[name] = value
}

func (s *EnvNode) GetVariable(name string) *Variable {
	if val, ok := s.Variables[name]; ok {
		return val
	}
	return nil
}

func (s *EnvNode) SetVariable(name string, value V.IValue) {
	if val, ok := s.Variables[name]; ok {
		val.SetValue(value)
	}
}

func (s *EnvNode) ResetEnvNode() {
	s.Variables = make(map[string]*Variable)
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
	for key, variable := range s.Variables {
		newNode.Variables[key] = variable.CopyVar()
	}
	for key, function := range s.Functions {
		newNode.Functions[key] = function
	}
	for key, object := range s.Structs {
		newNode.Structs[key] = object.Copy().(*ObjectV)
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

func (s *EnvTree) AddVariable(name string, value *Variable) {
	s.Current.AddVariable(name, value)
}

func (s *EnvTree) GetVariable(name string) *Variable {
	for node := s.Current; node != nil; node = node.Parent {
		if val, ok := node.Variables[name]; ok {
			return val
		}
	}

	return nil
}

func (s *EnvTree) SetVariable(name string, value V.IValue) {
	s.Current.SetVariable(name, value)
}

func (s *EnvTree) AddFunction(name string, value *Function) {
	s.Current.Functions[name] = value
}

func (s *EnvTree) GetFunction(name string) *Function {
	for node := s.Current; node != nil; node = node.Parent {
		if _, ok := node.Functions[name]; ok {
			return node.Functions[name]
		}
	}
	return nil
}

func (s *EnvTree) AddStruct(name string, value *ObjectV) {
	s.Current.Structs[name] = value
}

func (s *EnvTree) GetStruct(name string) *ObjectV {
	for node := s.Current; node != nil; node = node.Parent {
		if _, ok := node.Structs[name]; ok {
			return node.Structs[name]
		}
	}
	return nil
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

func (s *EnvTree) ResetEnv() {
	// Clean all variables inside current scope
	s.Current.ResetEnvNode()
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
	s.collectObjects(&allVariables)
	return allVariables
}

func (s *EnvNode) collectObjects(allVariables *[]ApiObject) {
	for _, variable := range s.Variables {
		apiVar := NewApiObject(
			variable.GetName(),
			variable.String(),
			variable.GetType(),
			variable.GetLine(),
			variable.GetColumn(),
			s.GetType(),
			"...",
			"...",
		)
		*allVariables = append(*allVariables, apiVar)
	}

	for _, function := range s.Functions {
		apiVar := NewApiObject(
			function.GetName(),
			"...",
			"Function",
			function.Line,
			function.Column,
			"...",
			function.GetParamsString(),
			function.ReturnDataType,
		)
		*allVariables = append(*allVariables, apiVar)
	}

	for _, child := range s.Child {
		child.collectObjects(allVariables)
	}
}
