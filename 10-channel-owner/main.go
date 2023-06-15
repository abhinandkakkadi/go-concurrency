package main

import "fmt"



func main() {

	owner := func() <-chan int {

		ch := make(chan int)

		go func(){
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		
		return ch
	}

	consumer := func(ch <-chan int) {
		
		for val := range ch {
			fmt.Println("the values are ",val)
		}
	}
	// here we are calling initialize function which return back a receive only channel which is being passed on to  pass()
	// initialize and pass is function expression which can be called in form of a normal function
	// here the initialize is the channel owner so by convention only it can create send and close the channel. 
	// al the other function can only receive from the channel
	ch := owner()
	consumer(ch)
}