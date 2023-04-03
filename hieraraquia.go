// Hierarquia de filosofos usando Mutex como garfos, o ultimo filoso pega os garfos na ordem invertida

package main

import (
	"fmt"
	"sync"
	"time"
)

type fork struct{ sync.Mutex }

var dining sync.WaitGroup

func Philosopher(id, iteration int, forkA *fork, forkB *fork) {
	time_to_think := 2
	time_to_eat := 3

	forkA.Lock() // pick up forks
	forkB.Lock()

	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration) // Thinking
	time.Sleep(time.Duration(time_to_think) * time.Second)

	// eat

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration) // Eating
	time.Sleep(time.Duration(time_to_eat) * time.Second)

	forkA.Unlock() // put down forks
	forkB.Unlock()

	dining.Done()
}

func main() {
	philosophers := 5
	rounds := 5

	// Create forks
	forks := make([]*fork, philosophers)
	for i := 0; i < philosophers; i++ {
		forks[i] = new(fork)
	}

	dining.Add(philosophers * rounds)

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	// run philosophers

	start := time.Now()

	for i := 0; i < rounds; i++ {
		for j := 0; j < philosophers; j++ {
			// if last philosopher, make him left handed
			if j == (philosophers - 1) {
				go Philosopher(j, i + 1, forks[0], forks[j])
			} else {
				go Philosopher(j, i + 1, forks[j], forks[j + 1])
			}
		}
	}

	dining.Wait() // wait for all philosophers to eat

	elapsed := time.Since(start)
	fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}
