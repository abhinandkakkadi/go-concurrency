package main

import (
	"fmt"
	"sync"
)

func main() {

	var balance float64 = 0
	var wg sync.WaitGroup
	var m sync.Mutex

	// in the 2 function expression given below the balance is a shared variable
	// i.e if one of them try to add balance to it and other one try to withdraw from it, some anomalies can happen
	// so we are using mutex lock and unlock functionality to protect the shared variable
	// here the balance will be zero, but if the lock is not provided we can't predict the actual output
	deposit := func(amount float64) {
		m.Lock()
		balance += amount
		m.Unlock()
	}

	withdraw := func(amount float64) {
		m.Lock()
		balance -= amount
		m.Unlock()
	}

	// here 100 go routine will be fired by passing values as arguments which ranges from 1 - 100
	// but we can't know in which order the go routine will be executed.
	// but we know that while a go routine is writing or reading on a shared variable (balance) all other go routine will be blocked abd can't enter the critical area
	// critical area -- area between the lock and unlock
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func(i int) {
			defer wg.Done()
			deposit(float64(i))
		}(i)
	}

	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func(i int) {
			defer wg.Done()
			withdraw(float64(i))
		}(i)
	}

	wg.Wait()
	fmt.Println(balance)

}
