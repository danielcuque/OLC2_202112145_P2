package interfaces

type ValueType int

const (
	STRING ValueType = iota
	INT
	DOUBLE
	BOOL
)

const (
	STRING_STR = "STRING"
	INT_STR    = "INT"
	DOUBLE_STR = "DOUBLE"
	BOOL_STR   = "BOOL"
)

func (v ValueType) String() string {
	switch v {
	case STRING:
		return STRING_STR
	case INT:
		return INT_STR
	case DOUBLE:
		return DOUBLE_STR
	case BOOL:
		return BOOL_STR
	default:
		return ""
	}
}

type Value struct {
	Type       ValueType
	ParseValue interface{}
}
