package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		elem := 1
		for {
			select {
			case <-ctx.Done():
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
