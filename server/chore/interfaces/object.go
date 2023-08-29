package interfaces

import (
	V "OLC2/chore/values"
	"fmt"
)

type ObjectV struct {
	Type string
	Body V.IValue
	Env  *EnvNode
}

func NewObjectV(Type string, body V.IValue, Env *EnvNode) *ObjectV {

	return &ObjectV{
		Type: Type,
		Body: body,
		Env:  Env,
	}
}

func (o *ObjectV) AddProp(name string, value *Variable) {
	o.Env.Variables[name] = value
}

func (o *ObjectV) AddMethod(name string, value interface{}) {
	o.Env.Functions[name] = value.(*Function)
}

func (o *ObjectV) GetProp(name string) interface{} {
	return o.Env.Variables[name]
}

func (o *ObjectV) GetMethod(name string) interface{} {
	return o.Env.Functions[name]
}

func (o *ObjectV) SetPropValue(name string, value V.IValue) {
	o.Env.Variables[name].SetValue(value)
}

func (o *ObjectV) GetType() string {
	return o.Type
}

func (o *ObjectV) GetValue() interface{} {
	return o.Body
}

func (o *ObjectV) String() string {
	return fmt.Sprint(o.Env.Variables, &o)
}

// Recursive function to get the value of a property

func GetPropValue(variable *Variable, props []string) interface{} {
	// There are two cases, variable can store another object, or not

	obj, ok := GetObject(variable).(*ObjectV)

	if !ok {
		return nil
	}

	if len(props) == 1 {
		return obj.GetProp(props[0])
	}

	return GetPropValue(obj.GetProp(props[0]).(*Variable), props[1:])
}

func GetObjectPropValue(object *ObjectV, props []string) interface{} {
	if len(props) == 1 {
		return object.GetProp(props[0])
	}

	return GetObjectPropValue(object.GetProp(props[0]).(*ObjectV), props[1:])
}

// Recursive until get the value of the object, then return the value
func GetObject(variable *Variable) interface{} {
	if CheckIsPointer(variable.Value) {
		return GetObject(variable.Value.(*Variable))
	}
	return variable.Value
}

func (o *ObjectV) Copy() V.IValue {
	return NewObjectV(o.Type, o.Body.Copy(), o.Env.Copy())
}
