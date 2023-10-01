package compiler

import "fmt"

func (c *Compiler) ConcatString(leftOp, rightOp *ValueResponse) *ValueResponse {
	// Check if is the first time that procedure is called, if it is, create the procedure
	procedure := c.TAC.GetStandar("concat")

	if procedure == nil {
		newProcedure := NewProcedure("concat")

		// Add arguments
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
			},
		)

		// Add code

		newProcedure.AddCode(
			[]string{
				fmt.Sprintf("%s = H;", newProcedure.GetArgumentByName("HeapPointer").Tmp()),
				c.TAC.NewLabel().String(),
			},
			"",
		)

		c.TAC.AddProcedure(newProcedure)
	}

	procedure = c.TAC.GetStandar("concat")

	// Set left Operator in temporal
	c.TAC.AppendCode(
		[]string{
			fmt.Sprintf("%v = %v", procedure.GetArgumentByName("leftOp").Tmp(), leftOp.GetValue()),
			fmt.Sprintf("%v = %v", procedure.GetArgumentByName("rightOp").Tmp(), rightOp.GetValue()),
		},
		"",
	)

	return &ValueResponse{}
}
