package main

import (
	"fmt"
	"time"
)

func Sleep(duration int) {
	start := time.Now()
	for time.Since(start) < time.Duration(duration)*time.Second {
	}
}

func main() {
	sleepingTimeSeconds := 5

	start := time.Now()

	Sleep(sleepingTimeSeconds)

	fmt.Printf("working time: %v \n", time.Since(start))
}
