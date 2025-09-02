package main

import (
	"fmt"
)

func main() {

	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	set := make(map[int]bool)

	for _, a := range A {
		set[a] = true
	}

	var answer []int

	for _, b := range B {
		if set[b] {
			answer = append(answer, b)
		}
	}

	fmt.Printf("%v", answer)
}
