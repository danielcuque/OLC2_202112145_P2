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

func NewVector(temporal *Temporal, metadata []int) *Matrix {
	return NewMatrix(1, temporal, metadata)
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
