package backoff

import (
	"math"
)

type Backoff interface {
	Backoff(int) int
}

type BackoffFunc func(int) int

func (b BackoffFunc) Backoff(step int) int {
	return b(step)
}

const (
	initialStep      = 0
	linearMultiplier = 1
	nullMultiplier   = 0
)

func Exponential() BackoffFunc {
	return BackoffFunc(func(step int) int {
		if step == initialStep {
			return nullMultiplier
		}

		return int(math.Exp2(float64(step - 1)))
	})
}

func Linear() BackoffFunc {
	return BackoffFunc(func(step int) int {
		return step
	})
}

func Fixed() BackoffFunc {
	return BackoffFunc(func(step int) int {
		if step == initialStep {
			return nullMultiplier
		}

		return linearMultiplier
	})
}
