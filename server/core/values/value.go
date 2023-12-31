package values

const (
	StringType  = "String"
	IntType     = "Int"
	FloatType   = "Float"
	BooleanType = "Bool"
	CharType    = "Character"
	NilType     = "nil"

	VectorType = "Vector"
	MatrixType = "Matrix"
)

type IValue interface {
	GetValue() interface{}
	GetType() string
	String() string
	Copy() IValue
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

func (a *RangeV) String() string {
	return ""
}

func (a *RangeV) Copy() IValue {
	return NewRangeValue(a.Value)
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

func IsBaseTypeString(value string) bool {
	switch value {
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
