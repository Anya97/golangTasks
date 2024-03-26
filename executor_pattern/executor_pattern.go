package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

func main() {
	fmt.Println(Run([]Task{
		func() error {
			time.Sleep(200)
			return errors.New("oshibka")
		},
		func() error {
			time.Sleep(200)
			return errors.New("oshibka")
		},
		func() error {
			time.Sleep(200)
			return nil
		},
		func() error {
			time.Sleep(200)
			return nil
		},
		func() error {
			time.Sleep(200)
			return nil
		},
		func() error {
			time.Sleep(200)
			return nil
		},
	}, 2, 2))
}

type Task func() error

func Run(tasks []Task, n, m int) error {
	jobChan := make(chan Task, n)
	resultChan := make(chan error, n)
	countOfError := 0
	wg := &sync.WaitGroup{}

	for i := 0; i < n; i++ {
		go Work(jobChan, wg, resultChan)
	}

	pushed := min(n, len(tasks))
	for i := 0; i < pushed; i++ {
		wg.Add(1)
		jobChan <- tasks[i]
	}

	for i := range tasks {
		result := <-resultChan
		if result != nil {
			countOfError += 1
			if countOfError >= m {
				break
			}
		}
		if i+pushed < len(tasks) {
			wg.Add(1)
			jobChan <- tasks[i+pushed]
		}
	}

	close(jobChan)
	wg.Wait()
	close(resultChan)

	if countOfError >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func Work(jobsChan <-chan Task, wg *sync.WaitGroup, resultChan chan<- error) {
	for {
		select {
		case job, ok := <-jobsChan:
			if !ok {
				return
			}
			resultChan <- job()
			wg.Done()
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
