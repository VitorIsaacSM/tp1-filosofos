// Solução utilizando um canal arbitrador, que permite apenas que um filoso coma por vez

package main

import (
    "fmt"
    "sync"
		"time"
)

type Philosopher struct {
    id                int
    left, right       *sync.Mutex
    arbitratorChannel chan bool
}

var time_to_think = 2
var time_to_eat = 3

func (p Philosopher) eat(iteration int) {
    p.right.Lock()
    p.left.Lock()

		for s := 1; s <= p.id*6; s++ {
			fmt.Printf(" ")
		}
		fmt.Printf(" E%d\n", iteration + 1) // Eating


		time.Sleep(time.Duration(time_to_eat) * time.Second)

    p.right.Unlock()
    p.left.Unlock()
}

func (p Philosopher) think(iteration int) {

		for s := 1; s <= p.id*6; s++ {
			fmt.Printf(" ")
		}
		fmt.Printf(" T%d\n", iteration + 1) // Thinking

		time.Sleep(time.Duration(time_to_think) * time.Second)
}

func (p Philosopher) start(arbitratorChannel chan bool, wg *sync.WaitGroup) {
  for i := 0; i < 5; i++ {
    wg.Add(1)
    <-arbitratorChannel // wait for permission to eat
    p.eat(i)
    arbitratorChannel <- true // inform arbitrator that eating is done
    p.think(i)
    wg.Done()
  }
}

func main() {
    // create 5 chopsticks
    chopsticks := make([]*sync.Mutex, 5)
    for i := 0; i < 5; i++ {
        chopsticks[i] = &sync.Mutex{}
    }

    // create arbitrator channel
    arbitratorChannel := make(chan bool, 1)
    arbitratorChannel <- true // start with permission to eat

		start := time.Now()

    // create 5 philosophers
    var wg sync.WaitGroup

		fmt.Println("\n [P1] [P2] [P3] [P4] [P5] \n")


    for i := 0; i < 5; i++ {
        left := chopsticks[i]
        right := chopsticks[(i+1)%5]
        philosopher := Philosopher{i, left, right, arbitratorChannel}
        go philosopher.start(arbitratorChannel, &wg)
    }
    wg.Wait()

		elapsed := time.Since(start)
		fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}