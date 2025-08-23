package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print("invalid arguments: should contain number of workers")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print("invalid arguments: error during parsing the number of workers")
		os.Exit(1)
	}
	if n < 1 {
		fmt.Print("invalid arguments: number of workers should be positive")
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	/*
		Создаем канал, в который придет уведомление о прерывании работы программы
	*/
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	go func() {
		<-shutdown
		fmt.Println("Shutdown signal was received")
		cancel()
	}()

	// Позволяет настроить ожидание завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(n)
	ch := make(chan int)
	for i := range n {
		go func() {
			defer wg.Done()
			for {
				/*
					Есть 2 пути развития:
					1) Прерывание программы (ожидание получения сигнала из ctx.Done())
					2) Получение данных
					Если получили прерывание, то завершаем функцию
				*/
				select {
				case <-ctx.Done():
					fmt.Printf("worker %d is shutdowned\n", i+1)
					return
				case output := <-ch:
					fmt.Println(output)
				}
			}
		}()
	}

	elem := 1
	for {
		select {
		/*
			Аналогично воркерам
		*/
		case <-ctx.Done():
			wg.Wait()
			close(ch)
			fmt.Println("main worker is shutdowned")
			fmt.Println("All Done")
			return
		case ch <- elem:
			elem = (elem + 1) % 1e9
		}
	}
}
