package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var data int

	wg.Add(1)

	// here the function is able to use data due to the property called closure.
	// this function can access all the variable present in it's lexical scope
	go func() {
		defer wg.Done()
		data++
	}()
	wg.Wait()

	fmt.Println("Now I can garantee the  of data is = ", data)
}
