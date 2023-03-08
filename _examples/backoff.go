package main

import (
	"fmt"
	"time"

	"github.com/edsonmichaque/backoffkit"
)

func main() {
	var backoff backoffkit.Backoff = backoffkit.Exponential()

	wrappers := []backoffkit.BackoffWrapper{
		backoffkit.MaxAttempts(8),
		backoffkit.InitialDelay(100 * time.Millisecond),
	}

	for _, wrapper := range wrappers {
		backoff = wrapper(backoff)
	}

	go func() {
		for i := 0; ; i++ {
			next, err := backoff.NextDelay(i)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%v %v %v\n", i, next, time.Duration(next))
			time.Sleep(time.Duration(next))
		}
	}()

	t := time.NewTimer(30 * time.Second)

	<-t.C
	fmt.Print("Max timeout reached\n")
}
