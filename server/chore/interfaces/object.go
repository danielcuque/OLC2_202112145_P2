package interfaces

import (
	V "OLC2/chore/values"
)

type ObjectV struct {
	Type   string
	Body   V.IValue
	Props  map[string]*Variable
	Method map[string]interface{}
}

func NewObjectV(Type string, body V.IValue) *ObjectV {
	return &ObjectV{
		Type:   Type,
		Body:   body,
		Props:  make(map[string]*Variable),
		Method: make(map[string]interface{}),
	}
}

func (o *ObjectV) AddProp(name string, value *Variable) {
	o.Props[name] = value
}

func (o *ObjectV) AddMethod(name string, value interface{}) {
	o.Method[name] = value
}

func (o *ObjectV) GetProp(name string) interface{} {
	return o.Props[name]
}

func (o *ObjectV) GetMethod(name string) interface{} {
	return o.Method[name]
}

func (o *ObjectV) GetType() string {
	return o.Type
}

func (o *ObjectV) GetValue() interface{} {
	return o.Body
}

func (o *ObjectV) String() string {
	return o.Body.String()
}

// Recursive function to get the value of a property

func GetPropValue(variable *Variable, props []string) interface{} {
	// There are two cases, variable can store another object, or not

	obj, ok := variable.Value.(*ObjectV)

	if !ok {
		return nil
	}

	if len(props) == 1 {
		return obj.GetProp(props[0])
	}

	return GetPropValue(obj.GetProp(props[0]).(*Variable), props[1:])
}
