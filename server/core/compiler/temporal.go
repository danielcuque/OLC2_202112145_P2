package compiler

type TemporalCast string

const (
	FloatTemporal TemporalCast = "float"
	IntTemporal   TemporalCast = "int"
	CharTemporal  TemporalCast = "char"
)

type Temporal struct {
	Index int
	Type  TemporalCast
}

func NewTemporal() *Temporal {
	return &Temporal{}
}

func (t *Temporal) String() string {
	return ""
}
