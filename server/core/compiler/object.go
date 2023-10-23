package compiler

type Object struct {
	Type  string
	Value interface{}
	Props *EnvNode
}

func NewObject(t string, v interface{}, env *EnvNode) *Object {
	return &Object{
		Type:  t,
		Value: v,
		Props: env,
	}
}
