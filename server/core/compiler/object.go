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

func (o *Object) AddProp(name string, value *Value) {
	o.Props.AddValue(name, value)
}

func (o *Object) GetProp(name string) *Value {
	return o.Props.Values[name]
}
