package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	chanFirst := make(chan int)
	chanSecond := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	go func() {
		<-shutdown
		fmt.Println("Work was interrupted")
		cancel()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chanFirst)
		defer fmt.Println("Generator end")

		fmt.Println("Generator start")

		for _, elem := range numbers {
			select {
			case <-ctx.Done():
				fmt.Println("Generator was interrupted")
				return
			case chanFirst <- elem:
			}
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chanSecond)
		defer fmt.Println("Sender end")

		fmt.Println("Sender start")

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Sender was interrupted")
				return
			case elem, ok := <-chanFirst:
				if !ok {
					return
				}
				chanSecond <- elem * 2
			}
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Output end")

		fmt.Println("Output start")

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Output was interrupted")
				return
			case elem, ok := <-chanSecond:
				if !ok {
					return
				}
				fmt.Println(elem)
			}
		}

	}()

	wg.Wait()
}
