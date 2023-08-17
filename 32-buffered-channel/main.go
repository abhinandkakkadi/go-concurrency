package main

import (
	"fmt"
	"time"
)

// buffered channel advantages

// when you know how many go routines you launched
// limit the number of go routines we launched
// limit the amount of work that's queued up -- this case
 
func listenToChan(ch chan int) {
	
	for {
		//  print a got data message
		// every time a message is received through this channel, a single buffer free space is created
		i := <- ch
		fmt.Println("Got",i,"from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}

}

func main() {

	ch := make(chan int,10)

	go listenToChan(ch)

	// since it is a buffered channel, first 10 values will be sent even if the receiving channel is not ready
	for i := 0; i <= 100; i++ {
		fmt.Println("sending",i,"to channel...")
		ch <- i
		fmt.Println("sent",i,"to channel!!")
	} 

	fmt.Println("Done")
	close(ch)

}