package main

import (
	"fmt"
)

func binSearch(arr []int, find int) int {

	l, r := 0, len(arr)

	for r-l > 1 {
		m := (r + l) / 2
		if arr[m] <= find {
			l = m
		} else {
			r = m
		}
	}

	if arr[l] == find {
		return l
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 3, 5, 6, 7, 7, 7, 8, 9}

	fmt.Println(binSearch(arr, 5))
	fmt.Println(binSearch(arr, 10))
	fmt.Println(binSearch(arr, 4))
}
