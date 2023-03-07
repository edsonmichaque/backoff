package backoff

import (
	"crypto/rand"
	"math/big"
)

func EqualJitter(b Backoff) BackoffFunc {
	return BackoffFunc(func(i int) int {
		dur := b.Backoff(i)

		dur = dur / 2

		if dur <= 0 {
			return dur
		}

		jitter, err := rand.Int(rand.Reader, big.NewInt(int64(dur)))
		if err != nil {
			panic(err)
		}

		return dur + int(jitter.Int64())
	})
}

func FullJitter(b BackoffFunc) BackoffFunc {
	return BackoffFunc(func(i int) int {
		dur := b(i)

		if dur <= 0 {
			return dur
		}

		jitter, err := rand.Int(rand.Reader, big.NewInt(int64(dur)))
		if err != nil {
			panic(err)
		}

		return int(jitter.Int64())
	})
}
