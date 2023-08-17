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
var thinkTime = 3 * time.Second  // philosopher think time (probably after eating)
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
func diningProblem(philosopher Philosopher,wg *sync.WaitGroup,forks map[int]*sync.Mutex,seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n",philosopher.name)
	seated.Done()

	// wait till all the philosophers are seated
	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {

		// chose the smaller number fork to avoid race condtion
		if philosopher.leftFork > philosopher.rightFork {
				forks[philosopher.rightFork].Lock()
				fmt.Printf("\t%s takes the right fork.\n",philosopher.name)	
				forks[philosopher.leftFork].Lock()
				fmt.Printf("\t%s takes the left fork.\n",philosopher.name)
		} else {
				forks[philosopher.leftFork].Lock()
				fmt.Printf("\t%s takes the left fork.\n",philosopher.name)
				forks[philosopher.rightFork].Lock()
				fmt.Printf("\t%s takes the right fork.\n",philosopher.name)	
		}


		// get a lck for both forks
		// we will bot be able to lock if it is already locked -- in simple terms - only one philosopher can use a set of forks at a time
	

		fmt.Printf("\t%s has both forks ans is eating.\n",philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n",philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n",philosopher.name)
	}

	fmt.Println(philosopher.name,"is satisfied.")
	fmt.Println(philosopher.name,"left the table.")
}


// there is a situation where all the philosopher take in one fork and 
// don't get a chance to take the other as it is already take by the other philosophers