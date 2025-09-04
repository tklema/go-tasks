package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var left []int
	var right []int

	mid := arr[0]

	for _, elem := range arr[1:] {
		if elem <= mid {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}

	leftSort := quickSort(left)
	rightSort := quickSort(right)

	var output []int
	output = append(output, leftSort...)
	output = append(output, mid)
	output = append(output, rightSort...)

	return output
}

func main() {
	arr := []int{5, 3, 7, 9, 2, 1, 4, 3}

	fmt.Println(quickSort(arr))
}
