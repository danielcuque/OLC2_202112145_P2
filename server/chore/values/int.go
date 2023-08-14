package values

import "strconv"

type IntV struct {
	Value int
}

func (i *IntV) GetValue() interface{} {
	return i.Value
}

func (i *IntV) GetType() string {
	return IntType
}

func (i *IntV) String() string {
	return strconv.Itoa(i.Value)
}

func NewIntValue(value int) *IntV {
	return &IntV{Value: value}
}
