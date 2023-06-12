package main

import (
	"fmt"
	"sync"
	"time"
)

func directCall(s string) {

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}

	wg.Done()
}

var wg sync.WaitGroup

func main() {

	wg.Add(3)
	// first goroutine function call
	go directCall("athira")

	//  goroutine with anonymous function
	go func(s string) {
		go directCall(s)
		wg.Done()
	}("geetha")
	// time.Sleep(2*time.Second)

	// go routine with function value call
	fValue := directCall
	go fValue("abhinand")
	wg.Wait()
}
