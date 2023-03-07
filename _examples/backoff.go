package main

import (
	"fmt"
	//"time"

	"github.com/edsonmichaque/backoff"
)

func main() {
	backoffFunc := backoff.Exponential()

	backoffFunc = backoff.EqualJitter(backoffFunc)

	for i := 0; i < 100; i++ {
		next := backoffFunc(i)

		fmt.Printf("%v %v %v\n", i, next, next*100)
	}
}
