package main

import "fmt"

// Here a race condition exists as we cannot be sure about if the if condition returns true or not
func main() {
	data := 0

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Println("the value of data: ",data)
	}

}