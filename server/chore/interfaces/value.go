package interfaces

import "strconv"

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
	String() string
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

func (s *StringV) String() string {
	return s.Value
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

func (i *IntV) String() string {
	return strconv.Itoa(i.Value)
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

func (f *FloatV) String() string {
	// format to 4 decimal places
	return strconv.FormatFloat(f.Value, 'f', 4, 64)
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

func (b *BooleanV) String() string {
	return strconv.FormatBool(b.Value)
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

func (c *CharV) String() string {
	return string(c.Value)
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

func (n *NilV) String() string {
	return "nil"
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

func IsBaseType(value IValue) bool {
	switch value.GetType() {
	case StringType:
		return true
	case IntType:
		return true
	case FloatType:
		return true
	case BooleanType:
		return true
	case CharType:
		return true
	case NilType:
		return true
	default:
		return false
	}
}
