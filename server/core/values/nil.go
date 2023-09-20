package values

type NilV struct {
	Value interface{}
}

func (n *NilV) GetValue() interface{} {
	return n.Value
}

func (n *NilV) GetType() string {
	return NilType
}

func (n *NilV) String() string {
	return "nil"
}

func NewNilValue(value interface{}) *NilV {
	return &NilV{Value: value}
}

func (n *NilV) Copy() IValue {
	return NewNilValue(n.Value)
}
