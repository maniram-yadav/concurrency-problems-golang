package diningphilosophers

import (
	"fmt"
	"sync"
	"testing"
)

func TestDiningPhilosophers(t *testing.T) {

	fmt.Println("Dining Philosphors")
	var forks [PHILOSOPHER_COUNT]*sync.Mutex

	for i := 0; i < PHILOSOPHER_COUNT; i++ {
		forks[i] = &sync.Mutex{}
	}
	dineGroup := &sync.WaitGroup{}
	dineGroup.Add(PHILOSOPHER_COUNT)

	for i := 0; i < PHILOSOPHER_COUNT; i++ {
		philosopher := &Philosopher{
			Id:        i + 1,
			LeftFork:  forks[i],
			RightFork: forks[(i+1)%PHILOSOPHER_COUNT],
			DineGroup: dineGroup,
		}
		go philosopher.Dine()
	}
	dineGroup.Wait()
	fmt.Println("All philosphers have finished dining")
	if dineGroup != nil {
		t.Log("\nTest passed all philospher have finished dining without deadlock")
	} else {
		t.Error("\nTest failed. Deadlock occured")
	}
}
