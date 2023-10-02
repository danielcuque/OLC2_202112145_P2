package compiler

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"
)

func (c *Compiler) ConcatString(leftOp, rightOp *ValueResponse) *ValueResponse {
	procedure := c.TAC.GetStandar("stdconcat")

	if procedure == nil {
		newProcedure := NewProcedure("stdconcat")

		newProcedure.AddArguments(
			[]*Parameter{
				{
					Name:     "HeapPointer",
					Temporal: c.TAC.NewTemporal("p", IntTemporal),
				},
				{
					Name:     "leftOp",
					Temporal: c.TAC.NewTemporal("q", IntTemporal),
				},
				{
					Name:     "rightOp",
					Temporal: c.TAC.NewTemporal("r", IntTemporal),
				},
				{
					Name:     "AccessChar",
					Temporal: c.TAC.NewTemporal("s", IntTemporal),
				},
			},
		)

		newProcedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("FirstString"),
				c.TAC.NewLabel("EndFirstString"),
				c.TAC.NewLabel("SecondString"),
				c.TAC.NewLabel("EndSecondString"),
			},
		)

		// Add code

		newProcedure.AddCode(
			[]string{
				// Get first free space in heap
				fmt.Sprintf("%s = H;", newProcedure.GetArgument("HeapPointer").Tmp()),

				// Declare first label to create a loop for first string
				newProcedure.GetLabel("FirstString").Declare(),

				// Get char from first string
				fmt.Sprintf(
					"%v = heap[%v];",
					newProcedure.GetArgument("AccessChar").Tmp(),
					newProcedure.GetArgument("leftOp").Temporal.Cast(),
				),

				// If char is -1, then go to end of first string
				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					newProcedure.GetArgument("AccessChar").Tmp(),
					newProcedure.GetLabel("EndFirstString"),
				),

				// Add char to heap
				fmt.Sprintf(
					"heap[(int)H] = %v;",
					newProcedure.GetArgument("AccessChar").Tmp(),
				),

				// Increase heap pointer
				"H = H + 1;",

				// Increase pointer to access char
				fmt.Sprintf(
					"%v = %v + 1;",
					newProcedure.GetArgument("AccessChar").Tmp(),
					newProcedure.GetArgument("AccessChar").Tmp(),
				),

				// Go to first label
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("FirstString").String()),

				// Declare end of first string
				newProcedure.GetLabel("EndFirstString").Declare(),
				newProcedure.GetLabel("SecondString").Declare(),

				// Get char from second string
				fmt.Sprintf(
					"%v = heap[%v];",
					newProcedure.GetArgument("AccessChar").Tmp(),
					newProcedure.GetArgument("rightOp").Temporal.Cast(),
				),

				// If char is -1, then go to end of second string
				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					newProcedure.GetArgument("AccessChar").Tmp(),
					newProcedure.GetLabel("EndSecondString"),
				),

				// Add char to heap
				fmt.Sprintf(
					"heap[(int)H] = %v;",
					newProcedure.GetArgument("AccessChar").Tmp(),
				),

				// Increase heap pointer
				"H = H + 1;",
				fmt.Sprintf(
					"%v = %v + 1;",
					newProcedure.GetArgument("rightOp").Tmp(),
					newProcedure.GetArgument("rightOp").Tmp(),
				),

				// Go to second label
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("SecondString").String()),

				// Declare end of second string
				newProcedure.GetLabel("EndSecondString").Declare(),

				// Add -1 to end of string
				"heap[(int)H] = -1;",
				"H = H + 1;",
			},
			"",
		)

		c.TAC.AddProcedure(newProcedure)
	}

	procedure = c.TAC.GetStandar("stdconcat")

	// Set left Operator in temporal
	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetArgument("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetArgument("rightOp").Tmp(), rightOp.GetValue()),
			"stdconcat();",
		},
		"Concatenación de cadenas",
	)

	return &ValueResponse{
		Type:        V.StringType,
		Value:       procedure.GetArgument("HeapPointer").Temporal,
		ContextType: TemporalType,
	}
}

func Print(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	// All args will be Value responses, we will traverse them to get the value, if the value is a string, we use strpint, then we use just printf("%type", value)

	for _, arg := range c.GetArgs(ctx) {
		if arg.Value.Type == V.StringType {
			// Standar print will be for strings
			procedure := c.TAC.GetStandar("stdprint")

			if procedure == nil {
				newProcedure := NewProcedure("stdprint")

				newProcedure.AddArguments(
					[]*Parameter{
						{
							Name:     "HeapPointer",
							Temporal: c.TAC.NewTemporal("p", IntTemporal),
						},
						{
							Name:     "AccessChar",
							Temporal: c.TAC.NewTemporal("q", IntTemporal),
						},
					},
				)

				newProcedure.AddLabels(
					[]*Label{
						c.TAC.NewLabel("StartString"),
						c.TAC.NewLabel("EndString"),
					},
				)

				// Add code

				newProcedure.AddCode(
					[]string{
						// Declare start of string
						newProcedure.GetLabel("StartString").Declare(),

						// Get char from string
						fmt.Sprintf(
							"%v = heap[%v];",
							newProcedure.GetArgument("AccessChar").Tmp(),
							newProcedure.GetArgument("HeapPointer").Temporal.Cast(),
						),

						// If char is -1, then go to end of string
						fmt.Sprintf(
							"if (%v == -1) goto %s;",
							newProcedure.GetArgument("AccessChar").Tmp(),
							newProcedure.GetLabel("EndString"),
						),

						// Print char
						fmt.Sprintf(
							"printf(\"%%c\",(int) %v);",
							newProcedure.GetArgument("AccessChar").Tmp(),
						),

						// Increase pointer to access char
						fmt.Sprintf(
							"%v = %v + 1;",
							newProcedure.GetArgument("HeapPointer").Tmp(),
							newProcedure.GetArgument("HeapPointer").Tmp(),
						),

						// Go to first label
						fmt.Sprintf("goto %s;", newProcedure.GetLabel("StartString").String()),

						// Declare end of string
						newProcedure.GetLabel("EndString").Declare(),
					},
					"",
				)

				c.TAC.AddProcedure(newProcedure)
			}

			procedure = c.TAC.GetStandar("stdprint")

			c.TAC.AppendCode(
				[]string{
					fmt.Sprintf("%v = %v;", procedure.GetArgument("HeapPointer").Tmp(), arg.Value.GetValue()),
					"stdprint();",
				},
				"Imprimiendo cadena",
			)
		} else {
			c.TAC.AppendCode(
				[]string{
					arg.Value.ToPrint(),
				},
				"Imprimiendo valor",
			)
		}

		// After print an arg, we need to print a space
		c.TAC.AppendCode(
			[]string{
				"printf(\"%c\", 32);",
			},
			"",
		)
	}

	c.TAC.AppendCode(
		[]string{
			"printf(\"%c\", 10);",
		},
		"",
	)

	return nil
}

func Int(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	return nil
}
func Float(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	return nil
}
func String(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	return nil
}
func Append(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	return nil
}
func Remove(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	return nil
}
func RemoveLast(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	return nil
}
