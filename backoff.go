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

func Exponential() BackoffFunc {
	return BackoffFunc(func(step int) int {
		if step == 0 {
			return 0
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
		if step == 0 {
			return 0
		}

		return 1
	})
}
