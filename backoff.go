package backoff

import (
	"math"
)

const (
	initialStep      = 0
	linearMultiplier = 1
	nullMultiplier   = 0
)

type Backoff interface {
	NextDelay(int) (int64, error)
}

type NextDelayFunc func(int) (int64, error)

func (n NextDelayFunc) NextDelay(step int) (int64, error) {
	return n(step)
}

func Exponential() NextDelayFunc {
	return NextDelayFunc(func(step int) (int64, error) {
		if step == initialStep {
			return nullMultiplier, nil
		}

		return int64(math.Exp2(float64(step - 1))), nil
	})
}

func Linear() NextDelayFunc {
	return NextDelayFunc(func(step int) (int64, error) {
		return int64(step), nil
	})
}

func Fixed() NextDelayFunc {
	return NextDelayFunc(func(step int) (int64, error) {
		if step == initialStep {
			return nullMultiplier, nil
		}

		return linearMultiplier, nil
	})
}
