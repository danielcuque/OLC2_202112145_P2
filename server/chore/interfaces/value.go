package interfaces

const (
	STRING_STR  = "string"
	INT_STR     = "int"
	FLOAT_STR   = "float"
	BOOLEAN_STR = "boolean"
	CHAR_STR    = "character"
	NIL_STR     = "nil"
)

type IValue interface {
	GetValue() interface{}
	GetType() string
}

type StringValue struct {
	Value string
}

func (s *StringValue) GetValue() interface{} {
	return s.Value
}

func (s *StringValue) GetType() string {
	return STRING_STR
}

func NewStringValue(value string) *StringValue {
	return &StringValue{Value: value}
}

type IntValue struct {
	Value int
}

func (i *IntValue) GetValue() interface{} {
	return i.Value
}

func (i *IntValue) GetType() string {
	return INT_STR
}

func NewIntValue(value int) *IntValue {
	return &IntValue{Value: value}
}

type FloatValue struct {
	Value float64
}

func (f *FloatValue) GetValue() interface{} {
	return f.Value
}

func (f *FloatValue) GetType() string {
	return FLOAT_STR
}

func NewFloatValue(value float64) *FloatValue {
	return &FloatValue{Value: value}
}

type BooleanValue struct {
	Value bool
}

func (b *BooleanValue) GetValue() interface{} {
	return b.Value
}

func (b *BooleanValue) GetType() string {
	return BOOLEAN_STR
}

func NewBooleanValue(value bool) *BooleanValue {
	return &BooleanValue{Value: value}
}

type CharValue struct {
	Value rune
}

func (c *CharValue) GetValue() interface{} {
	return c.Value
}

func (c *CharValue) GetType() string {
	return CHAR_STR
}

func NewCharValue(value rune) *CharValue {
	return &CharValue{Value: value}
}

type NilValue struct {
	Value interface{}
}

func (n *NilValue) GetValue() interface{} {
	return n.Value
}

func (n *NilValue) GetType() string {
	return NIL_STR
}

func NewNilValue(value interface{}) *NilValue {
	return &NilValue{Value: value}
}

type Value struct {
	Type       string
	ParseValue interface{}
}

type Variable struct {
	Name       string
	Value      Value
	IsConstant bool
}
