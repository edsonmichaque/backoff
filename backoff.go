package backoff

import (
	"math"
)

type Backoff interface {
	NextDelay(int) int64
}

type NextDelayFunc func(int) int64

func (b NextDelayFunc) NextDelay(step int) int64 {
	return b(step)
}

const (
	initialStep      = 0
	linearMultiplier = 1
	nullMultiplier   = 0
)

func Exponential() NextDelayFunc {
	return NextDelayFunc(func(step int) int64 {
		if step == initialStep {
			return nullMultiplier
		}

		return int64(math.Exp2(float64(step - 1)))
	})
}

func Linear() NextDelayFunc {
	return NextDelayFunc(func(step int) int64 {
		return int64(step)
	})
}

func Fixed() NextDelayFunc {
	return NextDelayFunc(func(step int) int64 {
		if step == initialStep {
			return nullMultiplier
		}

		return linearMultiplier
	})
}
