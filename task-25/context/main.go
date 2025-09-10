package main

import (
	"context"
	"fmt"
	"time"
)

func Sleep(duration int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel()

	<-ctx.Done()
}

func main() {
	sleepingTimeSeconds := 5

	start := time.Now()

	Sleep(sleepingTimeSeconds)

	fmt.Printf("working time: %v \n", time.Since(start))
}
