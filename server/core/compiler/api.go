package compiler

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
