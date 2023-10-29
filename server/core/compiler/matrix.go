package compiler

type Matrix struct {
	Type TemporalCast
	Body []interface{}
}

func NewMatrix(Type TemporalCast) *Matrix {
	return &Matrix{
		Type: Type,
	}
}

// func (m *Matrix) Print() {
// 	for _, value := range m.Body {
// 		fmt.Println(value)
// 	}
// }

// func RecursivePrint(body []interface{}) {
// 	for _, value := range body {
// 		switch value.(type) {
// 		case *Matrix:
// 			RecursivePrint(value.(*Matrix).Body)
// 		default:
// 			fmt.Println(value)
// 		}
// 	}
// }

func NewVector(Type TemporalCast) *Object {
	newVector := NewObject("vector", NewMatrix(Type))
	count := NewSimpleValue(0)
	count.SetData(IntTemporal, "")

	isEmpty := NewSimpleValue(1)
	isEmpty.SetData(BooleanTemporal, "")

	newVector.AddProp("count", count)
	newVector.AddProp("isEmpty", isEmpty)

	return newVector
}

func (m *Matrix) AddValue(value interface{}) {
	m.Body = append(m.Body, value)
}

func (m *Matrix) GetAddress() string {
	return ""
}

func (m *Matrix) GetValue() interface{} {
	return m.Body
}

func (m *Matrix) GetType() TemporalCast {
	return m.Type
}

func (m *Matrix) SetType(t TemporalCast) {
	m.Type = t
}
