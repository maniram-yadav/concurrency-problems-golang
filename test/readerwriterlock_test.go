package test

import (
	rwlock "concurrency-patterns/threaded-rwlcok"
	"sync"
	"testing"
)

func TestReaderWriterLock(t *testing.T) {

	rwLock := rwlock.ReaderWriterLock{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(readrId int) {
			defer wg.Done()
			rwLock.LockRead()
			println("reader ", readrId, " is reading")
			rwLock.UnlockRead()
			println("reader ", readrId, " have finished the reading")

		}(i + 1)
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		rwLock.LockWrite()
		println("writer is writing")
		rwLock.UnlockWrite()
		println("writer have finished the writing")

	}()
	wg.Wait()
	t.Log("RWLock implemented completed")

}
