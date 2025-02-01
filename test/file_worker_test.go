package test

import (
	fileprocessor "concurrency-patterns/file-processor"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const WORKER_SIZE = 4

const FILE_PATH = "abc.txt"

func TestFileProcessor(t *testing.T) {

	worker := fileprocessor.NewWorkerPool(WORKER_SIZE)
	totalRows := 40
	existingRows, _ := worker.CountLines(FILE_PATH)

	for i := 0; i < totalRows; i++ {
		worker.AddTask(fileprocessor.Task{FilePath: FILE_PATH, Content: "Text " + strconv.Itoa(i+1)})
	}
	worker.WaitForTaskCompletion()

	rows, _ := worker.CountLines(FILE_PATH)

	assert.Equal(t, rows-existingRows, totalRows)
	t.Logf("\nTotal rows before modifying file : %d", rows)
	t.Logf("\nTotal rows after modifying file : %d", existingRows)
	worker.StopWorker()

	t.Log("File Processer test completed")

}
