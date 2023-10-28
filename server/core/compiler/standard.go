package compiler

import (
	"OLC2/core/parser"
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
				c.StackPointer.IncreasePointerByOne(),
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
			procedure.Call(),
			fmt.Sprintf("%v = %v;", c.TAC.GetOffSetPointer(), initFreeAddress),
		},
		"Reservando espacio para la instancia de la función",
	)
}

func (c *Compiler) And(leftOp, rightOp *ValueResponse) *ValueResponse {

	procName := "std_and"
	if c.TAC.GetStandard(procName) == nil {
		newProcedure := NewProcedure(procName)

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

	procedure := c.TAC.GetStandard(procName)

	// Set left Operator in temporal

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("rightOp").Tmp(), rightOp.GetValue()),
			procedure.Call(),
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
	procName := "std_or"

	if c.TAC.GetStandard(procName) == nil {
		newProcedure := NewProcedure(procName)

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

	procedure := c.TAC.GetStandard(procName)

	// Set left Operator in temporal

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("rightOp").Tmp(), rightOp.GetValue()),
			procedure.Call(),
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

	procName := "std_concat"

	if c.TAC.GetStandard(procName) == nil {
		newProcedure := NewProcedure(procName)

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
				c.HeapPointer.IncreasePointer(),

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
				c.HeapPointer.IncreasePointer(),
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
				c.HeapPointer.IncreasePointer(),
			},
			"",
		)

		c.TAC.AddStandard(newProcedure)
	}

	procedure := c.TAC.GetStandard(procName)
	returnValue := c.TAC.NewTemporal(StringTemporal)

	// Set left Operator in temporal
	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("rightOp").Tmp(), rightOp.GetValue()),
			procedure.Call(),
			fmt.Sprintf("%v = %v;", returnValue, procedure.GetParameter("HeapPointer").Tmp()),
		},
		"Concatenación de cadenas",
	)

	return &ValueResponse{
		Type:        StringTemporal,
		Value:       returnValue,
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
			procedure.Call(),
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

func (c *Compiler) PrintBool(arg *Argument) *ValueResponse {

	name := "std_print_bool"

	if c.TAC.GetStandard(name) == nil {

		newProcedure := NewProcedure(name)

		newProcedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "HeapPointer",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "Result",
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
					"%v = H;",
					newProcedure.GetParameter("Result").Temporal,
				),

				fmt.Sprintf(
					"if (%v == 1) goto %s;",
					newProcedure.GetParameter("HeapPointer").Temporal.Cast(),
					newProcedure.GetLabel("True").String(),
				),

				fmt.Sprintf("goto %s;", newProcedure.GetLabel("False").String()),
				newProcedure.GetLabel("True").Declare(),
				"heap[(int)H] = 116;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = 114;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = 117;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = 101;",
				c.HeapPointer.IncreasePointer(),
				fmt.Sprintf("goto %s;", newProcedure.GetLabel("EndBool")),
				newProcedure.GetLabel("False").Declare(),
				"heap[(int)H] = 102;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = 97;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = 108;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = 115;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = 101;",
				c.HeapPointer.IncreasePointer(),
				newProcedure.GetLabel("EndBool").Declare(),
				"heap[(int)H] = -1;",
				c.HeapPointer.IncreasePointer(),
			},
			"",
		)

		c.TAC.AddStandard(newProcedure)

	}

	procedure := c.TAC.GetStandard(name)
	returnValue := c.TAC.NewTemporal(StringTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("HeapPointer").Tmp(), arg.Value.GetValue()),
			procedure.Call(),
			fmt.Sprintf("%v = %v;", returnValue, procedure.GetParameter("Result").Tmp()),
		},
		"Imprimiendo booleano",
	)

	return &ValueResponse{
		Type:        StringTemporal,
		Value:       returnValue,
		ContextType: TemporalType,
	}

}

func (c *Compiler) PrintString(value *ValueResponse) {

	name := "std_print"

	if c.TAC.GetStandard(name) == nil {

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

	procedure := c.TAC.GetStandard(name)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("HeapPointer").Tmp(), value.GetValue()),
			procedure.Call(),
		},
		"Imprimiendo cadena",
	)
}

func Print(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	for _, arg := range c.GetArgs(ctx) {
		if arg.Value.Type == StringTemporal {

			c.PrintString(arg.Value)

		} else if arg.Value.Type == BooleanTemporal {

			boolPointer := c.PrintBool(arg)

			c.PrintString(boolPointer)

		} else {
			c.TAC.AppendInstruction(arg.Value.ToPrint(), "Imprimiendo valor")
		}

		c.TAC.AppendInstruction("printf(\"%c\", 32);", "")
	}

	c.TAC.AppendInstruction("printf(\"%c\", 10);", "")

	return nil
}

func (c *Compiler) ZeroDivision(leftOp, rightOp *ValueResponse, op string) *ValueResponse {

	procName := "std_zero_division"

	if c.TAC.GetStandard(procName) == nil {
		newProcedure := NewProcedure(procName)

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

	procedure := c.TAC.GetStandard(procName)

	newTemporal := c.TAC.NewTemporal(IntTemporal)
	isZeroDivisionLabel := c.TAC.NewLabel("")

	if leftOp.GetType() == FloatTemporal {
		newTemporal.Type = FloatTemporal
	}

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("operand").Tmp(), rightOp.GetValue()),
			procedure.Call(),
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

	args := c.GetArgs(ctx)

	if len(args) <= 0 {
		return nil
	}

	if args[0].GetType() == StringTemporal {
		return c.StringToNum(args[0].Value)
	}

	newTemporal := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = (int) %v;", newTemporal, args[0].Value.GetValue()),
		},
		"Convirtiendo a entero",
	)

	return &ValueResponse{
		Type:        IntTemporal,
		Value:       newTemporal,
		ContextType: TemporalType,
	}
}

func Float(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	args := c.GetArgs(ctx)

	if len(args) <= 0 {
		return nil
	}

	return c.StringToFloat(args[0].Value)
}

func (c *Compiler) StringToFloat(value *ValueResponse) *ValueResponse {

	procedureName := "std_string_to_float"

	if c.TAC.GetStandard(procedureName) == nil {
		procedure := NewProcedure(procedureName)

		procedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "HeapPointer",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "HeapPointerValue",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "IntegerPart",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "DecimalPart",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "InitialDecimalPartPointer",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "DecimalCounter",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
			},
		)

		procedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("LookUpPoint"),
				c.TAC.NewLabel("EndLookUpPoint"),
				c.TAC.NewLabel("DecimalCounter"),
				c.TAC.NewLabel("EndDecimalCounter"),
				c.TAC.NewLabel("End"),
			},
		)

		c.TAC.Procedure = procedure

		integerPart := c.StringToNum(&ValueResponse{
			Type:        IntTemporal,
			Value:       procedure.GetParameter("HeapPointer").Temporal,
			ContextType: TemporalType,
		})

		procedure.AddCode(
			[]string{
				fmt.Sprintf(
					"%v = %v;",
					procedure.GetParameter("IntegerPart").Temporal,
					integerPart.GetValue(),
				),
				fmt.Sprintf(
					"%v = 1;",
					procedure.GetParameter("DecimalCounter").Temporal,
				),
				procedure.GetLabel("LookUpPoint").Declare(),

				fmt.Sprintf(
					"%v = heap[%v];",
					procedure.GetParameter("HeapPointerValue").Temporal,
					procedure.GetParameter("HeapPointer").Temporal.Cast(),
				),

				fmt.Sprintf(
					"if (%v == 46) goto %s;",
					procedure.GetParameter("HeapPointerValue").Temporal.Cast(),
					procedure.GetLabel("EndLookUpPoint"),
				),

				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					procedure.GetParameter("HeapPointerValue").Temporal.Cast(),
					procedure.GetLabel("End"),
				),

				fmt.Sprintf(
					"%v = %v + 1;",
					procedure.GetParameter("HeapPointer").Temporal,
					procedure.GetParameter("HeapPointer").Temporal,
				),

				fmt.Sprintf(
					"goto %s;",
					procedure.GetLabel("LookUpPoint").String(),
				),

				procedure.GetLabel("EndLookUpPoint").Declare(),
				fmt.Sprintf(
					"%v = %v + 1;",
					procedure.GetParameter("HeapPointer").Temporal,
					procedure.GetParameter("HeapPointer").Temporal,
				),
			},
			"",
		)

		decimalPart := c.StringToNum(&ValueResponse{
			Type:        IntTemporal,
			Value:       procedure.GetParameter("HeapPointer").Temporal,
			ContextType: TemporalType,
		})

		procedure.AddCode(
			[]string{
				fmt.Sprintf(
					"%v = %v;",
					procedure.GetParameter("DecimalPart").Temporal,
					decimalPart.GetValue(),
				),
				procedure.GetLabel("DecimalCounter").Declare(),

				fmt.Sprintf(
					"%v = heap[%v];",
					procedure.GetParameter("HeapPointerValue").Temporal,
					procedure.GetParameter("HeapPointer").Temporal.Cast(),
				),

				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					procedure.GetParameter("HeapPointerValue").Temporal,
					procedure.GetLabel("EndDecimalCounter"),
				),

				fmt.Sprintf(
					"%v = %v + 1;",
					procedure.GetParameter("HeapPointer").Temporal,
					procedure.GetParameter("HeapPointer").Temporal,
				),

				fmt.Sprintf(
					"%v = %v * 10;",
					procedure.GetParameter("DecimalCounter").Temporal,
					procedure.GetParameter("DecimalCounter").Temporal,
				),

				fmt.Sprintf(
					"goto %s;",
					procedure.GetLabel("DecimalCounter").String(),
				),

				procedure.GetLabel("EndDecimalCounter").Declare(),
				fmt.Sprintf(
					"%v = %v / %v;",
					procedure.GetParameter("DecimalPart").Temporal,
					procedure.GetParameter("DecimalPart").Temporal,
					procedure.GetParameter("DecimalCounter").Temporal,
				),

				fmt.Sprintf(
					"%v = %v + %v;",
					procedure.GetParameter("IntegerPart").Temporal,
					procedure.GetParameter("IntegerPart").Temporal,
					procedure.GetParameter("DecimalPart").Temporal,
				),

				procedure.GetLabel("End").Declare(),
			},
			"",
		)

		c.TAC.UnsetProcedure()

		c.TAC.AddStandard(procedure)
	}

	procedure := c.TAC.GetStandard(procedureName)

	returnValue := c.TAC.NewTemporal(FloatTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("HeapPointer").Tmp(), value.GetValue()),
			procedure.Call(),
			fmt.Sprintf("%v = %v;", returnValue, procedure.GetParameter("IntegerPart").Temporal),
		},
		"Convirtiendo a flotante",
	)

	return &ValueResponse{
		Type:        FloatTemporal,
		Value:       returnValue,
		ContextType: TemporalType,
	}
}

func (c *Compiler) StringToNum(value *ValueResponse) *ValueResponse {
	procedureName := "std_string_to_num"

	if c.TAC.GetStandard(procedureName) == nil {
		procedure := NewProcedure(procedureName)

		procedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "BaseChar",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "HeapPointer",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "HeapPointerValue",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "Result",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "IsNegative",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "Digit",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
			},
		)

		procedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("NotNegative"),
				c.TAC.NewLabel("StartString"),
				c.TAC.NewLabel("EndString"),
				c.TAC.NewLabel("End"),
			},
		)

		procedure.AddCode(
			[]string{
				fmt.Sprintf(
					"%v = 48;",
					procedure.GetParameter("BaseChar").Temporal,
				),
				fmt.Sprintf(
					"%v = 0;",
					procedure.GetParameter("Result").Temporal,
				),
				fmt.Sprintf(
					"%v = heap[%v];",
					procedure.GetParameter("HeapPointerValue").Temporal,
					procedure.GetParameter("HeapPointer").Temporal.Cast(),
				),

				// Check if is negative

				fmt.Sprintf(
					"if (%v != 45) goto %s;",
					procedure.GetParameter("HeapPointerValue").Temporal.Cast(),
					procedure.GetLabel("NotNegative"),
				),

				fmt.Sprintf(
					"%v = 1;",
					procedure.GetParameter("IsNegative").Temporal,
				),

				fmt.Sprintf(
					"%v = %v + 1;",
					procedure.GetParameter("HeapPointer").Temporal,
					procedure.GetParameter("HeapPointer").Temporal,
				),

				fmt.Sprintf(
					"goto %s;",
					procedure.GetLabel("StartString"),
				),

				procedure.GetLabel("NotNegative").Declare(),

				fmt.Sprintf(
					"%v = 0;",
					procedure.GetParameter("IsNegative").Temporal,
				),

				procedure.GetLabel("StartString").Declare(),

				fmt.Sprintf(
					"%v = heap[%v];",
					procedure.GetParameter("HeapPointerValue").Temporal,
					procedure.GetParameter("HeapPointer").Temporal.Cast(),
				),

				fmt.Sprintf(
					"if (%v == -1) goto %s;",
					procedure.GetParameter("HeapPointerValue").Temporal.Cast(),
					procedure.GetLabel("EndString"),
				),

				fmt.Sprintf(
					"if (%v == 46) goto %s;",
					procedure.GetParameter("HeapPointerValue").Temporal.Cast(),
					procedure.GetLabel("EndString"),
				),

				fmt.Sprintf(
					"%v = %v - %v;",
					procedure.GetParameter("HeapPointerValue").Temporal,
					procedure.GetParameter("HeapPointerValue").Temporal,
					procedure.GetParameter("BaseChar").Temporal,
				),

				fmt.Sprintf(
					"%v = %v * 10;",
					procedure.GetParameter("Result").Temporal,
					procedure.GetParameter("Result").Temporal,
				),

				fmt.Sprintf(
					"%v = %v + %v;",
					procedure.GetParameter("Result").Temporal,
					procedure.GetParameter("Result").Temporal,
					procedure.GetParameter("HeapPointerValue").Temporal,
				),

				fmt.Sprintf(
					"%v = %v + 1;",
					procedure.GetParameter("HeapPointer").Temporal,
					procedure.GetParameter("HeapPointer").Temporal,
				),

				fmt.Sprintf(
					"goto %s;",
					procedure.GetLabel("StartString"),
				),

				procedure.GetLabel("EndString").Declare(),

				fmt.Sprintf(
					"if (%v != 1) goto %s;",
					procedure.GetParameter("IsNegative").Temporal,
					procedure.GetLabel("End"),
				),

				fmt.Sprintf(
					"%v = %v * -1;",
					procedure.GetParameter("Result").Temporal,
					procedure.GetParameter("Result").Temporal,
				),

				procedure.GetLabel("End").Declare(),
			},
			"String a número",
		)

		c.TAC.AddStandard(procedure)
	}

	procedure := c.TAC.GetStandard(procedureName)

	returnValue := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("HeapPointer").Tmp(), value.GetValue()),
			procedure.Call(),
			fmt.Sprintf("%v = %v;", returnValue, procedure.GetParameter("Result").Temporal),
		},
		"String a número",
	)

	return &ValueResponse{
		Type:        IntTemporal,
		Value:       returnValue,
		ContextType: TemporalType,
	}
}

func (c *Compiler) NumberToString(value *ValueResponse) *ValueResponse {

	procName := "std_num_to_string"

	if c.TAC.GetStandard(procName) == nil {
		procedure := NewProcedure(procName)

		procedure.AddLabels(
			[]*Label{
				c.TAC.NewLabel("L5"),
				c.TAC.NewLabel("L6"),
				c.TAC.NewLabel("L7"),
				c.TAC.NewLabel("L8"),
				c.TAC.NewLabel("L9"),
				c.TAC.NewLabel("L10"),
				c.TAC.NewLabel("L11"),
			},
		)

		procedure.AddParameters(
			[]*Parameter{
				{
					ExternalName: "t2",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "t5",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "t6",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "t7",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
			},
		)

		procedure.AddCode(
			[]string{
				fmt.Sprintf(
					"%v = H;",
					procedure.GetParameter("t5").Temporal,
				),

				fmt.Sprintf(
					"if (%v > 0) goto %s;",
					procedure.GetParameter("t2").Temporal.Cast(),
					procedure.GetLabel("L5"),
				),

				"heap[(int)H] = 45;",
				c.HeapPointer.IncreasePointer(),

				fmt.Sprintf(
					"%v = %v * -1;",
					procedure.GetParameter("t2").Temporal,
					procedure.GetParameter("t2").Temporal,
				),

				procedure.GetLabel("L5").Declare(),

				fmt.Sprintf(
					"if(%v != 0) goto %s;",
					procedure.GetParameter("t2").Temporal,
					procedure.GetLabel("L6"),
				),
				"heap[(int)H] = 48;",
				c.HeapPointer.IncreasePointer(),

				"heap[(int)H] = 0;",
				c.HeapPointer.IncreasePointer(),

				fmt.Sprintf("goto %s;", procedure.GetLabel("L11")),
				procedure.GetLabel("L6").Declare(),
				fmt.Sprintf(
					"%v = H - 1;",
					procedure.GetParameter("t6").Temporal,
				),
				fmt.Sprintf(
					"%v = %v;",
					procedure.GetParameter("t7").Temporal,
					procedure.GetParameter("t2").Temporal,
				),
				procedure.GetLabel("L7").Declare(),
				fmt.Sprintf(
					"if (%v == 0) goto %s;",
					procedure.GetParameter("t7").Temporal,
					procedure.GetLabel("L8"),
				),
				fmt.Sprintf(
					"%v = %v / 10;",
					procedure.GetParameter("t7").Temporal,
					procedure.GetParameter("t7").Temporal,
				),
				fmt.Sprintf(
					"%v = %v;",
					procedure.GetParameter("t7").Temporal,
					procedure.GetParameter("t7").Temporal.Cast(),
				),
				fmt.Sprintf(
					"%v = %v + 1;",
					procedure.GetParameter("t6").Temporal,
					procedure.GetParameter("t6").Temporal,
				),
				fmt.Sprintf("goto %s;", procedure.GetLabel("L7")),
				procedure.GetLabel("L8").Declare(),
				fmt.Sprintf(
					"H = %v + 1;",
					procedure.GetParameter("t6").Temporal,
				),
				procedure.GetLabel("L9").Declare(),
				fmt.Sprintf(
					"if (%v == 0) goto %s;",
					procedure.GetParameter("t2").Temporal,
					procedure.GetLabel("L10"),
				),
				fmt.Sprintf(
					"%v = %s %% 10;",
					procedure.GetParameter("t7").Temporal,
					procedure.GetParameter("t2").Temporal.Cast(),
				),
				fmt.Sprintf(
					"%v = %v + 48;",
					procedure.GetParameter("t7").Temporal,
					procedure.GetParameter("t7").Temporal,
				),
				fmt.Sprintf(
					"heap[%v] = %v;",
					procedure.GetParameter("t6").Temporal.Cast(),
					procedure.GetParameter("t7").Temporal,
				),
				fmt.Sprintf(
					"%v = %v - 1;",
					procedure.GetParameter("t6").Temporal,
					procedure.GetParameter("t6").Temporal,
				),
				fmt.Sprintf(
					"%v = %v / 10;",
					procedure.GetParameter("t2").Temporal,
					procedure.GetParameter("t2").Temporal,
				),

				fmt.Sprintf(
					"%v = %v;",
					procedure.GetParameter("t2").Temporal,
					procedure.GetParameter("t2").Temporal.Cast(),
				),

				fmt.Sprintf("goto %s;", procedure.GetLabel("L9")),

				procedure.GetLabel("L10").Declare(),

				"heap[(int)H] = -1;",

				c.HeapPointer.IncreasePointer(),

				procedure.GetLabel("L11").Declare(),

				fmt.Sprintf(
					"%v = %v;",
					procedure.GetParameter("t2").Temporal,
					procedure.GetParameter("t5").Temporal,
				),
			},
			"",
		)

		c.TAC.AddStandard(procedure)
	}

	procedure := c.TAC.GetStandard(procName)
	returnValue := c.TAC.NewTemporal(StringTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("t2").Tmp(), value.GetValue()),
			procedure.Call(),
			fmt.Sprintf("%v = %v;", returnValue, procedure.GetParameter("t2").Temporal),
		},
		"Convirtiendo a cadena",
	)

	return &ValueResponse{
		Type:        StringTemporal,
		Value:       returnValue,
		ContextType: TemporalType,
	}
}

func (c *Compiler) FloatToString(value *ValueResponse) *ValueResponse {

	name := "std_float_to_string"

	if c.TAC.GetStandard(name) == nil {
		proc := NewProcedure(name)

		proc.AddParameters(
			[]*Parameter{
				{
					ExternalName: "Value",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "IntPart",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "DecimalPart",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "DotPointer",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
			},
		)

		proc.AddLabels(
			[]*Label{
				c.TAC.NewLabel("IsNotNegative"),
			},
		)

		proc.AddCode(
			[]string{
				fmt.Sprintf(
					"%v = (int)%v;",
					proc.GetParameter("IntPart").Temporal,
					proc.GetParameter("Value").Temporal,
				),
				fmt.Sprintf(
					"%v = %v - %v;",
					proc.GetParameter("DecimalPart").Temporal,
					proc.GetParameter("Value").Temporal,
					proc.GetParameter("IntPart").Temporal,
				),

				fmt.Sprintf(
					"%v = %v * 1000;",
					proc.GetParameter("DecimalPart").Temporal,
					proc.GetParameter("DecimalPart").Temporal,
				),

				fmt.Sprintf(
					"if (%v >= 0) goto %s;",
					proc.GetParameter("Value").Temporal,
					proc.GetLabel("IsNotNegative"),
				),

				fmt.Sprintf(
					"%v = %v * -1;",
					proc.GetParameter("DecimalPart").Temporal,
					proc.GetParameter("DecimalPart").Temporal,
				),

				proc.GetLabel("IsNotNegative").Declare(),
			},
			"",
		)

		c.TAC.Procedure = proc

		integerPart := c.NumberToString(&ValueResponse{
			Type:        IntTemporal,
			Value:       proc.GetParameter("IntPart").Temporal,
			ContextType: TemporalType,
		})

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf(
					"%v = H;",
					proc.GetParameter("DotPointer").Temporal,
				),

				"heap[(int)H] = 46;",
				c.HeapPointer.IncreasePointer(),
				"heap[(int)H] = -1;",
				c.HeapPointer.IncreasePointer(),
			},
			"",
		)

		concatInteger := c.ConcatString(integerPart, &ValueResponse{
			Type:        CharTemporal,
			Value:       proc.GetParameter("DotPointer").Temporal,
			ContextType: TemporalType,
		})

		decimalPart := c.NumberToString(&ValueResponse{
			Type:        IntTemporal,
			Value:       proc.GetParameter("DecimalPart").Temporal,
			ContextType: TemporalType,
		})

		concatDecimal := c.ConcatString(concatInteger, decimalPart)

		c.TAC.AppendInstructions(
			[]string{
				fmt.Sprintf(
					"%v = %v;",
					proc.GetParameter("Value").Temporal,
					concatDecimal.GetValue(),
				),
			},
			"",
		)

		c.TAC.UnsetProcedure()
		c.TAC.AddStandard(proc)
	}

	procedure := c.TAC.GetStandard(name)

	returnValue := c.TAC.NewTemporal(StringTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("Value").Tmp(), value.GetValue()),
			procedure.Call(),
			fmt.Sprintf("%v = %v;", returnValue, procedure.GetParameter("Value").Temporal),
		},
		"Convirtiendo a cadena",
	)

	return &ValueResponse{
		Type:        StringTemporal,
		Value:       returnValue,
		ContextType: TemporalType,
	}
}

func String(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	args := c.GetArgs(ctx)

	if len(args) <= 0 {
		return nil
	}

	arg := args[0]

	if arg.GetType() == IntTemporal {
		return c.NumberToString(arg.Value)
	}

	if arg.GetType() == FloatTemporal {
		return c.FloatToString(arg.Value)
	}

	if arg.GetType() == BooleanTemporal {
		return c.PrintBool(arg)
	}

	return nil
}

func Append(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	name := "std_append"

	if c.TAC.GetStandard(name) == nil {
		prc := NewProcedure(name)

		// Needd to copy all vector value

		// Get size of vector in position 0 of vector
		// [size, isEmpty, value1, value2, value3, ...]

		c.TAC.Procedure = prc

		response := c.InitNewMatrix()

		prc.AddParameters(
			[]*Parameter{
				{
					ExternalName: "vectorAddress",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "size",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "appendValue",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "newVectorAddress",
					Temporal:     response[0],
				},
				{
					ExternalName: "vectorAddressValue",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "vectorAccessPosition",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "counter",
					Temporal:     response[1],
				},
				{
					ExternalName: "IsEmptyAddress",
					Temporal:     response[2],
				},
			},
		)

		prc.AddLabels(
			[]*Label{
				c.TAC.NewLabel("Start"),
				c.TAC.NewLabel("End"),
			},
		)

		prc.AddCode(
			[]string{
				// Get size of the old vector, this will help us to copy all values and append the new one at the end

				// Obtenemos el valor del tamaño del vector viejo
				fmt.Sprintf(
					"%v = heap[(int)%v];",
					prc.GetParameter("counter").Temporal,
					prc.GetParameter("vectorAddress").Temporal,
				),

				// Iniciamos el contador en 0
				fmt.Sprintf(
					"%v = 0;",
					prc.GetParameter("size").Temporal,
				),

				// Obtenemos la dirección de donde inician los valores del vector
				fmt.Sprintf(
					"%v = %v + 2;",
					prc.GetParameter("vectorAddress").Temporal,
					prc.GetParameter("vectorAddress").Temporal,
				),

				// Declaramos la etiqueta de inicio del loop para copiar los valores
				prc.GetLabel("Start").Declare(),

				//
				fmt.Sprintf(
					"if (%v == %s) goto %s;",
					prc.GetParameter("size").Temporal,
					prc.GetParameter("counter").Temporal,
					prc.GetLabel("End"),
				),

				fmt.Sprintf(
					"%v = %v;",
					prc.GetParameter("vectorAccessPosition").Temporal,
					prc.GetParameter("size").Temporal,
				),

				fmt.Sprintf(
					"%v = %v + %v;",
					prc.GetParameter("vectorAccessPosition").Temporal,
					prc.GetParameter("vectorAccessPosition").Temporal,
					prc.GetParameter("vectorAddress").Temporal,
				),

				fmt.Sprintf(
					"%v = heap[(int)%v];",
					prc.GetParameter("vectorAddressValue").Temporal,
					prc.GetParameter("vectorAccessPosition").Temporal,
				),

				fmt.Sprintf(
					"heap[(int)H] = %v;",
					prc.GetParameter("vectorAddressValue").Temporal,
				),

				c.HeapPointer.IncreasePointer(),

				fmt.Sprintf(
					"%v = %v + 1;",
					prc.GetParameter("size").Temporal,
					prc.GetParameter("size").Temporal,
				),

				fmt.Sprintf(
					"goto %s;",
					prc.GetLabel("Start"),
				),

				prc.GetLabel("End").Declare(),

				// El nuevo valor del contador es contadorViejo + 1
				fmt.Sprintf(
					"%v = %v + 1;",
					prc.GetParameter("counter").Temporal,
					prc.GetParameter("counter").Temporal,
				),

				// Añadimos el valor al final del vector
				fmt.Sprintf(
					"heap[(int)H] = %v;",
					prc.GetParameter("appendValue").Temporal,
				),
				c.HeapPointer.IncreasePointer(),
			},
			"",
		)

		c.DeclareMatrixProps(response[0], response[1], response[2])

		c.TAC.UnsetProcedure()
		c.TAC.AddStandard(prc)
	}

	temporalResponse, value := c.GetTemporalResponse(ctx)

	if value == nil || temporalResponse == nil {
		return nil
	}

	args := c.GetArgs(ctx)

	if len(args) == 0 {
		return nil
	}

	procedure := c.TAC.GetStandard(name)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("vectorAddress").Tmp(), temporalResponse),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("appendValue").Tmp(), args[0].GetValue().GetValue()),
			procedure.Call(),
			fmt.Sprintf("stack[(int)%s] = %s;", c.TAC.GetValueAddress(value), procedure.GetParameter("newVectorAddress").Temporal),
		},
		"Función append",
	)

	return nil
}

func Remove(c *Compiler, ctx *parser.FunctionCallContext) interface{} {

	args := c.GetArgs(ctx)

	if len(args) == 0 {
		return nil
	}

	temporalResponse, value := c.GetTemporalResponse(ctx)

	if value == nil || temporalResponse == nil {
		return nil
	}

	returnValue := c.RemoveStandard(args[0].GetValue(), temporalResponse)

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[(int)%s] = %s;", c.TAC.GetValueAddress(value), returnValue.GetValue()),
		"",
	)

	return nil
}

func (c *Compiler) RemoveStandard(arg *ValueResponse, temporalResponse *Temporal) *ValueResponse {

	name := "std_remove"

	if c.TAC.GetStandard(name) == nil {
		prc := NewProcedure(name)

		prevProcedure := c.TAC.Procedure

		c.TAC.Procedure = prc

		response := c.InitNewMatrix()

		prc.AddParameters(
			[]*Parameter{
				{
					ExternalName: "vectorAddress",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "size",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "removeIndex",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "newVectorAddress",
					Temporal:     response[0],
				},
				{
					ExternalName: "vectorAddressValue",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "vectorAccessPosition",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "counter",
					Temporal:     response[1],
				},
				{
					ExternalName: "IsEmptyAddress",
					Temporal:     response[2],
				},
			},
		)

		prc.AddLabels(
			[]*Label{
				c.TAC.NewLabel("Start"),
				c.TAC.NewLabel("End"),
				c.TAC.NewLabel("Skip"),
			},
		)

		prc.AddCode(
			[]string{
				// Get size of the old vector, this will help us to copy all values and append the new one at the end

				// Obtenemos el valor del tamaño del vector viejo
				fmt.Sprintf(
					"%v = heap[(int)%v];",
					prc.GetParameter("counter").Temporal,
					prc.GetParameter("vectorAddress").Temporal,
				),

				// Iniciamos el contador en 0
				fmt.Sprintf(
					"%v = 0;",
					prc.GetParameter("size").Temporal,
				),

				// Obtenemos la dirección de donde inician los valores del vector
				fmt.Sprintf(
					"%v = %v + 2;",
					prc.GetParameter("vectorAddress").Temporal,
					prc.GetParameter("vectorAddress").Temporal,
				),

				// Declaramos la etiqueta de inicio del loop para copiar los valores
				prc.GetLabel("Start").Declare(),

				//
				fmt.Sprintf(
					"if (%v == %s) goto %s;",
					prc.GetParameter("size").Temporal,
					prc.GetParameter("counter").Temporal,
					prc.GetLabel("End"),
				),

				fmt.Sprintf(
					"if (%v == %v) goto %s;",
					prc.GetParameter("size").Temporal,
					prc.GetParameter("removeIndex").Temporal,
					prc.GetLabel("Skip"),
				),

				fmt.Sprintf(
					"%v = %v;",
					prc.GetParameter("vectorAccessPosition").Temporal,
					prc.GetParameter("size").Temporal,
				),

				fmt.Sprintf(
					"%v = %v + %v;",
					prc.GetParameter("vectorAccessPosition").Temporal,
					prc.GetParameter("vectorAccessPosition").Temporal,
					prc.GetParameter("vectorAddress").Temporal,
				),

				fmt.Sprintf(
					"%v = heap[(int)%v];",
					prc.GetParameter("vectorAddressValue").Temporal,
					prc.GetParameter("vectorAccessPosition").Temporal,
				),

				fmt.Sprintf(
					"heap[(int)H] = %v;",
					prc.GetParameter("vectorAddressValue").Temporal,
				),

				c.HeapPointer.IncreasePointer(),

				prc.GetLabel("Skip").Declare(),

				fmt.Sprintf(
					"%v = %v + 1;",
					prc.GetParameter("size").Temporal,
					prc.GetParameter("size").Temporal,
				),

				fmt.Sprintf(
					"goto %s;",
					prc.GetLabel("Start"),
				),

				prc.GetLabel("End").Declare(),
			},
			"",
		)

		c.DeclareMatrixProps(response[0], response[1], response[2])

		c.TAC.UnsetProcedure()
		c.TAC.Procedure = prevProcedure
		c.TAC.AddStandard(prc)
	}

	procedure := c.TAC.GetStandard(name)
	returnValue := c.TAC.NewTemporal(IntTemporal)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", procedure.GetParameter("vectorAddress").Tmp(), temporalResponse),
			fmt.Sprintf("%v = %v;", procedure.GetParameter("removeIndex").Tmp(), arg.GetValue()),
			procedure.Call(),
			fmt.Sprintf("%v = %v;", returnValue, procedure.GetParameter("newVectorAddress").Temporal),
		},
		"Función append",
	)

	return &ValueResponse{
		Type:        IntTemporal,
		Value:       returnValue,
		ContextType: TemporalType,
	}
}

func (c *Compiler) RemoveLastStandar(temporalResponse *Temporal) *ValueResponse {

	name := "std_remove_last"

	if c.TAC.GetStandard(name) == nil {
		prc := NewProcedure(name)

		c.TAC.Procedure = prc

		prc.AddParameters(
			[]*Parameter{
				{
					ExternalName: "vectorAddress",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "size",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
				{
					ExternalName: "newVectorAddress",
					Temporal:     c.TAC.NewTemporal(IntTemporal),
				},
			},
		)

		prc.AddCode(
			[]string{
				fmt.Sprintf(
					"%v = heap[(int)%v];",
					prc.GetParameter("size").Temporal,
					prc.GetParameter("vectorAddress").Temporal,
				),

				fmt.Sprintf(
					"%v = %v - 1;",
					prc.GetParameter("size").Temporal,
					prc.GetParameter("size").Temporal,
				),
			},
			"",
		)

		response := c.RemoveStandard(
			&ValueResponse{
				Type:        IntTemporal,
				Value:       prc.GetParameter("size").Temporal,
				ContextType: TemporalType,
			},
			prc.GetParameter("vectorAddress").Temporal,
		)

		prc.AddCode(
			[]string{
				fmt.Sprintf(
					"%v = %v;",
					prc.GetParameter("newVectorAddress").Temporal,
					response.GetValue(),
				),
			},
			"",
		)

		c.TAC.UnsetProcedure()
		c.TAC.AddStandard(prc)
	}

	returnValue := c.TAC.NewTemporal(IntTemporal)
	prc := c.TAC.GetStandard(name)

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%v = %v;", prc.GetParameter("vectorAddress").Tmp(), temporalResponse),
			prc.Call(),
			fmt.Sprintf("%v = %v;", returnValue, prc.GetParameter("newVectorAddress").Temporal),
		},
		"Remover el último elemento de un vector",
	)

	return &ValueResponse{
		Type:        IntTemporal,
		Value:       returnValue,
		ContextType: TemporalType,
	}
}

func RemoveLast(c *Compiler, ctx *parser.FunctionCallContext) interface{} {
	args := c.GetArgs(ctx)

	if len(args) > 0 {
		return nil
	}

	temporalResponse, value := c.GetTemporalResponse(ctx)

	if value == nil || temporalResponse == nil {
		return nil
	}

	returnValue := c.RemoveLastStandar(temporalResponse)

	c.TAC.AppendInstruction(
		fmt.Sprintf("stack[(int)%s] = %s;", c.TAC.GetValueAddress(value), returnValue.GetValue()),
		"",
	)

	return nil
}

func (c *Compiler) GetTemporalResponse(ctx *parser.FunctionCallContext) (*Temporal, *Value) {
	id, props := c.GetIds(ctx)

	value := c.Env.GetValue(id)

	if value == nil {
		return nil, nil
	}

	baseTemporal := c.TAC.NewTemporal(value.GetType())

	c.TAC.AppendInstructions(
		[]string{
			fmt.Sprintf("%s = stack[(int)%s];", baseTemporal, c.TAC.GetValueAddress(value)),
		},
		fmt.Sprintf("Acceso a la variable '%s'", id),
	)

	props = props[1 : len(props)-1]

	return c.GetProps(value, props, baseTemporal), value
}
