package main

import (
	"fmt"
	"time"

	"github.com/edsonmichaque/backoffkit"
)

func main() {
	var backoff backoffkit.Backoff = backoffkit.Exponential()

	backoff = backoffkit.MaxAttempts(8)(backoff)
	backoff = backoffkit.InitialDelay(100 * time.Millisecond)(backoff)

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
