package diningphilosophers

import (
	"fmt"
	"sync"
	"time"
)

const PHILOSOPHER_COUNT = 5

type Philosopher struct {
	Id        int
	LeftFork  *sync.Mutex
	RightFork *sync.Mutex
	DineGroup *sync.WaitGroup
}

func (p *Philosopher) think() {
	fmt.Printf("\nPhilospher %d is thinking", p.Id)
	time.Sleep(time.Duration(1+2*p.Id) * time.Second)
}

func (p *Philosopher) eat() {

	if p.Id%2 == 0 {
		p.LeftFork.Lock()
		p.RightFork.Lock()
	} else {
		p.RightFork.Lock()
		p.LeftFork.Lock()
	}
	fmt.Printf("\nPhilosopher %d is eating", p.Id)
	time.Sleep(2 * time.Second)
	p.LeftFork.Unlock()
	p.RightFork.Unlock()
	fmt.Printf("\nPhilosopher %d has finished the eating", p.Id)
}

func (p *Philosopher) Dine() {
	defer p.DineGroup.Done()
	for i := 0; i < 4; i++ {
		p.think()
		p.eat()
	}
}
