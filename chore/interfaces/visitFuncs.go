package interfaces

import (
	"github.com/antlr4-go/antlr/v4"
)

type VisitFunc func(ctx antlr.ParseTree) Value
