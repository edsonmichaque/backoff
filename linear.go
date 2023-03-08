package backoffkit

func Linear() NextDelayFunc {
	return NextDelayFunc(func(step int) (int64, error) {
		return int64(step), nil
	})
}
