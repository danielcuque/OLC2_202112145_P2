package compiler

import "fmt"

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

func NewTemporal(index int, Type TemporalCast) *Temporal {
	return &Temporal{
		Index: index,
		Type:  Type,
	}
}

func (t *Temporal) String() string {
	return fmt.Sprintf("(%s)t%d", t.Type, t.Index)
}
