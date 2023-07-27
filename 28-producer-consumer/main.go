package main

import (
	"errors"
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
	return <- ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making

	// run forever or until we receive a quit notification
	// try to make pizza's
	for {
		//  try to make pizza's 
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
