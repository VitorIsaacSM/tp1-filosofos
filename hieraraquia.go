// Hierarquia de filosofos usando Mutex como garfos, o ultimo filoso pega os garfos na ordem invertida

package main

import (
	"fmt"
	"sync"
	"time"
	"strings"
)

type Fork struct {
	sync.Mutex
	id int
}

type Philosopher struct {
	id              int
	leftFork, rightFork *Fork
	eatenCount      int
}

var time_to_think = 2
var time_to_eat = 3

func NewPhilosopher(id int, leftFork, rightFork *Fork) *Philosopher {
	return &Philosopher{
		id:         id,
		leftFork:   leftFork,
		rightFork:  rightFork,
		eatenCount: 0,
	}
}

func (p *Philosopher) dine() {
	p.acquireForks()
	p.eat()
	p.releaseForks()
}

func (p *Philosopher) start(wg *sync.WaitGroup, maxEatingCount int) {
	defer wg.Done()
	for p.eatenCount < maxEatingCount {
		p.dine()
		p.think()
	}
}

func (p *Philosopher) acquireForks() {
	if p.leftFork.id < p.rightFork.id {
		p.leftFork.Lock()
		p.rightFork.Lock()
	} else {
		p.rightFork.Lock()
		p.leftFork.Lock()
	}
}

func (p *Philosopher) releaseForks() {
	p.leftFork.Unlock()
	p.rightFork.Unlock()
}

func (p *Philosopher) eat() {
	var sb strings.Builder

	for s := 0; s <= p.id*5; s++ {
		sb.WriteString(" ")
	}
	fmt.Printf(sb.String() + " E%d\n", p.eatenCount + 1) // Eating
	p.eatenCount++

	time.Sleep(time.Duration(time_to_eat) * time.Second)
}

func (p *Philosopher) think() {
    var sb strings.Builder

    for s := 0; s <= p.id*5; s++ {
      sb.WriteString(" ")
		}
		fmt.Printf(sb.String() + " T%d\n",  p.eatenCount) // Thinking
		time.Sleep(time.Duration(time_to_think) * time.Second)
}

func main() {
	n := 5
	maxEatingCount := 5

	forks := make([]*Fork, n)
	for i := 0; i < n; i++ {
		forks[i] = &Fork{id: i}
	}

	philosophers := make([]*Philosopher, n)
	for i := 0; i < n; i++ {
		philosophers[i] = NewPhilosopher(i, forks[i], forks[(i+1)%n])
	}

	var wg sync.WaitGroup
	wg.Add(n)
	start := time.Now()

	fmt.Println("\n [P1] [P2] [P3] [P4] [P5] \n")

	for _, p := range philosophers {
		go p.start(&wg, maxEatingCount)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}
