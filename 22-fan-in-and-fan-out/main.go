package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _,n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
} 

// here we are assuming that square function is very expensive
// so we are fanning out to the work to multiple go routine
// and using the merge function we are fanning in
func square(in <-chan int) <-chan int {
	
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n*n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _,c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	in := generator(2,3)
	c1 := square(in)
	c2 := square(in)

	for n := range merge(c1,c2) {
		fmt.Println(n)
	}
}