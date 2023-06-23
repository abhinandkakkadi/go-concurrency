package main

import "fmt"

// here we are implementing pipelining
// The name itself says what it is

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

func square(in <-chan int) <-chan int {

	out := make(chan int)
	go func(){
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}

func main() {

	ch := generator(1,2,3,4,5)
	out := square(ch)
	
	for n := range out {
		fmt.Println(n)
	}
}