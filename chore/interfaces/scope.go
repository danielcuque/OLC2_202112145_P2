package interfaces

type Scope struct {
	Previous  *Scope
	Name      string
	Variables map[string]IValue
}
