package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			c <- "hello yopta after " + after.String()
		}()
		return c
	}

	start := time.Now()
	c := or(
		sig(2*time.Millisecond),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(2*time.Second),
		sig(1*time.Minute),
	)

L:
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println(v)
			} else {
				break L
			}
		}
	}

	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	resultChannel := make(chan interface{}, len(channels))
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	go func() {
		wg.Wait()
		close(resultChannel)
	}()
	for _, channel := range channels {
		channel := channel
		go func() {
			for {
				select {
				case v, ok := <-channel:
					if ok {
						resultChannel <- v
					} else {
						wg.Done()
						return
					}
				}
			}
		}()
	}
	return resultChannel
}
