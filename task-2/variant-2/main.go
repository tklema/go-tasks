package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := []int{2, 4, 6, 8, 10}
	length := len(arr)

	var wg sync.WaitGroup
	wg.Add(5)
	for i := range length {
		go func() {
			defer wg.Done()
			arr[i] *= arr[i]
		}()
	}

	wg.Wait()

	fmt.Println(arr)
}
