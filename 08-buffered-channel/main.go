package main

import "fmt"


func main() {

	ch := make(chan int,6)

	// here we want to send data to the channel without blocking, that's why we are using buffered channel
	go func(){
		defer close(ch)
		for i := 0; i < 7; i++ {
			ch <- i
		}

	}()
	
	// here for the first 6 elements 0-5 blocking does not happen since we have a buffer of size 6
	// But after that blocking will happen. In this case when sending value 6 blocking will happen
	for val := range ch {
		fmt.Println(val)
	}
}