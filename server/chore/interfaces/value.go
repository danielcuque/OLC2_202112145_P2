package interfaces

const (
	STRING_STR  = "String"
	INT_STR     = "Int"
	FLOAT_STR   = "Float"
	BOOLEAN_STR = "Bool"
	CHAR_STR    = "Character"
	NIL_STR     = "nil"
)

type IValue interface {
	GetValue() interface{}
	GetType() string
}

type StringV struct {
	Value string
}

func (s *StringV) GetValue() interface{} {
	return s.Value
}

func (s *StringV) GetType() string {
	return STRING_STR
}

func NewStringValue(value string) *StringV {
	return &StringV{Value: value}
}

type IntV struct {
	Value int
}

func (i *IntV) GetValue() interface{} {
	return i.Value
}

func (i *IntV) GetType() string {
	return INT_STR
}

func NewIntValue(value int) *IntV {
	return &IntV{Value: value}
}

type FloatV struct {
	Value float64
}

func (f *FloatV) GetValue() interface{} {
	return f.Value
}

func (f *FloatV) GetType() string {
	return FLOAT_STR
}

func NewFloatValue(value float64) *FloatV {
	return &FloatV{Value: value}
}

type BooleanV struct {
	Value bool
}

func (b *BooleanV) GetValue() interface{} {
	return b.Value
}

func (b *BooleanV) GetType() string {
	return BOOLEAN_STR
}

func NewBooleanValue(value bool) *BooleanV {
	return &BooleanV{Value: value}
}

type CharV struct {
	Value rune
}

func (c *CharV) GetValue() interface{} {
	return c.Value
}

func (c *CharV) GetType() string {
	return CHAR_STR
}

func NewCharValue(value rune) *CharV {
	return &CharV{Value: value}
}

type NilV struct {
	Value interface{}
}

func (n *NilV) GetValue() interface{} {
	return n.Value
}

func (n *NilV) GetType() string {
	return NIL_STR
}

func NewNilValue(value interface{}) *NilV {
	return &NilV{Value: value}
}

type ArrayV struct {
	Value []IValue
}

func (a *ArrayV) GetValue() interface{} {
	return a.Value
}

func (a *ArrayV) GetType() string {
	return "array"
}

func NewArrayValue(value []IValue) *ArrayV {
	return &ArrayV{Value: value}
}
