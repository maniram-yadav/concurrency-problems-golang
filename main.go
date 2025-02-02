package main

import (
	blockingqueue "concurrency-patterns/blocking-queue"
	diningphilosophers "concurrency-patterns/dining-philosophers"
	threadpool "concurrency-patterns/thread-pool"
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Dining Philosphors")
	var forks [diningphilosophers.PHILOSOPHER_COUNT]*sync.Mutex

	for i := 0; i < diningphilosophers.PHILOSOPHER_COUNT; i++ {
		forks[i] = &sync.Mutex{}
	}
	dineGroup := &sync.WaitGroup{}
	dineGroup.Add(diningphilosophers.PHILOSOPHER_COUNT)

	for i := 0; i < diningphilosophers.PHILOSOPHER_COUNT; i++ {
		philosopher := &diningphilosophers.Philosopher{
			Id:        i + 1,
			LeftFork:  forks[i],
			RightFork: forks[(i+1)%diningphilosophers.PHILOSOPHER_COUNT],
			DineGroup: dineGroup,
		}
		go philosopher.Dine()
	}
	dineGroup.Wait()
	fmt.Println("All philosphers have finished dining")
	threadpool := threadpool.NewThreadPool(4)
	threadpool.Shutdown()
	blockingqueue, _ := blockingqueue.NewBlockingQueue(5)
	blockingqueue.Enqueue(7)

}
