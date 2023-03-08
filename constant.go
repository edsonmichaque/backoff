package backoff

func Constant() ComputeDelayFunc {
	return ComputeDelayFunc(func(step int) (int64, error) {
		if step == initialStep {
			return nullMultiplier, nil
		}

		return linearMultiplier, nil
	})
}
