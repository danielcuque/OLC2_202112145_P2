package compiler

type Matrix struct {
	Type      TemporalCast
	Temporal  *Temporal
	Dimension int
	Body      []*Matrix
	Metadata  []int
}

func NewMatrix(dimension int, temporal *Temporal, metadata []int) *Matrix {
	return &Matrix{
		Dimension: dimension,
		Temporal:  temporal,
		Metadata:  metadata,
	}
}

func NewVector(temporal *Temporal, metadata []int) *Object {
	newVector := NewObject("vector", NewMatrix(1, temporal, metadata))
	count := NewSimpleValue(0)
	count.SetData(IntTemporal, metadata[1])

	isEmpty := NewSimpleValue(1)
	isEmpty.SetData(BooleanTemporal, metadata[0])

	newVector.AddProp("count", count)
	newVector.AddProp("isEmpty", isEmpty)

	return newVector
}

func (m *Matrix) GetValue() string {
	return m.Temporal.String()
}

func (m *Matrix) GetInit() int {
	return m.Metadata[0]
}

func (m *Matrix) GetSizeRow() int {
	return m.Metadata[1]
}
