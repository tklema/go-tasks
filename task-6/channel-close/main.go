package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for elem := range ch {
			fmt.Printf("elem: %d\n", elem)
		}
		fmt.Println("Go out from goroutine")
	}()

	for elem := range 10 {
		ch <- elem
	}

	close(ch)

	wg.Wait()

}
