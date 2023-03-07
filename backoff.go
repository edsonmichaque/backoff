package backoff

import (
	"math"
)

type Backoff interface {
	Backoff(int) int64
}

type BackoffFunc func(int) int64

func (b BackoffFunc) Backoff(step int) int64 {
	return b(step)
}

const (
	initialStep      = 0
	linearMultiplier = 1
	nullMultiplier   = 0
)

func Exponential() Backoff {
	return BackoffFunc(func(step int) int64 {
		if step == initialStep {
			return nullMultiplier
		}

		return int64(math.Exp2(float64(step - 1)))
	})
}

func Linear() Backoff {
	return BackoffFunc(func(step int) int64 {
		return int64(step)
	})
}

func Fixed() Backoff {
	return BackoffFunc(func(step int) int64 {
		if step == initialStep {
			return nullMultiplier
		}

		return linearMultiplier
	})
}
