package main

import (
	"fmt"
)

func main() {

	a := 7
	b := 25

	fmt.Printf("a: %d, b: %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a: %d, b: %d\n", a, b)
}
