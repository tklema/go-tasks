package main

import (
	"fmt"
)

func reverseString(s string) string {
	r := []rune(s)
	n := len(r)

	for i := range n / 2 {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}

	return string(r)
}

func main() {
	s1 := "главрыба"
	s2 := "гларыба"
	fmt.Println(reverseString(s1))
	fmt.Println(reverseString(s2))
}
