package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidWorkerAmount = errors.New("invalid workers amount")
)

type Task func() error

type errorCounter struct {
	counter int
	noLimit bool
	mu      sync.Mutex
}

func (ec *errorCounter) ErrorExceed() bool {
	if ec.noLimit {
		return false
	}

	ec.mu.Lock()
	defer ec.mu.Unlock()

	return ec.counter <= 0
}

func (ec *errorCounter) Update() error {
	if ec.noLimit {
		return nil
	}

	ec.mu.Lock()
	defer ec.mu.Unlock()

	if ec.counter <= 0 {
		return ErrErrorsLimitExceeded
	}

	ec.counter -= ec.counter

	return nil
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrInvalidWorkerAmount
	}

	wg := new(sync.WaitGroup)
	wg.Add(n)

	errorCounter := &errorCounter{
		mu:      sync.Mutex{},
		counter: m,
	}

	if m <= 0 {
		errorCounter.noLimit = true
	}

	taskCH := make(chan Task)

	for i := 0; i < n; i++ {
		go worker(wg, taskCH, errorCounter)
	}

	go func() {
		for _, t := range tasks {
			if errorCounter.ErrorExceed() {
				break
			}

			taskCH <- t
		}
		close(taskCH)
	}()

	wg.Wait()

	if errorCounter.ErrorExceed() {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(wg *sync.WaitGroup, taskCH chan Task, counter *errorCounter) {
	defer wg.Done()

	for task := range taskCH {
		if err := task(); err != nil {
			if cerr := counter.Update(); cerr != nil {
				break
			}
		}
	}
}
