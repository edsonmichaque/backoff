package backoff

func Linear() ComputeDelayFunc {
	return ComputeDelayFunc(func(step int) (int64, error) {
		return int64(step), nil
	})
}
