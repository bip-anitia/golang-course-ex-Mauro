package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	ID      int
	Data    interface{}
	Process func(interface{}) (interface{}, error)
}

type Result struct {
	TaskID int
	Value  interface{}
	Error  error
}

type WorkerPool struct {
	numWorkers int
	tasks      chan Task
	results    chan Result
	done       chan struct{}
	wg         sync.WaitGroup
}

func NewWorkerPool(n int) *WorkerPool {
	return &WorkerPool{
		numWorkers: n,
		tasks:      make(chan Task, n),
		results:    make(chan Result, n),
		done:       make(chan struct{}),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			wp.results <- Result{TaskID: -1, Error: fmt.Errorf("panic: %v", r)}
		}
	}()
	for task := range wp.tasks {
		val, err := task.Process(task.Data)
		wp.results <- Result{TaskID: task.ID, Value: val, Error: err}
	}
}

func (wp *WorkerPool) Submit(task Task) {
	wp.tasks <- task
}

func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp *WorkerPool) Stop() {
	close(wp.tasks)
	wp.wg.Wait()
	close(wp.results)
}

func main() {
	workers := flag.Int("workers", 5, "number of workers")
	tasks := flag.Int("tasks", 100, "number of tasks")
	flag.Parse()
	start := time.Now()
	var total, success, failed int64

	pool := NewWorkerPool(*workers)
	pool.Start()
	defer pool.Stop()

	numTasks := *tasks
	go func() {
		for i := 0; i < numTasks; i++ {
			task := Task{ID: i, Data: i, Process: func(d interface{}) (interface{}, error) {
				time.Sleep(100 * time.Millisecond)
				return fmt.Sprintf("done %v", d), nil
			}}
			pool.Submit(task)
		}
	}()

	for i := 0; i < numTasks; i++ {
		res := <-pool.Results()
		fmt.Printf("Task %d: %v (err=%v)\n", res.TaskID, res.Value, res.Error)
		atomic.AddInt64(&total, 1)
		if res.Error != nil {
			atomic.AddInt64(&failed, 1)
		} else {
			atomic.AddInt64(&success, 1)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("\nStatistics:\n")
	fmt.Printf("  Total tasks: %d\n", total)
	fmt.Printf("  Successful: %d\n", success)
	fmt.Printf("  Failed: %d\n", failed)
	fmt.Printf("  Total time: %s\n", elapsed)
	fmt.Printf("  Throughput: %.2f tasks/sec\n", float64(total)/elapsed.Seconds())

}
