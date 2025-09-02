package main

import (
	"fmt"
)

func main() {

	input := []string{"cat", "cat", "dog", "cat", "tree", "dog", "tree"}

	set := make(map[string]bool)

	for _, elem := range input {
		set[elem] = true
	}

	fmt.Printf("{ ")
	for k := range set {
		fmt.Printf("%s ", k)
	}
	fmt.Printf("}")
}
