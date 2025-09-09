package main

import (
	"fmt"
)

// для просто чисел, больших 2^20, можно использовать int64
func main() {

	a := int64(1<<20 + 100)
	b := int64(1<<20 + 7)

	fmt.Printf("a: %v \n", a)
	fmt.Printf("b: %v \n", b)

	fmt.Printf("a + b = %v \n", a+b)
	fmt.Printf("a - b = %v \n", a-b)
	fmt.Printf("a * b = %v \n", a*b)
	fmt.Printf("a / b = %v \n", a/b)

}
