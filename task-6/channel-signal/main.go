package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	exitGoroutine := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		elem := 1
		for {
			select {
			case <-exitGoroutine:
				fmt.Println("Go out from goroutine")
				return
			default:
				fmt.Printf("elem: %d\n", elem)
				elem++
			}
		}
	}()

	time.Sleep(time.Second)

	exitGoroutine <- 1

	wg.Wait()

}
