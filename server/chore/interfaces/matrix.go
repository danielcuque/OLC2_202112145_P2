package interfaces

import (
	V "OLC2/chore/values"
	"fmt"
)

type MatrixV struct {
	DataType   string
	Body       []V.IValue
	Dimentions int // 2 to n
}

func NewMatrix(Type string, body []V.IValue) *MatrixV {
	return &MatrixV{
		DataType: Type,
		Body:     body,
	}
}

func (m *MatrixV) GetValue() interface{} {
	return m.Body
}

func (m *MatrixV) String() string {
	str := ""
	// Add comma except for last element
	for i, value := range m.Body {
		str += fmt.Sprintf("%v", value)
		if i != len(m.Body)-1 {
			str += ", "
		}
	}
	return str
}

func (m *MatrixV) GetType() string {
	return m.DataType
}
