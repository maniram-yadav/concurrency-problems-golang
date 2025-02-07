package threadpool

import (
	"fmt"
	"sync"
)

type ThreadPool struct {
	workers     int
	taskQueue   chan Task
	stopChannel chan struct{}
	resizeMutex sync.Mutex
	workerGroup sync.WaitGroup
}

func NewThreadPool(initWorkers int) *ThreadPool {
	if initWorkers <= 0 {
		panic("Invalid number of workers provided")
	}
	threadPool := ThreadPool{
		workers:     initWorkers,
		taskQueue:   make(chan Task, 100),
		stopChannel: make(chan struct{}, initWorkers),
	}
	threadPool.startWorkers(initWorkers)
	return &threadPool
}

func (t *ThreadPool) startWorkers(workersCount int) {

	t.workerGroup.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func(workerId int) {
			defer t.workerGroup.Done()
			for {
				select {
				case task := <-t.taskQueue:
					fmt.Printf("\n Worker %d executing task %d ", workerId, task.Id)
					task.Execute()
				case <-t.stopChannel:
					fmt.Printf("\nWorker %d stopping", workerId)
					return
				}
			}
		}(i)
	}

}

func (t *ThreadPool) Resize(workersCount int) {
	currentWorkers := t.workers

	t.resizeMutex.Lock()
	defer t.resizeMutex.Unlock()

	if workersCount <= 0 {
		panic("Worker count is less than 1")
	}
	if workersCount == currentWorkers {
		return
	}

	if workersCount < currentWorkers {
		workersToStop := currentWorkers - workersCount
		for i := 0; i < workersToStop; i++ {
			t.stopChannel <- struct{}{}
		}

	} else {
		additionalWorkers := workersCount - currentWorkers
		t.startWorkers(additionalWorkers)
	}
	t.workers = workersCount
}

func (t *ThreadPool) Submit(task Task) {
	t.taskQueue <- task
}

func (t *ThreadPool) Shutdown() {
	close(t.stopChannel)
	t.workerGroup.Wait()
	close(t.taskQueue)
}
