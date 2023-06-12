package main

import (
	"fmt"
	"sync"
)




var wg sync.WaitGroup

func main() {

	var data int

	wg.Add(1)
	go func(){
		defer wg.Done()
		data++
	}()
	wg.Wait()

	fmt.Println("Now I can garantee the  of data is = ",data)
}