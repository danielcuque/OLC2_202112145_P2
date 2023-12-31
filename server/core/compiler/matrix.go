package compiler

type Matrix struct {
	Type TemporalCast
	Body []interface{}
}

type InitMatrix struct {
	InitVector     *Temporal
	Counter        *Temporal
	IsEmptyAddress *Temporal
}

func (c *Compiler) NewInitMatrix() *InitMatrix {
	return &InitMatrix{
		InitVector:     c.TAC.NewTemporal(IntTemporal),
		Counter:        c.TAC.NewTemporal(IntTemporal),
		IsEmptyAddress: c.TAC.NewTemporal(BooleanTemporal),
	}
}

func NewMatrix(Type TemporalCast) *Matrix {
	return &Matrix{
		Type: Type,
	}
}

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
