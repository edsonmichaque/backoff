package backoff

func Constant() NextDelayFunc {
	return NextDelayFunc(func(step int) (int64, error) {
		if step == initialStep {
			return nullMultiplier, nil
		}

		return linearMultiplier, nil
	})
}
