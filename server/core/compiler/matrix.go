package compiler

type Matrix struct {
	Type      TemporalCast
	Temporal  *Temporal
	Dimension int
	Body      []*Matrix
}

func NewMatrix(dimension int, temporal *Temporal) *Matrix {
	return &Matrix{
		Dimension: dimension,
		Temporal:  temporal,
	}
}

func NewVector(temporal *Temporal) *Object {
	newVector := NewObject("vector", NewMatrix(1, temporal))
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
