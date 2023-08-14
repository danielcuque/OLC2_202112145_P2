package values

import "strconv"

type BooleanV struct {
	Value bool
}

func (b *BooleanV) GetValue() interface{} {
	return b.Value
}

func (b *BooleanV) GetType() string {
	return BooleanType
}

func (b *BooleanV) String() string {
	return strconv.FormatBool(b.Value)
}

func NewBooleanValue(value bool) *BooleanV {
	return &BooleanV{Value: value}
}
