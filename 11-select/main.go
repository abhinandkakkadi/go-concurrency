package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "abhinand"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "athira"
	}()

	for i := 0; i < 2; i++ {
		// here whichever one is ready will be executed. In this case since the second go routine have a lesser sleep time
		// that case will be met and printed first and since we have provided a for loop i will wait for the next condition is met
		// which is the data that is sent over the other channel
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		case m2 := <-ch2:
			fmt.Println(m2)
		}

	}

}
