package main

import (
	"fmt"
	"sync"
)

func main() {

	simpleMap := make(map[int]int)
	var mutex sync.Mutex

	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range 10 {
				mutex.Lock()
				simpleMap[j] = i
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()

	for k, v := range simpleMap {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}

}
