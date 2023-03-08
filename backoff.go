package backoff

const (
	initialStep      = 0
	linearMultiplier = 1
	nullMultiplier   = 0
)

type BackoffWrapper func(Backoff) Backoff

type Backoff interface {
	NextDelay(int) (int64, error)
}

type NextDelayFunc func(int) (int64, error)

func (n NextDelayFunc) NextDelay(step int) (int64, error) {
	return n(step)
}
