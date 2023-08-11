package interfaces

import "fmt"

type ScopeType string

const (
	RootScope   ScopeType = "Root"
	FuncScope   ScopeType = "Func"
	WhileScope  ScopeType = "While"
	ForScope    ScopeType = "For"
	IfScope     ScopeType = "If"
	ElseScope   ScopeType = "Else"
	SwitchScope ScopeType = "Switch"
)

type TokenSymbol struct {
	Scope    string
	Type     string
	Name     string
	DataType string
	Value    string
	Params   []string
}

func NewTokenSymbol(scope, tokenType, name, dataType, value string, params []string) *TokenSymbol {
	return &TokenSymbol{
		Scope:    scope,
		Type:     tokenType,
		Name:     name,
		DataType: dataType,
		Value:    value,
		Params:   params,
	}
}

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

func (s *ScopeNode) ResetScopeNode() {
	s.Variables = make(map[string]*Variable)
}

func (s *ScopeNode) GetType() string {
	return string(s.ScopeType)
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

func (s *ScopeTree) GetSymbolTable() []ApiVariable {
	// Traverse tree to get symbol table
	return s.Root.GetAllVariables()
}

func (s *ScopeTree) String() string {
	return s.Root.String()
}

func (s *ScopeTree) ResetScope() {
	// Clean all variables inside current scope
	s.Current.ResetScopeNode()
}

func (s *ScopeTree) GetCurrentScope() *ScopeNode {
	return s.Current
}

// ApiVariable is a struct to represent variables in api
type ApiVariable struct {
	Name    string
	IsConst bool
	Value   interface{}
	Type    string
	Line    int
	Column  int
	Scope   string
}

func (s *ScopeNode) GetAllVariables() []ApiVariable {
	var allVariables []ApiVariable
	s.collectVariables(&allVariables)
	return allVariables
}

func (s *ScopeNode) collectVariables(allVariables *[]ApiVariable) {
	for _, variable := range s.Variables {
		apiVar := ApiVariable{
			Name:    variable.GetName(),
			IsConst: variable.IsConstant(),
			Value:   variable.GetValue(), // Puedes decidir cómo manejar el valor aquí
			Type:    variable.GetType(),
			Line:    variable.GetLine(),
			Column:  variable.GetColumn(),
			Scope:   s.GetType(),
		}
		*allVariables = append(*allVariables, apiVar)
	}
	for _, child := range s.Child {
		child.collectVariables(allVariables)
	}
}
