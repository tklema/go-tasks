package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	seconds := 2

	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case elem := <-ch:
				fmt.Printf("Read: %d\n", elem)
			case <-ctx.Done():
				fmt.Println("End of reading")
				return
			}
		}
	}()

	timer := time.NewTimer(time.Duration(seconds) * time.Second)

	elem := 1
	for {
		select {
		case ch <- elem:
			elem++
		case <-timer.C:
			cancel()
			wg.Wait()
			close(ch)
			fmt.Println("End of work")
			return
		}
	}
}
