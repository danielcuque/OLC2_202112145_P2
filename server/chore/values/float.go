package values

import "strconv"

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
