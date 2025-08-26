package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		elem := 1
		for {
			switch {
			case elem > 5:
				fmt.Println("Go out from goroutine")
				return
			default:
				fmt.Printf("elem: %d\n", elem)
				elem++
			}
		}
	}()

	wg.Wait()

}
