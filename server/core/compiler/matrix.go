package compiler

type Matrix struct {
	Type      TemporalCast
	Dimension int
	Body      []*Matrix
}

func NewMatrix(dimension int, Type TemporalCast) *Matrix {
	return &Matrix{
		Dimension: dimension,
		Type:      Type,
	}
}

func NewVector(Type TemporalCast) *Object {
	newVector := NewObject("vector", NewMatrix(1, Type))
	count := NewSimpleValue(0)
	count.SetData(IntTemporal, "")

	isEmpty := NewSimpleValue(1)
	isEmpty.SetData(BooleanTemporal, "")

	newVector.AddProp("count", count)
	newVector.AddProp("isEmpty", isEmpty)

	return newVector
}

func (m *Matrix) GetType() TemporalCast {
	return m.Type
}

func (m *Matrix) SetType(t TemporalCast) {
	m.Type = t
}
