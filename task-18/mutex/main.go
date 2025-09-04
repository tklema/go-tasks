package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	count int
}

func newCounter() *Counter {
	return &Counter{count: 0}
}

func (counter *Counter) increment() {
	counter.mutex.Lock()
	counter.count++
	counter.mutex.Unlock()
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

	fmt.Println(counter.count)
}
