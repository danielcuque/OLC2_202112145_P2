package interfaces

type Function struct {
	Name     string
	Params   []string
	DataType string
	Scope    *ScopeNode
}
