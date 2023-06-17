package main

import "sync"

var sharedRes = make(map[string]interface{})
func main() {

	var wg sync.WaitGroup
	m := sync.Mutex{}
	c := sync.NewCond(&m)

	wg.Add(1)
	go func ()  {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRes) == 0 {
			// here the go routine will be in a suspended form,
			// it will be in this form until some signal or broadcast
			// message is received

			// if we think about this, this almost works like a channel,
			// with the difference that It doesn't need to send data, it can be in wait
			// sate and it will be suspended and when signal or broadcast  is executed ,
			// it will come out of it's suspended state.
			c.Wait()
		}


		fmt.Println("the values of shared resources",sharedRes["one"])
		c.L.Unlock()
	}

	c.L.Lock()
	sharedRes["one"] = "abhinand"
	c.Signal()
	c.L.Unlock()

	wg,Wait()
}