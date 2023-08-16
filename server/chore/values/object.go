package values

type ObjectV struct {
	Props  map[string]interface{}
	Method map[string]interface{}
}

func NewObjectV(name string) *ObjectV {
	return &ObjectV{
		Props:  make(map[string]interface{}),
		Method: make(map[string]interface{}),
	}
}

func (o *ObjectV) AddProp(name string, value interface{}) {
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
