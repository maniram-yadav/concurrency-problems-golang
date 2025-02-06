package jobscheduler

type Job struct {
	Id      string
	Execute func() error
	Depends []string
}
