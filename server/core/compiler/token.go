package compiler

type TokenSymbol struct {
	Env      string
	Type     string
	Name     string
	DataType string
	Value    string
	Params   []string
}

func NewTokenSymbol(env, tokenType, name, dataType, value string, params []string) *TokenSymbol {
	return &TokenSymbol{
		Env:      env,
		Type:     tokenType,
		Name:     name,
		DataType: dataType,
		Value:    value,
		Params:   params,
	}
}
