package interpreter

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
)

func (v *Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	condition, ok := v.Visit(ctx.Expr()).(V.IValue).GetValue().(bool)

	if !ok {
		v.NewError("Se esperaba una expresión booleana", ctx.GetStart())
		return nil
	}

	v.Env.PushEnv(WhileEnv)

	v.Stack.Push(NewStackItem(
		"While",
		V.NewNilValue(nil),
		[]StackItemType{BreakType, ContinueType},
	))

	v.ExecuteWhile(condition, ctx)

	v.Stack.Pop()
	v.Env.PopEnv()

	return nil
}

func (v *Visitor) ExecuteWhile(condition bool, ctx *parser.WhileStatementContext) {

	defer func() {
		peek, ok := recover().(*StackItem)

		if !ok {
			return
		}

		if peek.Trigger == ReturnType {
			panic(peek)
		}

		if peek.Trigger == BreakType {
			return
		}

		if peek.Trigger == ContinueType {
			condition = v.Visit(ctx.Expr()).(V.IValue).GetValue().(bool)
			v.ExecuteWhile(condition, ctx)

		}
	}()

	for condition {
		v.Env.ResetEnv()
		v.Visit(ctx.Block())
		condition = v.Visit(ctx.Expr()).(V.IValue).GetValue().(bool)
	}

}
