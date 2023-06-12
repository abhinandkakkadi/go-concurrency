package main

import (
	"fmt"
	"sync"
)

// In order to add function inside another function we can do 2 things
// one - use anonymous function
// two - use function expression

// task is
// we need to create a function expression which takes in pointer to waitgroups as arguments and inside that a int will be declared and
// a anonymous concurrent functoin will be declared. The goroutine uses the int data which is present in it's lexical scope.
// we have to find if the anonymous go routine can access the the data even after the enclosing function has returned

// if we execute the program most probably the enclosing function will return before the go routine executes.
// bu still the go routine can access the value of data since it is in it's lexical scope

func main() {

	var wg sync.WaitGroup

	enclosing := func(wg *sync.WaitGroup) {

		var data int
		wg.Add(1)

		go func() {
			defer wg.Done()
			data++
			fmt.Println("the value of data ", data)
		}()

		fmt.Println("returning at this point")
		return
	}

	enclosing(&wg)
	wg.Wait()
	fmt.Println("finished")

}
