package main

import "fmt"

// her ech1 can only send to channel since we are specifying this is a receive only channel
func sendMsg(ch1 chan<- string){
	ch1 <- "abhinand"
}


// here ch1 can only receive and ch2 can only send since we have mentioned in that way
func relayMsg(ch1 <-chan string, ch2 chan<- string){
	val := <- ch1
	ch2 <- val
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMsg(ch1)
	go relayMsg(ch1,ch2)

	// this have to be sent from relayMsg
	fmt.Println("final value : ",<- ch2)
}


// var  chan interface{}
// the above declaration is a nil channel, it does not allocate any memory
// if we try to close a nil channel an error will be thrown

// the 