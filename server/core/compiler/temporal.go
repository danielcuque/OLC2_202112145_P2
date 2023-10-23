package compiler

import "fmt"

type TemporalCast string

const (
	BooleanTemporal TemporalCast = "b"
	CharTemporal    TemporalCast = "c"
	FloatTemporal   TemporalCast = "f"
	IntTemporal     TemporalCast = "d"
	MatrixTemporal  TemporalCast = "m"
	StringTemporal  TemporalCast = "s"
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

func (t *Temporal) Cast() string {
	switch t.Type {
	case FloatTemporal:
		return fmt.Sprintf("(float) t%d", t.ID)
	case IntTemporal:
		return fmt.Sprintf("(int) t%d", t.ID)
	case CharTemporal:
		return fmt.Sprintf("(char) t%d", t.ID)
	}
	return ""
}

func (t *Temporal) String() string {
	return fmt.Sprintf("t%d", t.ID)
}
