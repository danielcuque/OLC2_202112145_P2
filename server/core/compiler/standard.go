package compiler

import (
	V "OLC2/core/values"
	"fmt"
)

func (c *Compiler) ConcatString(leftOp, rightOp *ValueResponse) *ValueResponse {
	procedure := c.TAC.GetStandar("stdconcat")

	if procedure == nil {
		newProcedure := NewProcedure("stdconcat")

		newProcedure.AddArguments(
			[]*Argument{
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
					"%v = Heap[%v];",
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
					"%v = Heap[%v];",
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
			fmt.Sprintf("%v = %v", procedure.GetArgument("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v", procedure.GetArgument("rightOp").Tmp(), rightOp.GetValue()),
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
