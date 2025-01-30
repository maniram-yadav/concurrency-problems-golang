package readerwriterlock

import "sync"

type ReaderWriterLock struct {
	lock        sync.Mutex
	writerLock  sync.Mutex
	readerCount int
}

func NewReaderWriterLock() *ReaderWriterLock {
	return &ReaderWriterLock{}
}

func (rw *ReaderWriterLock) LockRead() {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	if rw.readerCount == 0 {
		rw.writerLock.Lock()
	}
	rw.readerCount++
}

func (rw *ReaderWriterLock) UnlockRead() {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.readerCount--
	if rw.readerCount == 0 {
		rw.writerLock.Unlock()
	}

}

func (rw *ReaderWriterLock) LockWrite() {
	rw.writerLock.Lock()
}

func (rw *ReaderWriterLock) UnlockWrite() {
	rw.writerLock.Unlock()
}
