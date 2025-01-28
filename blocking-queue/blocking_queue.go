package blockingqueue

import (
	"container/list"
	"sync"
)

type BoundedBlockingQueue struct {
	capacity int
	list     *list.List
	mutex    sync.Mutex
	cond     *sync.Cond
}

func NewBlockingQueue(capacity int) (*BoundedBlockingQueue, error) {
	if capacity <= 0 {
		panic("capacity should be greater than 0")
	}
	blockingQueue := &BoundedBlockingQueue{
		capacity: capacity,
		list:     list.New(),
	}
	blockingQueue.cond = sync.NewCond(&blockingQueue.mutex)
	return blockingQueue, nil
}

func (q *BoundedBlockingQueue) Enqueue(value interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.list.Len() >= q.capacity {
		q.cond.Wait()
	}

	q.list.PushBack(value)
	q.cond.Broadcast()

}

func (q *BoundedBlockingQueue) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.list.Len() == 0 {
		q.cond.Wait()
	}
	value := q.list.Remove(q.list.Front())
	q.cond.Broadcast()
	return value
}

func (q *BoundedBlockingQueue) Peek() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.list.Len() == 0 {
		q.cond.Wait()
	}
	value := q.list.Front().Value
	return value
}

func (q *BoundedBlockingQueue) Size() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return q.list.Len()
}

func (q *BoundedBlockingQueue) Capacity() int {
	return q.capacity
}
