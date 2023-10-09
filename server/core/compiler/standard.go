package compiler

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"
)

func (c *Compiler) And(leftOp, rightOp *ValueResponse) *ValueResponse {
	if c.TAC.GetStandar("std_and") == nil {
		newProcedure := NewProcedure("std_and")

		newProcedure.AddArguments(
			[]*Parameter{
				{
					Name:     "leftOp",
					Temporal: c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					Name:     "rightOp",
					Temporal: c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					Name:     "result",
					Temporal: c.TAC.NewTemporal(BooleanTemporal),
				},
			},
		)

		newProcedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("FalseCondition"),
				c.TAC.NewLabel("TrueCondition"),
			},
		)

		// Add code

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf(
					"if (%v == 0) goto %v;",
					newProcedure.GetArgument("leftOp").Temporal,
					newProcedure.GetLabel("FalseCondition"),
				),
				fmt.Sprintf(
					"if (%v == 0) goto %v;",
					newProcedure.GetArgument("rightOp").Temporal,
					newProcedure.GetLabel("FalseCondition"),
				),
				fmt.Sprintf(
					"%v = 1;",
					newProcedure.GetArgument("result").Temporal,
				),
				fmt.Sprintf("goto %v;", newProcedure.GetLabel("TrueCondition")),
				newProcedure.GetLabel("FalseCondition").Declare(),
				fmt.Sprintf(
					"%v = 0;",
					newProcedure.GetArgument("result").Temporal,
				),
				newProcedure.GetLabel("TrueCondition").Declare(),
			},
			"Operación AND",
		)

		c.TAC.AddProcedure(newProcedure)
	}

	procedure := c.TAC.GetStandar("std_and")

	// Set left Operator in temporal

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetArgument("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetArgument("rightOp").Tmp(), rightOp.GetValue()),
			"std_and();",
		},
		"Operación AND",
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       procedure.GetArgument("result").Temporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) Or(leftOp, rightOp *ValueResponse) *ValueResponse {
	if c.TAC.GetStandar("std_or") == nil {
		newProcedure := NewProcedure("std_or")

		newProcedure.AddArguments(
			[]*Parameter{
				{
					Name:     "leftOp",
					Temporal: c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					Name:     "rightOp",
					Temporal: c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					Name:     "result",
					Temporal: c.TAC.NewTemporal(BooleanTemporal),
				},
			},
		)

		newProcedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("TrueCondition"),
				c.TAC.NewLabel("FalseCondition"),
			},
		)

		// Add code

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf(
					"if (%v == 1) goto %v;",
					newProcedure.GetArgument("leftOp").Temporal,
					newProcedure.GetLabel("TrueCondition"),
				),
				fmt.Sprintf(
					"if (%v == 1) goto %v;",
					newProcedure.GetArgument("rightOp").Temporal,
					newProcedure.GetLabel("TrueCondition"),
				),
				fmt.Sprintf(
					"%v = 0;",
					newProcedure.GetArgument("result").Temporal,
				),
				fmt.Sprintf("goto %v;", newProcedure.GetLabel("FalseCondition")),
				newProcedure.GetLabel("TrueCondition").Declare(),
				fmt.Sprintf(
					"%v = 1;",
					newProcedure.GetArgument("result").Temporal,
				),
				newProcedure.GetLabel("FalseCondition").Declare(),
			},
			"Operación OR",
		)

		c.TAC.AddProcedure(newProcedure)
	}

	procedure := c.TAC.GetStandar("std_or")

	// Set left Operator in temporal

	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetArgument("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetArgument("rightOp").Tmp(), rightOp.GetValue()),
			"std_or();",
		},
		"Operación OR",
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       procedure.GetArgument("result").Temporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) ConcatString(leftOp, rightOp *ValueResponse) *ValueResponse {

	if c.TAC.GetStandar("stdconcat") == nil {
		newProcedure := NewProcedure("stdconcat")

		newProcedure.AddArguments(
			[]*Parameter{
				{
					Name:     "HeapPointer",
					Temporal: c.TAC.NewTemporal(IntTemporal),
				},
				{
					Name:     "leftOp",
					Temporal: c.TAC.NewTemporal(IntTemporal),
				},
				{
					Name:     "rightOp",
					Temporal: c.TAC.NewTemporal(IntTemporal),
				},
				{
					Name:     "AccessChar",
					Temporal: c.TAC.NewTemporal(IntTemporal),
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

	procedure := c.TAC.GetStandar("stdconcat")

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

func (c *Compiler) PrintBool(name string) {
	newProcedure := NewProcedure(name)

	newProcedure.AddArguments(
		[]*Parameter{
			{
				Name:     "HeapPointer",
				Temporal: c.TAC.NewTemporal(IntTemporal),
			},
		},
	)

	newProcedure.AddLabels(
		[]*Label{
			c.TAC.NewLabel("True"),
			c.TAC.NewLabel("False"),
			c.TAC.NewLabel("EndBool"),
		},
	)

	newProcedure.AddCode(
		[]string{
			fmt.Sprintf(
				"if (%v == 1) goto %s;",
				newProcedure.GetArgument("HeapPointer").Temporal.Cast(),
				newProcedure.GetLabel("True").String(),
			),
			fmt.Sprintf("goto %s;", newProcedure.GetLabel("False").String()),
			newProcedure.GetLabel("True").Declare(),
			"printf(\"%c\", (char) 116);",
			"printf(\"%c\", (char) 114);",
			"printf(\"%c\", (char) 117);",
			"printf(\"%c\", (char) 101);",
			fmt.Sprintf("goto %s;", newProcedure.GetLabel("EndBool")),
			newProcedure.GetLabel("False").Declare(),
			"printf(\"%c\", (char) 102);",
			"printf(\"%c\", (char) 97);",
			"printf(\"%c\", (char) 108);",
			"printf(\"%c\", (char) 115);",
			"printf(\"%c\", (char) 101);",
			newProcedure.GetLabel("EndBool").Declare(),
		},
		"",
	)

	c.TAC.AddProcedure(newProcedure)

}

func (c *Compiler) PrintString(name string) {
	newProcedure := NewProcedure(name)

	newProcedure.AddArguments(
		[]*Parameter{
			{
				Name:     "HeapPointer",
				Temporal: c.TAC.NewTemporal(IntTemporal),
			},
			{
				Name:     "AccessChar",
				Temporal: c.TAC.NewTemporal(IntTemporal),
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

func Print(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	for _, arg := range c.GetArgs(ctx) {
		if arg.Value.Type == StringTemporal {

			if c.TAC.GetStandar("stdprint") == nil {
				c.PrintString("stdprint")
			}

			procedure := c.TAC.GetStandar("stdprint")

			c.TAC.AppendCode(
				[]string{
					fmt.Sprintf("%v = %v;", procedure.GetArgument("HeapPointer").Tmp(), arg.Value.GetValue()),
					"stdprint();",
				},
				"Imprimiendo cadena",
			)

		} else if arg.Value.Type == BooleanTemporal {

			if c.TAC.GetStandar("stdprintbool") == nil {
				c.PrintBool("stdprintbool")
			}

			procedure := c.TAC.GetStandar("stdprintbool")

			c.TAC.AppendCode(
				[]string{
					fmt.Sprintf("%v = %v;", procedure.GetArgument("HeapPointer").Tmp(), arg.Value.GetValue()),
					"stdprintbool();",
				},
				"Imprimiendo booleano",
			)

		} else {
			c.TAC.AppendCode(
				[]string{
					arg.Value.ToPrint(),
				},
				"Imprimiendo valor",
			)
		}

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
