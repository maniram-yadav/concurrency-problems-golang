package jobscheduler

import (
	"fmt"
	"sync"
)

type JobScheduler struct {
	jobs       map[string]*Job
	workers    int
	readyQueue chan string
	mutex      sync.Mutex
	results    map[string]error
	wg         sync.WaitGroup
}

func NewJobScheduler(workers int) *JobScheduler {
	js := &JobScheduler{workers: workers, jobs: make(map[string]*Job), results: make(map[string]error)}
	return js
}

// naukri
// trent
// srf
// sbicard
// macrotech
// kalyan
// maxhealth
// hindcopper
// indigo
// affle
// j&kbank

func (js *JobScheduler) dfs(jobId string, visited map[string]bool) {
	if visited[jobId] {
		return
	}
	visited[jobId] = true
	js.mutex.Lock()
	job, jobExists := js.jobs[jobId]
	js.mutex.Unlock()
	if !jobExists {
		return
	}
	for _, dep := range job.Depends {
		js.dfs(dep, visited)
	}
	js.wg.Add(1)
	js.readyQueue <- jobId
}
func (js *JobScheduler) AddJob(job *Job) {
	js.jobs[job.Id] = job
}

func (js *JobScheduler) Run() {

	for i := 0; i < js.workers; i++ {
		go js.worker(i + 1)
	}

	var visited = make(map[string]bool)

	for jobId := range js.jobs {
		// if len(js.jobs[jobId].Depends) == 0 {
		js.dfs(jobId, visited)
		// }
	}

	js.wg.Wait()
	close(js.readyQueue)
}

func (js *JobScheduler) worker(workerId int) {
	fmt.Printf("\n=> started Worker with Id %d", workerId)
	for jobId := range js.readyQueue {
		js.mutex.Lock()
		job, jobExists := js.jobs[jobId]
		js.mutex.Unlock()
		if !jobExists {
			continue
		}
		err := job.Execute()
		fmt.Printf("\n=> Worker %s executed Job with Job Id %d", workerId, jobId)
		js.mutex.Lock()
		js.results[jobId] = err
		js.mutex.Unlock()
		js.wg.Done()
	}
}

func (js *JobScheduler) GetExecutionRsults() map[string]error {
	return js.results
}
