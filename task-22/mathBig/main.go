package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// для очень больших чисел используем math/big
func main() {

	a := strconv.Itoa(1<<20 + 100)
	b := strconv.Itoa(1<<20 + 7)

	fmt.Printf("a: %v \n", a)
	fmt.Printf("b: %v \n", b)

	aBig := new(big.Int)
	bBig := new(big.Int)

	aBig.SetString(a, 10)
	bBig.SetString(b, 10)

	sum := new(big.Int)
	sum.Add(aBig, bBig)
	fmt.Printf("a + b = %v \n", sum)

	sub := new(big.Int)
	sub.Sub(aBig, bBig)
	fmt.Printf("a - b = %v \n", sub)

	mul := new(big.Int)
	mul.Mul(aBig, bBig)
	fmt.Printf("a * b = %v \n", mul)

	div := new(big.Int)
	div.Div(aBig, bBig)
	fmt.Printf("a / b = %v \n", div)
}
