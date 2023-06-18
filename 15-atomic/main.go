package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < 50; i++ {

		// if atomic is not given counter variable goes through a race condition since multiple function are
		// trying to access the same value
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddUint64(&counter, 1)
			}

		}()
	}

	wg.Wait()
	fmt.Println(counter)

}
