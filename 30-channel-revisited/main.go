package main

import (
	"fmt"
	"strings"
)

// here we are specifying receive only and send only channel
func shout(ping <-chan string, pong chan<- string) {
	for {

		// ok will be false when channel is closed -- until then the channel is blocked
		s,ok := <-ping
		if !ok {
				//  do something
		}

		pong <- fmt.Sprintf("%s!!!",strings.ToUpper(s))

	}
}

func main() {

	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get a user input
		var userInput string
		_,_ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for response

		response := <- pong

		fmt.Println("Response:",response)

	}

	fmt.Println("All done closing channels.")

	// created channels should always be closed - in order to avoid resource leak
	close(ping)
	close(pong)

}
