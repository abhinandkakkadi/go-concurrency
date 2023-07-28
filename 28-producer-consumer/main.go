package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

// data will be sending details about the order
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	// we have to return error, which will be null if we successfully cancel the channel
	// it won't be null if we try to close the channel and something went wrong
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		// + 1 since we want to delay it for at least one second
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number #%d!\n", pizzaNumber)

		// now we will create a condition where the pizza making process might fail
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		// total represents both pizzaMade and pizzaFailed
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf(" *** we ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0
	// run forever or until we receive a quit notification
	// try to make pizza's
	for {
		currentPizza := makePizza(i)
		//  try to make pizza's
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
				// we tried to make a pizza (we sent something to the data channel)
				case pizzaMaker.data <- *currentPizza:

				case quitChan := <- pizzaMaker.quit:
					// close channels - (always close channel if created)
					close(pizzaMaker.data)
					// close quitChan that we have created above (intialzed means created)
					close(quitChan)
					return
			}
		}

	}
}

func main() {

	// seed the random number generator
	// if we don't do this it will get same details every time we run this program
	rand.Seed(time.Now().UnixNano())

	// create a producer
	// once we have created a channel, we have to close that
	// that's why we are maintaining the chan here
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer
	go pizzeria(pizzaJob)

}
