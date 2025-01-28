package test

import (
	bq "concurrency-patterns/blocking-queue"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockingQueue(t *testing.T) {

	blockingQueue, err := bq.NewBlockingQueue(3)
	if err != nil {
		t.Error("Error creating blocking queue")
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {

		blockingQueue.Enqueue(5)
		blockingQueue.Enqueue(51)
		blockingQueue.Enqueue(23)
		t.Log("All Elements Enqued")
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	assert.Equal(t, blockingQueue.Size(), 3, "Blocking Queue is having some elements")

	go func() {

		t.Logf("Element Dequed %d ", blockingQueue.Dequeue().(int))
		t.Logf("Element Dequed %d ", blockingQueue.Dequeue().(int))
		t.Logf("Element Dequed %d ", blockingQueue.Dequeue().(int))
		t.Log("All Elements Dequed")
		wg.Done()
	}()

	wg.Wait()
	assert.Equal(t, blockingQueue.Size(), 0, "Blocking Queue is now empty")
	t.Log("Blocking Queue test completed")

}
