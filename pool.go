package pool

import (
	"sync"
)

type Task func()

type ThreadPoolExecutor struct {
	workerCount int
	taskQueue   chan Task
	wg          sync.WaitGroup
}

func NewThreadPoolExecutor(workerCount int) *ThreadPoolExecutor {
	return &ThreadPoolExecutor{
		workerCount: workerCount,
		taskQueue:   make(chan Task),
	}
}

func (executor *ThreadPoolExecutor) Start() {
	for i := 0; i < executor.workerCount; i++ {
		go executor.worker()
	}
}

func (executor *ThreadPoolExecutor) Submit(task func()) {
	executor.wg.Add(1)
	executor.taskQueue <- task
}

func (executor *ThreadPoolExecutor) Wait() {
	close(executor.taskQueue)
	executor.wg.Wait()
}

func (executor *ThreadPoolExecutor) worker() {
	for task := range executor.taskQueue {
		task()
		executor.wg.Done()
	}
}
