package interfaces

import (
	V "OLC2/chore/values"
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
