package main

import (
	"fmt"
	"sync"
)

// sync Cond is useful in situation where multiple readers are waiting for a shared resources
// other wise we can just use mutex to lock and unlock -- both are doing he same thing
var sharedRes = make(map[string]interface{})

func main() {

	var wg sync.WaitGroup
	m := sync.Mutex{}

	// conditional require a pointer to mutex as argument
	c := sync.NewCond(&m)

	wg.Add(2)

	go func() {

		defer wg.Done()
		c.L.Lock()
		for len(sharedRes) < 1 {
			wg.Wait()
		}

		fmt.Println("wait for one record : ",sharedRes["one"])
		c.L.Unlock()
	}()

	go func() {

		defer wg.Done()
		c.L.Lock()
		for len(sharedRes) < 2 {
			c.Wait()
		}

		fmt.Println("wait for 2 records : ",sharedRes["two"])
		c.L.Unlock()
	}()

	c.L.Lock()
	sharedRes["one"] = "abhinand"
	sharedRes["two"] = "athira"
	// Broadcast send a broadcast signal to all the go routines which are in it's suspended form to continue it's execution.
	c.Broadcast()
	c.L.Unlock()

	wg.Wait()
	fmt.Println("Done")
}