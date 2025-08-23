package main

import "fmt"

func main() {

	arr := []int{2, 4, 6, 8, 10}
	length := len(arr)

	c := make(chan int, length)
	for i := range length {
		go func() {
			arr[i] *= arr[i]
			c <- 1
		}()
	}

	for range length {
		<-c
	}

	fmt.Println(arr)
}
