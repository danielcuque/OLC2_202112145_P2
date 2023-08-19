package interfaces

import (
	V "OLC2/chore/values"
	"fmt"
)

type VectorV struct {
	DataType string
	Body     []V.IValue
}

func NewVector(Type string, body []V.IValue) *VectorV {
	return &VectorV{
		DataType: Type,
		Body:     body,
	}
}

func (v *VectorV) GetValue() interface{} {
	return v.Body
}

func (v *VectorV) String() string {
	str := ""
	// Add comma except for last element
	for i, value := range v.Body {
		str += fmt.Sprintf("%v", value)
		if i != len(v.Body)-1 {
			str += ", "
		}
	}
	return str
}

func (v *VectorV) GetType() string {
	return v.DataType
}

func (v *VectorV) GetDataType() string {
	return v.DataType
}

func (v *VectorV) Append(value V.IValue) {
	v.Body = append(v.Body, value)
}

// Remove last
func (v *VectorV) Pop() V.IValue {
	if v.IsEmpty() {
		return nil
	}

	last := v.Body[v.Count()-1]
	v.Body = v.Body[:v.Count()-1]
	return last
}

// Remove at index
func (v *VectorV) Remove(index int) V.IValue {
	if v.IsEmpty() {
		return nil
	}

	if index < 0 || index >= v.Count() {
		return nil
	}

	removed := v.Body[index]
	v.Body = append(v.Body[:index], v.Body[index+1:]...)
	return removed
}

// Is empty
func (v *VectorV) IsEmpty() bool {
	return v.Count() == 0
}

// Count
func (v *VectorV) Count() int {
	return len(v.Body)
}
