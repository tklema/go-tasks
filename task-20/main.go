package main

import (
	"fmt"
)

func reverseRune(r []rune) []rune {
	n := len(r)

	for i := range n / 2 {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}

	return r
}

func reverseWordsOrder(s string) string {
	r := []rune(s)

	reverseRune(r)

	start := 0
	for i := range r[1:] {
		if r[i] == ' ' {
			reverseRune(r[start:i])
			start = i + 1
		}
	}

	reverseRune(r[start:])

	return string(r)
}

func main() {
	s1 := "snow dog sun"
	fmt.Println(reverseWordsOrder(s1))
}
