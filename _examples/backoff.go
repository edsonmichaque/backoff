package main

import (
	"fmt"
	//"time"

	"github.com/edsonmichaque/backoff"
)

func main() {
	backoffFunc := func(res int64, b backoff.Backoff) backoff.Backoff {
		return backoff.NextDelayFunc(func(i int) int64 {
			return 1000 * b.NextDelay(i)
		})
	}

	var bo backoff.Backoff = backoffFunc(1000, backoff.Linear())

	bo = backoff.EqualJitter(bo)

	for i := 0; i < 10; i++ {
		next := bo.NextDelay(i)

		fmt.Printf("%v %v %v\n", i, next, next*100)
	}
}
