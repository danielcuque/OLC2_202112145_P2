package values

type CharV struct {
	Value rune
}

func (c *CharV) GetValue() interface{} {
	return c.Value
}

func (c *CharV) GetType() string {
	return CharType
}

func (c *CharV) String() string {
	return string(c.Value)
}

func NewCharValue(value rune) *CharV {
	return &CharV{Value: value}
}

func (c *CharV) Copy() IValue {
	return NewCharValue(c.Value)
}
