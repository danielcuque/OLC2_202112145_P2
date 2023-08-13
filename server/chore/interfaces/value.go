package interfaces

const (
	StringType  = "String"
	IntType     = "Int"
	FloatType   = "Float"
	BooleanType = "Bool"
	CharType    = "Character"
	NilType     = "nil"
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
	return StringType
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
	return IntType
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
	return FloatType
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
	return BooleanType
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
	return CharType
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
	return NilType
}

func NewNilValue(value interface{}) *NilV {
	return &NilV{Value: value}
}

type RangeV struct {
	Value []IValue
}

func (a *RangeV) GetValue() interface{} {
	return a.Value
}

func (a *RangeV) GetType() string {
	return IntType
}

func NewRangeValue(value []IValue) *RangeV {
	return &RangeV{Value: value}
}
