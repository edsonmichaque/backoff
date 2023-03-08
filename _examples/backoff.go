package main

import (
	"fmt"
	"time"

	"github.com/edsonmichaque/go-backoff"
)

func main() {
	var backoff backoff.Backoff = backoff.Exponential()

	wrappers := []backoff.BackoffWrapper{
		backoff.MaxAttempts(8),
		backoff.InitialDelay(100 * time.Millisecond),
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
