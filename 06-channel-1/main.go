package main

import "fmt"

func main() {

	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	result := <-ch
	fmt.Println("the result is =", result)

}
