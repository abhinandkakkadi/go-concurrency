package main

import (
	"fmt"
	"sync"
)

// Here race condition will not occur as it is resolved
func main() {
	data := 0
	var memoryAccess sync.Mutex

	go func() {

		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()

	}()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("the value of data: ",data)
	} else {
		fmt.Println("the value of data: ",data)
	}
	memoryAccess.Unlock()

}