package main

import (
	"fmt"
	"sync"
	"time"
)

//

// when someone is using the left fork or right fork there will be lock upon those forks
// philosopher is a struct which stores information about a philosopher
type Philosopher struct {
	name string 
	rightFork int
	leftFork int
} 

// philosopher is a list of all philosophers
// there are 5 forks - 0 -  (look at the diagram in case of any doubts)
// note that left fork of socrates is right fork of plato
var philosophers = []Philosopher{
	{name: "Plato", leftFork:4, rightFork:0},
	{name: "Socrates", leftFork:0, rightFork:1},
	{name: "Aristotle", leftFork:1, rightFork:2},
	{name: "Pascal", leftFork:2, rightFork:3},
	{name: "Locke", leftFork:3, rightFork:4},

}

// define some variables
var hunger = 3 // how many time a person eats
var eatTime = 1 * time.Second // time for eating food 
var think = 3 * time.Second  // philosopher think time (probably after eating)
var sleep = 1 * time.Second



func main() {

	// program starts
	fmt.Println("dining philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is currently empty.")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("the table is empty")

	// 

}

func dine() {
	// this will be done when everyone is done eating
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// this will be done when everyone is seated
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	// we need to lock the access to the fork by the philosopher (when they are eating)
	// once we create a mutex, we should never copy that. But we can use pointer (i.e use address)
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}


	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire of the go routine for the current philosopher
		go diningProblem(philosophers[i],wg,forks,seated)
	}

	wg.Wait()
}

// this function we be run for all the philosophers in the philospohers slice
func diningProblem(pilosopher Philosopher,wg *sync.WaitGroup,forks map[int]*sync.Mutex,seated *sync.WaitGroup) {
	defer wg.Done()
	
}