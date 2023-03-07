package main

import (
	"fmt"
	"time"

	"github.com/edsonmichaque/backoff"
)

func main() {
	var bo backoff.Backoff = backoff.Initialdelay(
		150*time.Millisecond, backoff.MaxAttemps(5, backoff.Exponential()),
	)

	go func() {
		for i := 0; ; i++ {
			next, err := bo.NextDelay(i)
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
