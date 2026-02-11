package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	errsNum := 0

	workersNum := len(tasks)

	if n < workersNum {
		workersNum = n
	}

	errs := make(chan error, workersNum)
	defer close(errs)

	for i := 0; i < len(tasks); i += workersNum {
		wg.Add(workersNum)

		for j := 0; j < workersNum; j++ {
			go func(taskIdx int) {
				defer wg.Done()
				errs <- tasks[taskIdx]()
			}(i + j)
		}

		wg.Wait()

		for j := 0; j < workersNum; j++ {
			if err := <-errs; err != nil {
				errsNum++
			}
		}

		if m > 0 && errsNum >= m {
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}
