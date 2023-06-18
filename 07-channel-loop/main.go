package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// here ch will wait until a value is received and when the channel is close the range automatically break
	// ranging along channel does not give back 2 return values
	for v := range ch {
		fmt.Println(v)
	}

}
