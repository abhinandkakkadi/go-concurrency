package main

import (
	"fmt"
	"sync"
)

// here is goroutine is able to access the for loop variable even after the loop finished
// this is because the goroutine forms a closure around it's lexical scope. In this case the for loop (just like inside a function (for loop is kind of like a block only))


func main() {
	var wg sync.WaitGroup
	// In this case all of them print 4 since loop reaches 4 before executing the goroutine
	for i := 0; i <=3 ; i++ {
		wg.Add(1)
		go func(){
			wg.Done()
			fmt.Println(i)
		}()

	}
	wg.Wait()

	// In this case that problem is resolved since we are passing each value in iteration as arguments.
	// what we can't guarantee the order
	for i := 0; i <=3 ; i++ {
		
		wg.Add(1)
		go func(i int){
			wg.Done()
			fmt.Println(i)
		}(i)

	}
	wg.Wait()
}