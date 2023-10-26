package compiler

type Matrix struct {
	Type      TemporalCast
	Temporal  *Temporal
	Dimension int
	Body      []*Matrix
}

func NewMatrix(dimension int, temporal *Temporal, Type TemporalCast) *Matrix {
	return &Matrix{
		Dimension: dimension,
		Temporal:  temporal,
		Type:      Type,
	}
}

func NewVector(temporal *Temporal, Type TemporalCast) *Object {
	newVector := NewObject("vector", NewMatrix(1, temporal, Type))
	count := NewSimpleValue(0)
	count.SetData(IntTemporal, "")

	isEmpty := NewSimpleValue(1)
	isEmpty.SetData(BooleanTemporal, "")

	newVector.AddProp("count", count)
	newVector.AddProp("isEmpty", isEmpty)

	return newVector
}

func (m *Matrix) GetValue() string {
	return m.Temporal.String()
}

func (m *Matrix) GetType() TemporalCast {
	return m.Type
}

func (m *Matrix) SetType(t TemporalCast) {
	m.Type = t
}
