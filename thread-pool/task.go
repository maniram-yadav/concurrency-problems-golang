package threadpool

type Task struct {
	Id      int
	Execute func()
}
