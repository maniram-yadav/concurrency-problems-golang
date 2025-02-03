package fileprocessor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Task struct {
	FilePath string
	Content  string
}

type WorkerPool struct {
	taskQueue    chan Task
	wg           sync.WaitGroup
	messageGroup sync.WaitGroup
}

func NewWorkerPool(queueSize int) *WorkerPool {
	wp := &WorkerPool{
		taskQueue: make(chan Task, queueSize),
	}
	for i := 0; i < queueSize; i++ {
		go wp.startWorker(i)
	}
	return wp
}

func (wp *WorkerPool) startWorker(workerId int) {
	wp.wg.Add(1)
	defer wp.wg.Done()
	for task := range wp.taskQueue {
		fmt.Printf("\nWorker %d is processing File Path : %s", workerId, task.FilePath)
		err := wp.processFile(task, workerId)
		if err != nil {
			fmt.Printf("\nError processing file Path %s, Error : %v", task.FilePath, err)
		}
	}

}

func (wp *WorkerPool) processFile(task Task, workerid int) error {
	file, err := os.OpenFile(task.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("\nError in opening the file %s", task.FilePath)
		return err
	}

	defer file.Close()
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(task.Content + "\n")
	if err != nil {
		fmt.Printf("\nError in modifying the content file %s", task.FilePath)
		return err
	}

	writer.Flush()
	time.Sleep(time.Duration(workerid+1) * time.Second)
	fmt.Printf("\nContent moidified for the file %s by worker id %d", task.FilePath, workerid)
	wp.messageGroup.Done()
	return nil
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.taskQueue <- task
	wp.messageGroup.Add(1)
}

func (wp *WorkerPool) CountLines(filePath string) (int, error) {

	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)

	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := reader.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func (wp *WorkerPool) WaitForTaskCompletion() {
	wp.messageGroup.Wait()
	fmt.Println("\nAll Task Completed")
}

func (wp *WorkerPool) StopWorker() {
	close(wp.taskQueue)
	wp.wg.Wait()
	fmt.Println("\nAll worker stopped")
}
