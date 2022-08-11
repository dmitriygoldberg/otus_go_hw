package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// max 0 errors
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	errorLimit := int32(m)
	taskChannel := make(chan Task)
	wg := sync.WaitGroup{}
	wg.Add(n)

	// producer
	go func() {
		for _, task := range tasks {
			if atomic.LoadInt32(&errorLimit) <= 0 {
				break
			}
			taskChannel <- task
		}
		close(taskChannel)
	}()

	// consumer
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range taskChannel {
				if err := task(); err != nil {
					atomic.AddInt32(&errorLimit, -1)
				}
			}
		}()
	}

	wg.Wait()

	if errorLimit <= 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
