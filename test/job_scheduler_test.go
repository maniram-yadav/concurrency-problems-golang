package test

import (
	js "concurrency-patterns/jobscheduler"
	"fmt"
	"testing"
)

func TestJobScheduler(t *testing.T) {

	jobScheduler := js.NewJobScheduler(4)

	jobScheduler.AddJob(&js.Job{
		Id: "A",
		Execute: func() error {
			fmt.Println("Executing Job A")
			return nil
		},
	})
	jobScheduler.AddJob(&js.Job{
		Id:      "B",
		Depends: []string{"C"},
		Execute: func() error {
			fmt.Println("\nExecuting Job B")
			return nil
		},
	})

	jobScheduler.AddJob(&js.Job{
		Id:      "C",
		Depends: []string{"A"},
		Execute: func() error {
			fmt.Println("\nExecuting Job B")
			return nil
		},
	})

	jobScheduler.Run()
	fmt.Println("\nAll jobs completed")
}
