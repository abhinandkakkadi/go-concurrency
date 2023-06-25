package main

import (
	"fmt"
	"sync"
)



func generator(done <-chan struct{}, nums ...int) <-chan int {

	out := make(chan int)

	go func(){
		defer close(out)
		for _,n := range nums {
			select {
			case out <- n:
			case <- done:
				return
			}
		}
	}()

	return out
}	


func square(done <-chan struct{},ch <-chan int) <-chan int  {

	out := make(chan int)

	go func(){
		defer close(out)
		for n := range ch {
			select{
				case out <- n*n:
				case <- done:
					return
			} 
		}
	}()

	return out
}

func merge(done <-chan struct{},chs ...<-chan int) <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup

	output := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {

			select {
			case 	out <- n:
			case <- done:
					return
			}
		
		}
	}

	wg.Add(len(chs))
	for _,ch := range chs {
		go output(ch)
	}

	go func(){
		wg.Wait()
		close(out)

	}()

	return out
}

func main() {

	done := make(chan struct{})
	defer close(done)

	in := generator(done,1,2)
	ch1 := square(done,in)
	ch2 := square(done,in)
	out := merge(done,ch1,ch2)

	fmt.Println(<-out)
}

// here even if the receiving channel is not declared the execution will not be blocked as it is sent from a go routine which is independent of main routine
// but conversely if the main go routine is sending from the channel, some one have to receive or the program will be blocked
// empty struct is a common go idiom which is used to synchronize go routine where data might not be sent between go routine. å