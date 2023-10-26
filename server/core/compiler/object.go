package compiler

type Object struct {
	Type  string
	Value interface{}
	Props *EnvNode
}

func NewObject(t string, v interface{}) *Object {
	return &Object{
		Type:  t,
		Value: v,
		Props: NewEnvNode(nil, t),
	}
}

func (o *Object) GetValue() interface{} {
	return o.Value
}

func (o *Object) AddProp(name string, value *Value) {
	o.Props.AddValue(name, value)
}

func (o *Object) GetProp(name string) *Value {
	return o.Props.Values[name]
}

func (o *Object) GetType() string {
	return o.Type
}
