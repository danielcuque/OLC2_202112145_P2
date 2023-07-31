package expressions

import (
	"fmt"

	I "OLC2/chore/interfaces"
	U "OLC2/chore/utils"
)

func Add(l, r interface{}) (error, error) {
	leftT := U.GetType(l)
	rightT := U.GetType(r)

	fmt.Println(leftT, rightT)

	r = I.Value{
		Type:       I.INT,
		ParseValue: 0,
	}

	return nil, nil
}
