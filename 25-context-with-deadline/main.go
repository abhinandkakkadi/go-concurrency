package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	val string
}

func main() {

	deadline := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	compute := func(ctx context.Context) <-chan data {

		ch := make(chan data)

		go func() {

			defer close(ch)
			deadline, ok := ctx.Deadline()
			if ok {
				if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
					fmt.Println("deadline is too short, can't complete")
					return
				}
			}

			time.Sleep(50 * time.Millisecond)

			select {
			case ch <- data{"abhinand"}:
			case <-ctx.Done():
				return
			}

		}()

		return ch

	}

	ch := compute(ctx)
	data, ok := <-ch
	if ok {
		fmt.Println("work is complete and data is ", data)
	}

}
