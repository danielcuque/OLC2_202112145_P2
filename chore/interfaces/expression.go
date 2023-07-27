package interfaces

import (
	"reflect"
	"strconv"
	"strings"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

func (v *Visitor) Visit(tree antlr.ParseTree) Value {

	visitFuncs := map[reflect.Type]func(ctx antlr.ParseTree) Value{
		reflect.TypeOf((*parser.ProgramContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitProg(ctx.(*parser.ProgramContext))
		},
		reflect.TypeOf((*parser.BlockContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitBlock(ctx.(*parser.BlockContext))
		},
		reflect.TypeOf((*parser.StatementContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitStmt(ctx.(*parser.StatementContext))
		},
		reflect.TypeOf((*parser.AssignmentContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitAssignstmt(ctx.(*parser.AssignmentContext))
		},
		reflect.TypeOf((*parser.IfstmtContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIfstmt(ctx.(*parser.IfstmtContext))
		},
		reflect.TypeOf((*parser.WhilestmtContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitWhilestmt(ctx.(*parser.WhilestmtContext))
		},
		reflect.TypeOf((*parser.OpExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitOpExpr(ctx.(*parser.OpExprContext))
		},
		reflect.TypeOf((*parser.IntExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIntExpr(ctx.(*parser.IntExprContext))
		},
		reflect.TypeOf((*parser.IdExprContext)(nil)): func(ctx antlr.ParseTree) Value {
			return v.VisitIdExpr(ctx.(*parser.IdExprContext))
		},
	}

	ctxType := reflect.TypeOf(tree)

	if fn, ok := visitFuncs[ctxType]; ok {
		return fn(tree)
	}

	// If the type is not found in the map we panic
	panic("Unknown context " + ctxType.String())
}

func (v *Visitor) VisitProg(ctx *parser.ProgramContext) Value {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value {
	for i := 0; ctx.Statement(i) != nil; i++ {
		v.Visit(ctx.Statement(i))
	}
	return Value{ParseValue: true}
}

func (v *Visitor) VisitStmt(ctx *parser.StatementContext) Value {
	if ctx.Assignment() != nil {
		return v.Visit(ctx.Assignment())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Whilestmt() != nil {
		return v.Visit(ctx.Whilestmt())
	}
	return Value{ParseValue: true}
}

func (v *Visitor) VisitAssignstmt(ctx *parser.AssignmentContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.Memory[id] = value
	return Value{ParseValue: true}
}

func (v *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).ParseValue.(bool)
	if ok && value {
		return v.Visit(ctx.Block())
	}
	return Value{ParseValue: false}
}

func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).ParseValue.(bool)
	for ok && value {
		v.Visit(ctx.Block())
		value, ok = v.Visit(ctx.Expr()).ParseValue.(bool)
	}
	return Value{ParseValue: true}
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{ParseValue: i}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	value := strings.Trim(ctx.GetText(), "\"")
	return Value{ParseValue: value}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()
	value, ok := v.Memory[id]
	if ok {
		return value
	} else {
		panic("no such variable: " + id)
	}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{ParseValue: value}
}

func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext) Value {
	l := v.Visit(ctx.GetLeft()).ParseValue.(int64)
	r := v.Visit(ctx.GetRight()).ParseValue.(int64)
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		return Value{ParseValue: l + r}
	case "-":
		return Value{ParseValue: l - r}
	case "*":
		return Value{ParseValue: l * r}
	case "/":
		return Value{ParseValue: l / r}
	case "%":
		return Value{ParseValue: l % r}
	case "<":
		if l < r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case ">":
		if l > r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "<=":
		if l <= r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case ">=":
		if l >= r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "==":
		if l == r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "!=":
		if l != r {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "&&":
		if l != 0 && r != 0 {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "||":
		if l != 0 || r != 0 {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	case "!":
		if r == 0 {
			return Value{ParseValue: true}
		} else {
			return Value{ParseValue: false}
		}
	}
	return Value{ParseValue: false}
}
