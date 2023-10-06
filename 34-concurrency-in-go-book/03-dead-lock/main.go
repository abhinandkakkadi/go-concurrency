package main

import (
	"fmt"
	"sync"
	"time"
)


type value struct {

	mu sync.Mutex
	value int

}

var wg sync.WaitGroup

// here we are in a dead lock situation as goroutine 1 has locked a and goroutine 2 has locked b and now it's trying to lock 
// using mutex value which is already in a wait state which will inevitably put it in a wait state.
var printSum = func(a, b *value) {

	defer wg.Done()

	a.mu.Lock()
	defer a.mu.Unlock()

	time.Sleep(2 * time.Second)

	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Println("The values are: ",a.value,b.value)

}

func main() {

	var a, b value
	
	a.value = 1
	b.value = 2

	a.mu = sync.Mutex{}
	b.mu = sync.Mutex{}

	// Add two go routines
	wg.Add(2)
	
	go printSum(&a, &b)
	go printSum(&b, &a)

	// wait for all go routines to be in done state
	wg.Wait()

}