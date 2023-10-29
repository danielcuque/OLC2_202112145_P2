package compiler

import (
	"OLC2/core/parser"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func (c *Compiler) arithmeticOp(l, r interface{}, op string, lc antlr.Token) interface{} {
	leftT := l.(*ValueResponse).GetType()
	rightT := r.(*ValueResponse).GetType()

	lV := l.(*ValueResponse).GetValue()
	rV := r.(*ValueResponse).GetValue()

	if op == "/" || op == "%" || op == "*" {
		lV = l.(*ValueResponse).Cast()
		rV = r.(*ValueResponse).Cast()
	}

	var response *ValueResponse

	if (leftT == StringTemporal || leftT == CharTemporal) || (rightT == StringTemporal || rightT == CharTemporal) {
		return c.ConcatString(l.(*ValueResponse), r.(*ValueResponse))
	}

	if (op == "/" || op == "%") && (rightT == IntTemporal) {
		return c.ZeroDivision(l.(*ValueResponse), r.(*ValueResponse), op)
	}

	if leftT == IntTemporal && rightT == IntTemporal {
		response = &ValueResponse{
			Type:        IntTemporal,
			Value:       c.TAC.NewTemporal(IntTemporal), // Temporal
			ContextType: TemporalType,
		}
	}

	if leftT == FloatTemporal || rightT == FloatTemporal {
		response = &ValueResponse{
			Type:        FloatTemporal,
			Value:       c.TAC.NewTemporal(FloatTemporal), // Temporal
			ContextType: TemporalType,
		}
	}

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = %s %s %s;", response.GetValue(), lV, op, rV),
		},
		fmt.Sprintf("Operación aritmética %s", op))
	return response
}

func (c *Compiler) VisitArithmeticExpr(ctx *parser.ArithmeticExprContext) interface{} {
	var l, r *ValueResponse

	op := ctx.GetOp().GetText()

	_, isCall := ctx.GetRight().(*parser.FunctionCallExprContext)

	if isCall {
		r = c.Visit(ctx.GetRight()).(*ValueResponse)
		l = c.Visit(ctx.GetLeft()).(*ValueResponse)
	} else {
		l = c.Visit(ctx.GetLeft()).(*ValueResponse)
		r = c.Visit(ctx.GetRight()).(*ValueResponse)
	}

	return c.arithmeticOp(l, r, op, ctx.GetStart())
}

func (c *Compiler) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	value := "1"

	if ctx.GetText() == "false" {
		value = "0"
	}

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       value,
		ContextType: LiteralType,
	}
}

func (c *Compiler) VisitCharExpr(ctx *parser.CharExprContext) interface{} {

	newTemporal := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = H;", newTemporal),
			fmt.Sprintf("heap[(int)H] = %d;", ctx.GetText()[1]),
			"H = H + 1;",
		},
		fmt.Sprintf("Almacenamiento de caracter '%d' en el heap", ctx.GetText()[1]),
	)

	c.HeapPointer.AddPointer()

	return &ValueResponse{
		Type:        CharTemporal,
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitComparisonExpr(ctx *parser.ComparisonExprContext) interface{} {
	left := c.Visit(ctx.GetLeft()).(*ValueResponse)
	right := c.Visit(ctx.GetRight()).(*ValueResponse)

	op := ctx.GetOp().GetText()

	if left.GetType() == StringTemporal && right.GetType() == StringTemporal {
		return c.CompareString(left, right, op)
	}

	trueLabel := c.TAC.NewLabel("")
	falseLabel := c.TAC.NewLabel("")
	temporalResult := c.TAC.NewTemporal(BooleanTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("if (%s %s %s) goto %s;", left.GetValue(), op, right.GetValue(), trueLabel.String()),
			fmt.Sprintf("%s = 0;", temporalResult),
			fmt.Sprintf("goto %s;", falseLabel.String()),
			trueLabel.Declare(),
			fmt.Sprintf("%s = 1;", temporalResult),
			falseLabel.Declare(),
		},
		fmt.Sprintf("Operación %s %s %s", left.GetValue(), op, right.GetValue()),
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       temporalResult,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return &ValueResponse{
		Type:        FloatTemporal,
		Value:       ctx.GetText(),
		ContextType: LiteralType,
	}
}

func (c *Compiler) VisitFunctionCallExpr(ctx *parser.FunctionCallExprContext) interface{} {
	return c.Visit(ctx.FunctionCall())
}

func (c *Compiler) VisitIDChain(ctx *parser.IDChainContext) interface{} {
	return ctx.AllID()
}

func (c *Compiler) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id, props := c.GetPropsAsString(ctx.IdChain().(*parser.IDChainContext))

	value := c.Env.GetValue(id)

	if value == nil {
		return nil
	}

	newTemporal := c.TAC.NewTemporal(value.GetType())

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = stack[(int)%s];", newTemporal, c.TAC.GetValueAddress(value)),
		},
		fmt.Sprintf("Acceso a la variable '%s'", id),
	)

	temporalPointer := c.GetProps(value, props[1:], newTemporal)

	return &ValueResponse{
		Type:        temporalPointer.GetType(),
		Value:       temporalPointer,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return &ValueResponse{
		Type:        IntTemporal,
		Value:       ctx.GetText(),
		ContextType: LiteralType,
	}
}

func (c *Compiler) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	left := c.Visit(ctx.GetLeft()).(*ValueResponse)
	right := c.Visit(ctx.GetRight()).(*ValueResponse)

	op := ctx.GetOp().GetText()

	if op == "&&" {
		return c.And(left, right)
	}

	return c.Or(left, right)
}

func (c *Compiler) VisitMatrixAccessExpr(ctx *parser.MatrixAccessExprContext) interface{} {
	return c.Visit(ctx.MatrixAccess())
}

func (c *Compiler) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	return &ValueResponse{
		Type:        FloatTemporal,
		Value:       DefaultNil(),
		ContextType: LiteralType,
	}
}

func (c *Compiler) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	value := c.Visit(ctx.Expr()).(*ValueResponse)

	if c.TAC.GetProcedure("std_not") == nil {
		newProcedure := NewProcedure("std_not")

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "operand",
					Temporal:     c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					ExternalName: "result",
					Temporal:     c.TAC.NewTemporal(BooleanTemporal),
				},
			},
		)

		newProcedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("FalseCondition"),
				c.TAC.NewLabel("TrueCondition"),
			},
		)

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf(
					"if (%v == 0) goto %s;",
					newProcedure.GetParameter("operand").Temporal,
					newProcedure.GetLabel("TrueCondition"),
				),
				fmt.Sprintf(
					"%s = 0;",
					newProcedure.GetParameter("result").Temporal,
				),
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("FalseCondition")),
				newProcedure.GetLabel("TrueCondition").Declare(),
				fmt.Sprintf(
					"%s = 1;",
					newProcedure.GetParameter("result").Temporal,
				),
				newProcedure.GetLabel("FalseCondition").Declare(),
			},
			"Operación not",
		)

		c.TAC.AddProcedure(newProcedure)
	}

	procedure := c.TAC.GetProcedure("std_not")

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf(
				"%s = %s;",
				procedure.GetParameter("operand").Temporal,
				value.GetValue(),
			),
			"std_not();",
		},
		"Operación not",
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       procedure.GetParameter("result").Temporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitParExpr(ctx *parser.ParExprContext) interface{} {
	return c.Visit(ctx.Expr())
}

func (c *Compiler) VisitRangeExpr(ctx *parser.RangeExprContext) interface{} {

	left := c.Visit(ctx.GetLeft()).(*ValueResponse)
	right := c.Visit(ctx.GetRight()).(*ValueResponse)

	validRange := c.TAC.NewLabel("")

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("if (%s > %s) goto %s;", right.GetValue(), left.GetValue(), validRange.String()),
			"printf(\"%c\", 79);",  //O
			"printf(\"%c\", 117);", //u
			"printf(\"%c\", 116);", //t
			"printf(\"%c\", 32);",  //
			"printf(\"%c\", 111);", //o
			"printf(\"%c\", 102);", //f
			"printf(\"%c\", 32);",  //
			"printf(\"%c\", 66);",  //B
			"printf(\"%c\", 111);", //o
			"printf(\"%c\", 117);", //u
			"printf(\"%c\", 110);", //n
			"printf(\"%c\", 100);", //d
			"printf(\"%c\", 115);", //s
			"printf(\"%c\", 10);",  //
			validRange.Declare(),
		},
		"Validación de rango",
	)

	return map[string]interface{}{
		"left":  left.GetValue(),
		"right": right.GetValue(),
	}
}

func (c *Compiler) VisitStrExpr(ctx *parser.StrExprContext) interface{} {
	s := strings.Trim(ctx.GetText(), "\"")

	var stringType TemporalCast

	if len(s) <= 1 {
		stringType = CharTemporal
	} else {
		stringType = StringTemporal
	}

	// if len(s) == 0 {
	// 	return &ValueResponse{
	// 		Type:        CharTemporal,
	// 		Value:       fmt.Sprintf("%d", 0),
	// 		ContextType: LiteralType,
	// 	}
	// }

	// if len(s) == 1 {
	// 	return &ValueResponse{
	// 		Type:        CharTemporal,
	// 		Value:       fmt.Sprintf("%d", s[0]),
	// 		ContextType: LiteralType,
	// 	}
	// }

	// Replace scape characters: double quote, backslash, new line, carriage return, tab
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\r", "\r")
	s = strings.ReplaceAll(s, "\\t", "\t")

	// Traverse string to insert chars in heap
	newTemporal := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = H;", newTemporal),
		},
		"Obtención de la posición inicial del heap",
	)

	for _, char := range s {
		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("heap[(int)H] = %d;", char),
				"H = H + 1;",
			},
			"",
		)
		c.HeapPointer.AddPointer()
	}

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("heap[(int)H] = %d;", -1),
			"H = H + 1;",
		},
		"Almacenamiento del caracter nulo en el heap",
	)
	c.HeapPointer.AddPointer()

	return &ValueResponse{
		Type:        stringType,
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	response := c.Visit(ctx.Expr()).(*ValueResponse)

	newTemporal := c.TAC.NewTemporal(response.GetType())

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = %s * -1;", newTemporal, response.GetValue()),
		},
		fmt.Sprintf("Operación aritmética %s", "-"),
	)

	return &ValueResponse{
		Type:        response.GetType(),
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) VisitVariableType(ctx *parser.VariableTypeContext) interface{} {
	switch ctx.GetText() {
	case "Int":
		return IntTemporal
	case "Float":
		return FloatTemporal
	case "Character":
		return CharTemporal
	case "Boolean":
		return BooleanTemporal
	case "String":
		return StringTemporal
	default:
		return ctx.GetText()
	}
}

func (c *Compiler) VisitVectorAccessExpr(ctx *parser.VectorAccessExprContext) interface{} {
	evalValue := c.Visit(ctx.VectorAccess())

	if evalValue == nil {
		return nil
	}

	return c.GetVectorValue(evalValue.(*ValueResponse))
}

func (c *Compiler) GetVectorValue(response *ValueResponse) *ValueResponse {
	vectorPosition := response.GetValue()
	returnTemporal := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = heap[(int)%s];", returnTemporal, vectorPosition),
		},
		"Acceso a vector",
	)

	return &ValueResponse{
		Type:        response.GetType(),
		Value:       returnTemporal,
		ContextType: TemporalType,
	}
}

// Utils

func (c *Compiler) GetPropsAsString(ctx *parser.IDChainContext) (string, []string) {
	props := strings.Split(ctx.GetText(), ".")
	id := props[0]

	return id, props
}

func (c *Compiler) GetProps(value *Value, props []string, auxiliarTemporal *Temporal) *Temporal {
	if len(props) == 0 {
		return auxiliarTemporal
	}

	// If length of props is 1, then we are accessing to a property of object

	obj, notObject := value.GetValue().(*Object)

	if !notObject {
		return auxiliarTemporal
	}

	// Check if object value is nested object or not

	objectValue, ok := obj.GetValue().(*Object)

	if !ok {
		val := obj.GetProp(props[0])

		if val == nil {
			return auxiliarTemporal
		}

		relativePosition := c.TAC.NewTemporal(IntTemporal)
		heapTemporal := c.TAC.NewTemporal(val.GetType())

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf("%s = %s + %s;", relativePosition, auxiliarTemporal, val.GetAddress()),
				fmt.Sprintf("%s = heap[%s];", heapTemporal, relativePosition.Cast()),
			},
			fmt.Sprintf("Propiedad '%s'", props[0]),
		)

		return c.GetProps(val, props[1:], heapTemporal)
	}

	value = objectValue.GetProp(props[0])

	stackTemporal := c.TAC.NewTemporal(value.GetType())

	if value == nil {
		return auxiliarTemporal
	}

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = stack[(int)%s];", stackTemporal, c.TAC.GetValueAddress(value)),
		},
		fmt.Sprintf("Propiedad '%s'", props[0]),
	)

	return c.GetProps(value, props[1:], stackTemporal)
}
