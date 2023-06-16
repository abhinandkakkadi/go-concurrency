package main

import (
	"fmt"
	"time"
)


func main() {

	ch := make(chan string)
	
	go func(){
		time.Sleep(3 * time.Second)
		ch <- "abhinand"
	}()

	// here the time.After will sent the current time after 4 second through a receive only channel of type time.Time
	// which ever condition satisfy first will be executed
	select {
	case m := <- ch:
		fmt.Println(m)
	case <- time.After(4 * time.Second):
		fmt.Println("timeout")
	}


 }