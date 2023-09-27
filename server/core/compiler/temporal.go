package compiler

import "fmt"

type TemporalCast string

const (
	FloatTemporal TemporalCast = "float"
	IntTemporal   TemporalCast = "int"
	CharTemporal  TemporalCast = "char"
)

type Temporal struct {
	ID   int
	Type TemporalCast
}

func NewTemporal(index int, Type TemporalCast) *Temporal {
	return &Temporal{
		ID:   index,
		Type: Type,
	}
}

func (t *Temporal) String() string {
	return fmt.Sprintf("(%s)t%d", t.Type, t.ID)
}
