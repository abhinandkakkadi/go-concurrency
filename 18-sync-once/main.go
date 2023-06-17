package main

import (
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	var once sync.Once

	ok := func(){
		fmt.Println("this function should only be called once")
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		
		go func() {
			defer wg.Done()
			// here since we passed the function expression to the method of Once struct, it will make sure that this function is only called once,
			// Even thought it is called by multiple go routines
			once.Do(ok)
		}()

	}

	wg.Wait()

}