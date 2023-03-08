package backoffkit

import (
	"math"
)

func Exponential() NextDelayFunc {
	return NextDelayFunc(func(step int) (int64, error) {
		if step == initialStep {
			return nullMultiplier, nil
		}

		return int64(math.Exp2(float64(step - 1))), nil
	})
}
