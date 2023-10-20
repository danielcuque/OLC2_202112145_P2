package compiler

import (
	"OLC2/core/parser"
	V "OLC2/core/values"
	"fmt"
)

func (c *Compiler) AllocateStack(size int) {
	allocateStackName := "std_allocate_stack"

	if c.TAC.GetStandard(allocateStackName) == nil {
		newProcedure := NewProcedure(allocateStackName)

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "size",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "basePointer",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
			},
		)

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf(
					"stack[(int)P] = %v;",
					newProcedure.GetParameter("basePointer").Temporal.Cast(),
				),
				fmt.Sprintf(
					"P = P + %v;",
					newProcedure.GetParameter("size").Temporal.Cast(),
				),
			},
			"",
		)

		c.TAC.AddStandard(newProcedure)
	}

	procedure := c.TAC.GetStandard(allocateStackName)
	initFreeAddress := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = P;", initFreeAddress),
			fmt.Sprintf("%v = %d;", procedure.GetParameter("size").Tmp(), size),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("basePointer").Tmp(), c.TAC.GetOffSetPointer()),
			fmt.Sprintf("%v();", allocateStackName),
			fmt.Sprintf("%v = %v;", c.TAC.GetOffSetPointer(), initFreeAddress),
		},
		"Reservando espacio para la instancia de la función",
	)
}

func (c *Compiler) And(leftOp, rightOp *ValueResponse) *ValueResponse {
	if c.TAC.GetStandard("std_and") == nil {
		newProcedure := NewProcedure("std_and")

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "leftOp",
					Temporal:     c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					ExternalName: "rightOp",
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

		// Add code

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf(
					"if (%v == 0) goto %v;",
					newProcedure.GetParameter("leftOp").Temporal,
					newProcedure.GetLabel("FalseCondition"),
				),
				fmt.Sprintf(
					"if (%v == 0) goto %v;",
					newProcedure.GetParameter("rightOp").Temporal,
					newProcedure.GetLabel("FalseCondition"),
				),
				fmt.Sprintf(
					"%v = 1;",
					newProcedure.GetParameter("result").Temporal,
				),
				fmt.Sprintf("goto %v;", newProcedure.GetLabel("TrueCondition")),
				newProcedure.GetLabel("FalseCondition").Declare(),
				fmt.Sprintf(
					"%v = 0;",
					newProcedure.GetParameter("result").Temporal,
				),
				newProcedure.GetLabel("TrueCondition").Declare(),
			},
			"Operación AND",
		)

		c.TAC.AddStandard(newProcedure)
	}

	procedure := c.TAC.GetStandard("std_and")

	// Set left Operator in temporal

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("rightOp").Tmp(), rightOp.GetValue()),
			"std_and();",
		},
		"Operación AND",
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       procedure.GetParameter("result").Temporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) Or(leftOp, rightOp *ValueResponse) *ValueResponse {
	if c.TAC.GetStandard("std_or") == nil {
		newProcedure := NewProcedure("std_or")

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "leftOp",
					Temporal:     c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					ExternalName: "rightOp",
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
				c.TAC.NewLabel("TrueCondition"),
				c.TAC.NewLabel("FalseCondition"),
			},
		)

		// Add code

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf(
					"if (%v == 1) goto %v;",
					newProcedure.GetParameter("leftOp").Temporal,
					newProcedure.GetLabel("TrueCondition"),
				),
				fmt.Sprintf(
					"if (%v == 1) goto %v;",
					newProcedure.GetParameter("rightOp").Temporal,
					newProcedure.GetLabel("TrueCondition"),
				),
				fmt.Sprintf(
					"%v = 0;",
					newProcedure.GetParameter("result").Temporal,
				),
				fmt.Sprintf("goto %v;", newProcedure.GetLabel("FalseCondition")),
				newProcedure.GetLabel("TrueCondition").Declare(),
				fmt.Sprintf(
					"%v = 1;",
					newProcedure.GetParameter("result").Temporal,
				),
				newProcedure.GetLabel("FalseCondition").Declare(),
			},
			"Operación OR",
		)

		c.TAC.AddStandard(newProcedure)
	}

	procedure := c.TAC.GetStandard("std_or")

	// Set left Operator in temporal

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("rightOp").Tmp(), rightOp.GetValue()),
			"std_or();",
		},
		"Operación OR",
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       procedure.GetParameter("result").Temporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) ConcatString(leftOp, rightOp *ValueResponse) *ValueResponse {

	if c.TAC.GetStandard("stdconcat") == nil {
		newProcedure := NewProcedure("stdconcat")

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "HeapPointer",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "leftOp",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "rightOp",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "AccessChar",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
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
				fmt.Sprintf("%s = H;", newProcedure.GetParameter("HeapPointer").Tmp()),

				// Declare first label to create a loop for first string
				newProcedure.GetLabel("FirstString").Declare(),

				// Get char from first string
				fmt.Sprintf(
					"%v = heap[%v];",
					newProcedure.GetParameter("AccessChar").Tmp(),
					newProcedure.GetParameter("leftOp").Temporal.Cast(),
				),

				// If char is -1, then go to end of first string
				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					newProcedure.GetParameter("AccessChar").Tmp(),
					newProcedure.GetLabel("EndFirstString"),
				),

				// Add char to heap
				fmt.Sprintf(
					"heap[(int)H] = %v;",
					newProcedure.GetParameter("AccessChar").Tmp(),
				),

				// Increase heap pointer
				"H = H + 1;",

				// Increase pointer to access char
				fmt.Sprintf(
					"%v = %v + 1;",
					newProcedure.GetParameter("leftOp").Tmp(),
					newProcedure.GetParameter("leftOp").Tmp(),
				),

				// Go to first label
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("FirstString").String()),

				// Declare end of first string
				newProcedure.GetLabel("EndFirstString").Declare(),
				newProcedure.GetLabel("SecondString").Declare(),

				// Get char from second string
				fmt.Sprintf(
					"%v = heap[%v];",
					newProcedure.GetParameter("AccessChar").Tmp(),
					newProcedure.GetParameter("rightOp").Temporal.Cast(),
				),

				// If char is -1, then go to end of second string
				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					newProcedure.GetParameter("AccessChar").Tmp(),
					newProcedure.GetLabel("EndSecondString"),
				),

				// Add char to heap
				fmt.Sprintf(
					"heap[(int)H] = %v;",
					newProcedure.GetParameter("AccessChar").Tmp(),
				),

				// Increase heap pointer
				"H = H + 1;",
				fmt.Sprintf(
					"%v = %v + 1;",
					newProcedure.GetParameter("rightOp").Tmp(),
					newProcedure.GetParameter("rightOp").Tmp(),
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

		c.TAC.AddStandard(newProcedure)
	}

	procedure := c.TAC.GetStandard("stdconcat")

	// Set left Operator in temporal
	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("rightOp").Tmp(), rightOp.GetValue()),
			"stdconcat();",
		},
		"Concatenación de cadenas",
	)

	return &ValueResponse{
		Type:        V.StringType,
		Value:       procedure.GetParameter("HeapPointer").Temporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) CompareString(leftOp, rightOp *ValueResponse, op string) *ValueResponse {
	// There are two cases, when the operator is == or !=

	procedureName := "std_compare_str"
	if c.TAC.GetStandard(procedureName) == nil {
		newProcedure := NewProcedure(procedureName)

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "strPointer1",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "strPointer2",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "AccessChar1",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "AccessChar2",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "result",
					Temporal:     c.TAC.NewTemporal(BooleanTemporal),
				},
				{
					ExternalName: "result2",
					Temporal:     c.TAC.NewTemporal(BooleanTemporal),
				},
			},
		)

		newProcedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("5"),
				c.TAC.NewLabel("6"),
				c.TAC.NewLabel("7"),
				c.TAC.NewLabel("8"),
				c.TAC.NewLabel("9"),
				c.TAC.NewLabel("10"),

				c.TAC.NewLabel("TrueCondition"),
				c.TAC.NewLabel("EndCondition"),
			},
		)

		// Add code
		newProcedure.AddCode(
			[]string{
				newProcedure.GetLabel("5").Declare(),
				fmt.Sprintf(
					"%v = heap[%v];",
					newProcedure.GetParameter("AccessChar1").Tmp(),
					newProcedure.GetParameter("strPointer1").Temporal.Cast(),
				),
				fmt.Sprintf(
					"%v = heap[%v];",
					newProcedure.GetParameter("AccessChar2").Tmp(),
					newProcedure.GetParameter("strPointer2").Temporal.Cast(),
				),
				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					newProcedure.GetParameter("AccessChar1").Tmp(),
					newProcedure.GetLabel("6"),
				),
				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					newProcedure.GetParameter("AccessChar2").Tmp(),
					newProcedure.GetLabel("7"),
				),
				fmt.Sprintf(
					"if (%v < %v) goto %s;",
					newProcedure.GetParameter("AccessChar1").Tmp(),
					newProcedure.GetParameter("AccessChar2").Tmp(),
					newProcedure.GetLabel("8"),
				),
				fmt.Sprintf(
					"if (%v > %v) goto %s;",
					newProcedure.GetParameter("AccessChar1").Tmp(),
					newProcedure.GetParameter("AccessChar2").Tmp(),
					newProcedure.GetLabel("7"),
				),
				fmt.Sprintf(
					"%v = %v + 1;",
					newProcedure.GetParameter("strPointer1").Tmp(),
					newProcedure.GetParameter("strPointer1").Tmp(),
				),
				fmt.Sprintf(
					"%v = %v + 1;",
					newProcedure.GetParameter("strPointer2").Tmp(),
					newProcedure.GetParameter("strPointer2").Tmp(),
				),
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("5").String()),
				newProcedure.GetLabel("6").Declare(),
				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					newProcedure.GetParameter("AccessChar2").Tmp(),
					newProcedure.GetLabel("9"),
				),
				fmt.Sprintf(
					"goto %s;",
					newProcedure.GetLabel("8"),
				),
				newProcedure.GetLabel("9").Declare(),
				fmt.Sprintf(
					"%v = 0;",
					newProcedure.GetParameter("result").Tmp(),
				),
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("10").String()),
				newProcedure.GetLabel("8").Declare(),
				fmt.Sprintf(
					"%v = -1;",
					newProcedure.GetParameter("result").Tmp(),
				),
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("10").String()),
				newProcedure.GetLabel("7").Declare(),
				fmt.Sprintf(
					"%v = 1;",
					newProcedure.GetParameter("result").Tmp(),
				),
				newProcedure.GetLabel("10").Declare(),
			},
			"Comparación de cadenas",
		)

		c.TAC.AddStandard(newProcedure)
	}

	procedure := c.TAC.GetStandard(procedureName)

	// Set left Operator in temporal
	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("strPointer1").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("strPointer2").Tmp(), rightOp.GetValue()),
			"std_compare_str();",
			fmt.Sprintf(
				"if(%v %s 0) goto %s;",
				procedure.GetParameter("result").Temporal,
				op,
				procedure.GetLabel("TrueCondition"),
			),
			fmt.Sprintf(
				"%s = 0;",
				procedure.GetParameter("result2").Temporal,
			),
			fmt.Sprintf("goto %s;", procedure.GetLabel("EndCondition")),
			procedure.GetLabel("TrueCondition").Declare(),
			fmt.Sprintf(
				"%s = 1;",
				procedure.GetParameter("result2").Temporal,
			),
			procedure.GetLabel("EndCondition").Declare(),
		},
		"Comparación de cadenas",
	)

	return &ValueResponse{
		Type:        BooleanTemporal,
		Value:       procedure.GetParameter("result2").Temporal,
		ContextType: TemporalType,
	}
}

func (c *Compiler) PrintBool(name string) {
	newProcedure := NewProcedure(name)

	newProcedure.AddParameters(
		[]*Parameter{
			{
				ExternalName: "HeapPointer",
				Temporal:     c.TAC.NewTemporal(IntTemporal),
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
				newProcedure.GetParameter("HeapPointer").Temporal.Cast(),
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

	c.TAC.AddStandard(newProcedure)
}

func (c *Compiler) PrintString(name string) {
	newProcedure := NewProcedure(name)

	newProcedure.AddParameters(
		[]*Parameter{
			{
				ExternalName: "HeapPointer",
				Temporal:     c.TAC.NewTemporal(IntTemporal),
			},
			{
				ExternalName: "AccessChar",
				Temporal:     c.TAC.NewTemporal(IntTemporal),
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
				newProcedure.GetParameter("AccessChar").Tmp(),
				newProcedure.GetParameter("HeapPointer").Temporal.Cast(),
			),

			// If char is -1, then go to end of string
			fmt.Sprintf(
				"if (%v == -1) goto %s;",
				newProcedure.GetParameter("AccessChar").Tmp(),
				newProcedure.GetLabel("EndString"),
			),

			// Print char
			fmt.Sprintf(
				"printf(\"%%c\",(int) %v);",
				newProcedure.GetParameter("AccessChar").Tmp(),
			),

			// Increase pointer to access char
			fmt.Sprintf(
				"%v = %v + 1;",
				newProcedure.GetParameter("HeapPointer").Tmp(),
				newProcedure.GetParameter("HeapPointer").Tmp(),
			),

			// Go to first label
			fmt.Sprintf("goto %s;", newProcedure.GetLabel("StartString").String()),

			// Declare end of string
			newProcedure.GetLabel("EndString").Declare(),
		},
		"",
	)

	c.TAC.AddStandard(newProcedure)
}

func Print(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	for _, arg := range c.GetArgs(ctx) {
		if arg.Value.Type == StringTemporal {

			if c.TAC.GetStandard("stdprint") == nil {
				c.PrintString("stdprint")
			}

			procedure := c.TAC.GetStandard("stdprint")

			c.TAC.AppendInstructions(
				[]string{
					fmt.Sprintf("%v = %v;", procedure.GetParameter("HeapPointer").Tmp(), arg.Value.GetValue()),
					"stdprint();",
				},
				"Imprimiendo cadena",
			)

		} else if arg.Value.Type == BooleanTemporal {

			if c.TAC.GetStandard("stdprintbool") == nil {
				c.PrintBool("stdprintbool")
			}

			procedure := c.TAC.GetStandard("stdprintbool")

			c.TAC.AppendInstructions(
				[]string{
					fmt.Sprintf("%v = %v;", procedure.GetParameter("HeapPointer").Tmp(), arg.Value.GetValue()),
					"stdprintbool();",
				},
				"Imprimiendo booleano",
			)

		} else {
			c.TAC.AppendInstruction(arg.Value.ToPrint(), "Imprimiendo valor")
		}

		c.TAC.AppendInstruction("printf(\"%c\", 32);", "")
	}

	c.TAC.AppendInstruction("printf(\"%c\", 10);", "")

	return nil
}

func (c *Compiler) ZeroDivision(leftOp, rightOp *ValueResponse, op string) *ValueResponse {

	if c.TAC.GetStandard("std_zero_division") == nil {
		newProcedure := NewProcedure("std_zero_division")

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "operand",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
			},
		)

		newProcedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("NotZeroDivision"),
				c.TAC.NewLabel("EndZeroDivision"),
				c.TAC.NewLabel("IsZeroDivision"),
			},
		)

		// Add code

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf(
					"if (%v != 0) goto %s;",
					newProcedure.GetParameter("operand").Temporal,
					newProcedure.GetLabel("NotZeroDivision"),
				),

				"printf(\"%c\", 77);",
				"printf(\"%c\", 97);",
				"printf(\"%c\", 116);",
				"printf(\"%c\", 104);",
				"printf(\"%c\", 69);",
				"printf(\"%c\", 114);",
				"printf(\"%c\", 114);",
				"printf(\"%c\", 111);",
				"printf(\"%c\", 114);",
				"printf(\"%c\", 10);",
				fmt.Sprintf(
					"%s = 1;",
					newProcedure.GetParameter("operand").Temporal,
				),
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("EndZeroDivision")),
				newProcedure.GetLabel("NotZeroDivision").Declare(),
				fmt.Sprintf(
					"%s = 0;",
					newProcedure.GetParameter("operand").Temporal,
				),
				newProcedure.GetLabel("EndZeroDivision").Declare(),
			},
			"División entre cero",
		)

		c.TAC.AddStandard(newProcedure)
	}

	procedure := c.TAC.GetStandard("std_zero_division")

	newTemporal := c.TAC.NewTemporal(IntTemporal)
	isZeroDivisionLabel := c.TAC.NewLabel("")

	if leftOp.GetType() == FloatTemporal {
		newTemporal.Type = FloatTemporal
	}

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("operand").Tmp(), rightOp.GetValue()),
			"std_zero_division();",
			fmt.Sprintf(
				"if (%v == 1) goto %s;",
				procedure.GetParameter("operand").Temporal,
				isZeroDivisionLabel,
			),
			fmt.Sprintf(
				"%s = %v %s %v;",
				newTemporal,
				leftOp.Cast(),
				op,
				rightOp.Cast(),
			),
			isZeroDivisionLabel.Declare(),
		},
		"División entre cero",
	)

	return &ValueResponse{
		Type:        IntTemporal,
		Value:       newTemporal,
		ContextType: TemporalType,
	}
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
