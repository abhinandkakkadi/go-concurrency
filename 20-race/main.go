package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	start := time.Now()
	var t *time.Timer
	ch := make(chan bool)
	t = time.AfterFunc(randomDuration(),func() {
		fmt.Println(time.Now().Sub(start))
		ch <- true
	})
	
	// so here te ch is created until the callback function get's executed,
	// And then it receives value from channel and then reset the time to the duration specified
	// and then the call back is again to scheduled to call after the duration passed through the
	// reset function.
	for time.Since(start) < 5 *time.Second {
		<-ch
		t.Reset(randomDuration())
	}
}


func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}