package interfaces

import (
	V "OLC2/core/values"
	"fmt"
)

type MatrixNode struct {
	DataType string
	Body     []V.IValue
}

func NewMatrixNode(Type string, body []V.IValue) *MatrixNode {
	return &MatrixNode{
		DataType: Type,
		Body:     body,
	}
}

func (m *MatrixNode) GetValue() interface{} {
	return m.Body
}

func (m *MatrixNode) String() string {
	return fmt.Sprintf("%v", m.Body)
}

func (m *MatrixNode) GetType() string {
	return m.DataType
}

func (m *MatrixNode) Copy() V.IValue {
	body := make([]V.IValue, len(m.Body))

	for i, value := range m.Body {
		body[i] = value.Copy()
	}

	return NewMatrixNode(m.DataType, body)
}

type MatrixV struct {
	DataType   string
	Body       *MatrixNode
	Dimensions int // 2 to n
}

func NewMatrix(Type string, body *MatrixNode) *MatrixV {
	return &MatrixV{
		DataType: Type,
		Body:     body,
	}
}

func (m *MatrixV) GetValue() interface{} {
	return m.Body
}

func (m *MatrixV) String() string {
	return fmt.Sprintf("%v", m.Body)
}

func (m *MatrixV) GetType() string {
	return m.DataType
}

func (m *MatrixV) Copy() V.IValue {
	body := make([]V.IValue, len(m.Body.Body))

	for i, value := range m.Body.Body {
		body[i] = value.Copy()
	}

	return NewMatrix(m.DataType, NewMatrixNode(m.DataType, body))

}
