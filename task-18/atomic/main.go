package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count atomic.Int64
}

func newCounter() *Counter {
	counter := Counter{count: atomic.Int64{}}
	counter.count.Store(0)
	return &counter
}

func (counter *Counter) increment() {
	counter.count.Add(1)
}

func main() {

	var wg sync.WaitGroup
	counter := newCounter()

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter.count.Load())
}
