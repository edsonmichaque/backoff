package backoff

const (
	initialStep      = 0
	linearMultiplier = 1
	nullMultiplier   = 0
)

type BackoffWrapper func(Backoff) Backoff

type Backoff interface {
	ComputeDelay(int) (int64, error)
}

type ComputeDelayFunc func(int) (int64, error)

func (n ComputeDelayFunc) ComputeDelay(step int) (int64, error) {
	return n(step)
}
