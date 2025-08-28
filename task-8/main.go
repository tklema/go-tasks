package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("invalid number of arguments")
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

	fmt.Printf("n bitnumber: %b, n: %d\n", n, n)
	for i := range 64 {
		if 1<<i > n {
			break
		}
		m1 := n | (1 << i)
		fmt.Printf("i: %d, bit: 1, bitnumber: %b, value: %d\n", i, m1, m1)
		m2 := n &^ (1 << i)
		fmt.Printf("i: %d, bit: 0, bitnumber: %b, value: %d\n", i, m2, m2)
	}

}
