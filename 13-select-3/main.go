package main

import "fmt"

// non blocking select operation
func main() {

	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- "athira"
		}
	}()

	for i := 0; i < 2; i++ {
		// here blocking is not done. if there is a message it will be printed. else the default message will be printed

		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("sorry no message at this time")
		}

	}
}
