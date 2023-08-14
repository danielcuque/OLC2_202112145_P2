package values

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
