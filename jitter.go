package backoff

import (
	"crypto/rand"
	"math/big"
)

func EqualJitter(b Backoff) Backoff {
	return BackoffFunc(func(i int) int {
		dur := b.Backoff(i)

		dur = dur / 2

		jitter, err := rand.Int(rand.Reader, big.NewInt(int64(dur)))
		if err != nil {
			panic(err)
		}

		return dur + int(jitter.Int64())
	})
}

func FullJitter(b Backoff) Backoff {
	return BackoffFunc(func(i int) int {
		dur := b.Backoff(i)

		jitter, err := rand.Int(rand.Reader, big.NewInt(int64(dur)))
		if err != nil {
			panic(err)
		}

		return int(jitter.Int64())
	})
}
