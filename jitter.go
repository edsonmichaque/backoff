package backoff

import (
	"crypto/rand"
	"math/big"
)

func EqualJitter(b Backoff) Backoff {
	return NextDelayFunc(func(i int) int64 {
		dur := b.NextDelay(i)

		dur = dur / 2

		if dur <= 0 {
			return dur
		}

		jitter, err := rand.Int(rand.Reader, big.NewInt(int64(dur+1)))
		if err != nil {
			panic(err)
		}

		return dur + jitter.Int64()
	})
}

func FullJitter(b Backoff) Backoff {
	return NextDelayFunc(func(i int) int64 {
		dur := b.NextDelay(i)

		if dur <= 0 {
			return dur
		}

		jitter, err := rand.Int(rand.Reader, big.NewInt(int64(dur+1)))
		if err != nil {
			panic(err)
		}

		return jitter.Int64()
	})
}
