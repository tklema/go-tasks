package main

import (
	"fmt"
	"sync"
)

func main() {

	var cmap sync.Map

	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range 10 {
				cmap.Store(j, i)
			}
		}()
	}

	wg.Wait()

	cmap.Range(func(k, v interface{}) bool {
		fmt.Printf("key: %d, value: %d\n", k, v)
		return true
	})

}
